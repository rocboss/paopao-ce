package api

import (
	"image"
	"net/url"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/rocboss/paopao-ce/global"
	"github.com/rocboss/paopao-ce/internal/model"
	"github.com/rocboss/paopao-ce/internal/service"
	"github.com/rocboss/paopao-ce/pkg/app"
	"github.com/rocboss/paopao-ce/pkg/convert"
	"github.com/rocboss/paopao-ce/pkg/errcode"
)

func GeneratePath(s string) string {
	n := len(s)
	if n <= 2 {
		return s
	}

	return GeneratePath(s[:n-2]) + "/" + s[n-2:]
}

func GetFileExt(s string) (string, error) {
	switch s {
	case "image/png":
		return ".png", nil
	case "image/jpg":
		return ".jpg", nil
	case "image/jpeg":
		return ".jpeg", nil
	case "image/gif":
		return ".gif", nil
	case "video/mp4":
		return ".mp4", nil
	case "video/quicktime":
		return ".mov", nil
	case "application/zip":
		return ".zip", nil
	default:
		return "", errcode.FileInvalidExt.WithDetails("仅允许 png/jpg/gif/mp4/mov/zip 类型")
	}
}
func GetImageSize(img image.Rectangle) (int, int) {
	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y
	return width, height
}

func fileCheck(uploadType string, size int64) error {
	if uploadType != "public/video" &&
		uploadType != "public/image" &&
		uploadType != "public/avatar" &&
		uploadType != "attachment" {
		return errcode.InvalidParams
	}

	if size > 1024*1024*100 {
		return errcode.FileInvalidSize.WithDetails("最大允许100MB")
	}

	return nil
}

func UploadAttachment(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c)

	uploadType := c.Request.FormValue("type")
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		global.Logger.Errorf("api.UploadAttachment err: %v", err)
		response.ToErrorResponse(errcode.FileUploadFailed)
		return
	}
	defer file.Close()

	if err = fileCheck(uploadType, fileHeader.Size); err != nil {
		cErr, _ := err.(*errcode.Error)
		response.ToErrorResponse(cErr)
		return
	}

	fileExt, err := GetFileExt(fileHeader.Header.Get("Content-Type"))
	if err != nil {
		global.Logger.Errorf("GetFileExt err: %v", err)
		response.ToErrorResponse(err.(*errcode.Error))
		return
	}

	fileReader, err := fileHeader.Open()
	if err != nil {
		global.Logger.Errorf("Attachment file read err: %v", err)
		response.ToErrorResponse(errcode.FileUploadFailed)
		return
	}
	defer fileReader.Close()

	// 生成随机路径
	randomPath := uuid.Must(uuid.NewV4()).String()
	ossSavePath := uploadType + "/" + GeneratePath(randomPath[:8]) + "/" + randomPath[9:] + fileExt

	client, err := oss.New(global.AliossSetting.AliossEndpoint, global.AliossSetting.AliossAccessKeyID, global.AliossSetting.AliossAccessKeySecret)
	if err != nil {
		global.Logger.Errorf("oss.New err: %v", err)
		response.ToErrorResponse(errcode.FileUploadFailed)
		return
	}

	bucket, err := client.Bucket("paopao-assets")
	if err != nil {
		global.Logger.Errorf("client.Bucket err: %v", err)
		response.ToErrorResponse(errcode.FileUploadFailed)
		return
	}

	err = bucket.PutObject(ossSavePath, fileReader)
	if err != nil {
		global.Logger.Errorf("bucket.PutObject err: %v", err)
		response.ToErrorResponse(errcode.FileUploadFailed)
		return
	}

	// 构造附件Model
	attachment := &model.Attachment{
		FileSize: fileHeader.Size,
		Content:  "https://" + global.AliossSetting.AliossDomain + "/" + ossSavePath,
	}

	if userID, exists := c.Get("UID"); exists {
		attachment.UserID = userID.(int64)
	}

	var uploadAttachmentTypeMap = map[string]model.AttachmentType{
		"public/image":  model.ATTACHMENT_TYPE_IMAGE,
		"public/avatar": model.ATTACHMENT_TYPE_IMAGE,
		"public/video":  model.ATTACHMENT_TYPE_VIDEO,
		"attachment":    model.ATTACHMENT_TYPE_OTHER,
	}

	attachment.Type = uploadAttachmentTypeMap[uploadType]
	if attachment.Type == model.ATTACHMENT_TYPE_IMAGE {
		var src image.Image
		src, err = imaging.Decode(file)
		if err == nil {
			attachment.ImgWidth, attachment.ImgHeight = GetImageSize(src.Bounds())
		}
	}

	attachment, err = svc.CreateAttachment(attachment)
	if err != nil {
		global.Logger.Errorf("svc.CreateAttachment err: %v", err)
		response.ToErrorResponse(errcode.FileUploadFailed)
	}

	response.ToResponse(attachment)
}

