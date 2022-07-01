package core

import (
	"github.com/rocboss/paopao-ce/internal/model"
)

// MessageService 消息服务
type MessageService interface {
	CreateMessage(msg *model.Message) (*model.Message, error)
	GetUnreadCount(userID int64) (int64, error)
	GetMessageByID(id int64) (*model.Message, error)
	ReadMessage(message *model.Message) error
	GetMessages(conditions *model.ConditionsT, offset, limit int) ([]*model.MessageFormated, error)
	GetMessageCount(conditions *model.ConditionsT) (int64, error)
}
