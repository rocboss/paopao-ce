package jinzhu

import (
	"github.com/rocboss/paopao-ce/internal/core"
	"github.com/rocboss/paopao-ce/internal/model"
	"gorm.io/gorm"
)

var (
	_ core.MessageService = (*messageServant)(nil)
)

type messageServant struct {
	db *gorm.DB
}

func newMessageService(db *gorm.DB) core.MessageService {
	return &messageServant{
		db: db,
	}
}

func (d *messageServant) CreateMessage(msg *model.Message) (*model.Message, error) {
	return msg.Create(d.db)
}

func (d *messageServant) GetUnreadCount(userID int64) (int64, error) {
	return (&model.Message{}).Count(d.db, &model.ConditionsT{
		"receiver_user_id": userID,
		"is_read":          model.MsgStatusUnread,
	})
}

func (d *messageServant) GetMessageByID(id int64) (*model.Message, error) {
	return (&model.Message{
		Model: &model.Model{
			ID: id,
		},
	}).Get(d.db)
}

func (d *messageServant) ReadMessage(message *model.Message) error {
	message.IsRead = 1
	return message.Update(d.db)
}

func (d *messageServant) GetMessages(conditions *model.ConditionsT, offset, limit int) ([]*model.MessageFormated, error) {
	messages, err := (&model.Message{}).List(d.db, conditions, offset, limit)
	if err != nil {
		return nil, err
	}

	mfs := []*model.MessageFormated{}
	for _, message := range messages {
		mf := message.Format()
		mfs = append(mfs, mf)
	}

	return mfs, nil
}

func (d *messageServant) GetMessageCount(conditions *model.ConditionsT) (int64, error) {
	return (&model.Message{}).Count(d.db, conditions)
}
