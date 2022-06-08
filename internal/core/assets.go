package core

type AttachmentCheckService interface {
	Check(uri string) error
}
