package core

type AttachmentCheckService interface {
	Check(cUrl string) error
}
