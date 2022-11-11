package localoss

import (
	"github.com/rocboss/paopao-ce/pkg/xerror"
)

var (
	errFileUploadFailed = xerror.NewError(10200, "文件上传失败")
	errFileInvalidExt   = xerror.NewError(10201, "文件类型不合法")
	errFileInvalidSize  = xerror.NewError(10202, "文件大小超限")
)