func DownloadAttachmentPrecheck(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c)

	contentID := convert.StrTo(c.Query("id")).MustInt64()
	// 加载content
	content, err := svc.GetPostContentByID(contentID)
	if err != nil {
		global.Logger.Errorf("svc.GetPostContentByID err: %v", err)
		response.ToErrorResponse(errcode.InvalidDownloadReq)
	}
	user, _ := c.Get("USER")
	if content.Type == model.CONTENT_TYPE_CHARGE_ATTACHMENT {
		// 加载post
		post, err := svc.GetPost(content.PostID)
		if err != nil {
			global.Logger.Errorf("svc.GetPost err: %v", err)
			response.ToResponse(gin.H{
				"paid": false,
			})
			return
		}

		// 发布者或管理员免费下载
		if post.UserID == user.(*model.User).ID || user.(*model.User).IsAdmin {
			response.ToResponse(gin.H{
				"paid": true,
			})
			return
		}

		// 检测是否有购买记录
		response.ToResponse(gin.H{
			"paid": svc.CheckPostAttachmentIsPaid(post.ID, user.(*model.User).ID),
		})
		return
	}
	response.ToResponse(gin.H{
		"paid": false,
	})
}

func DownloadAttachment(c *gin.Context) {
	response := app.NewResponse(c)
	svc := service.New(c)

	contentID := convert.StrTo(c.Query("id")).MustInt64()

	// 加载content
	content, err := svc.GetPostContentByID(contentID)
	if err != nil {
		global.Logger.Errorf("svc.GetPostContentByID err: %v", err)
		response.ToErrorResponse(errcode.InvalidDownloadReq)
	}

	// 收费附件
	if content.Type == model.CONTENT_TYPE_CHARGE_ATTACHMENT {
		user, _ := c.Get("USER")

		// 加载post
		post, err := svc.GetPost(content.PostID)
		if err != nil {
			global.Logger.Errorf("svc.GetPost err: %v", err)
			response.ToResponse(gin.H{
				"paid": false,
			})
			return
		}

		paidFlag := false

		// 发布者或管理员免费下载
		if post.UserID == user.(*model.User).ID || user.(*model.User).IsAdmin {
			paidFlag = true
		}

		// 检测是否有购买记录
		if svc.CheckPostAttachmentIsPaid(post.ID, user.(*model.User).ID) {
			paidFlag = true
		}

		if !paidFlag {
			// 未购买，则尝试购买
			err := svc.BuyPostAttachment(&model.Post{
				Model: &model.Model{
					ID: post.ID,
				},
				UserID:          post.UserID,
				AttachmentPrice: post.AttachmentPrice,
			}, user.(*model.User))
			if err != nil {
				global.Logger.Errorf("svc.BuyPostAttachment err: %v", err)
				if err == errcode.InsuffientDownloadMoney {

					response.ToErrorResponse(errcode.InsuffientDownloadMoney)
				} else {

					response.ToErrorResponse(errcode.DownloadExecFail)
				}
				return
			}
		}
	}

	// 开始构造下载地址
	client, err := oss.New(global.AliossSetting.AliossEndpoint, global.AliossSetting.AliossAccessKeyID, global.AliossSetting.AliossAccessKeySecret)
	if err != nil {
		global.Logger.Errorf("oss.New err: %v", err)
		response.ToErrorResponse(errcode.DownloadReqError)
		return
	}

	bucket, err := client.Bucket("paopao-assets")
	if err != nil {
		global.Logger.Errorf("client.Bucket err: %v", err)
		response.ToErrorResponse(errcode.DownloadReqError)
		return
	}

	// 签名
	objectKey := strings.Replace(content.Content, "https://"+global.AliossSetting.AliossDomain+"/", "", -1)
	signedURL, err := bucket.SignURL(objectKey, oss.HTTPGet, 60)
	if err != nil {
		global.Logger.Errorf("client.SignURL err: %v", err)
		response.ToErrorResponse(errcode.DownloadReqError)
		return
	}
	ur, err := url.Parse(signedURL)
	if err != nil {
		global.Logger.Errorf("url.Parse err: %v", err)
		response.ToErrorResponse(errcode.DownloadReqError)
		return
	}
	epath, err := url.PathUnescape(ur.Path)
	if err != nil {
		global.Logger.Errorf("url.PathUnescape err: %v", err)
		response.ToErrorResponse(errcode.DownloadReqError)
		return
	}

	ur.Path = epath
	ur.RawPath = epath
	response.ToResponse(ur.String())
}
