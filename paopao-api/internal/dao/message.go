package dao

import "github.com/rocboss/paopao-api/internal/model"

func (d *Dao) CreateMessage(msg *model.Message) (*model.Message, error) {
	return msg.Create(d.engine)
}

func (d *Dao) GetUnreadCount(userID int64) (int64, error) {
	return (&model.Message{}).Count(d.engine, &model.ConditionsT{
		"receiver_user_id": userID,
		"is_read":          model.MSG_UNREAD,
	})
}

func (d *Dao) GetMessageByID(id int64) (*model.Message, error) {
	return (&model.Message{
		Model: &model.Model{
			ID: id,
		},
	}).Get(d.engine)
}

func (d *Dao) ReadMessage(message *model.Message) error {
	message.IsRead = 1
	return message.Update(d.engine)
}

func (d *Dao) GetMessages(conditions *model.ConditionsT, offset, limit int) ([]*model.MessageFormated, error) {
	messages, err := (&model.Message{}).List(d.engine, conditions, offset, limit)
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

func (d *Dao) GetMessageCount(conditions *model.ConditionsT) (int64, error) {
	return (&model.Message{}).Count(d.engine, conditions)
}
