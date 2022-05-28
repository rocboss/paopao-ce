package model

import "gorm.io/gorm"

type MessageT int8

const (
	MESSAGE_POST MessageT = iota + 1
	MESSAGE_COMMENT
	MESSAGE_REPLY
	MESSAGE_WHISPER
)
const MESSAGE_SYSTEM MessageT = 99

const MSG_UNREAD = 0
const MSG_READED = 1

type Message struct {
	*Model
	SenderUserID   int64    `json:"sender_user_id"`
	ReceiverUserID int64    `json:"receiver_user_id"`
	Type           MessageT `json:"type"`
	Breif          string   `json:"breif"`
	Content        string   `json:"content"`
	PostID         int64    `json:"post_id"`
	CommentID      int64    `json:"comment_id"`
	ReplyID        int64    `json:"reply_id"`
	IsRead         int8     `json:"is_read"`
}

type MessageFormated struct {
	ID             int64         `json:"id"`
	SenderUserID   int64         `json:"sender_user_id"`
	SenderUser     *UserFormated `json:"sender_user"`
	ReceiverUserID int64         `json:"receiver_user_id"`
	Type           MessageT      `json:"type"`
	Breif          string        `json:"breif"`
	Content        string        `json:"content"`
	PostID         int64         `json:"post_id"`
	Post           *PostFormated `json:"post"`
	CommentID      int64         `json:"comment_id"`
	Comment        *Comment      `json:"comment"`
	ReplyID        int64         `json:"reply_id"`
	Reply          *CommentReply `json:"reply"`
	IsRead         int8          `json:"is_read"`
	CreatedOn      int64         `json:"created_on"`
	ModifiedOn     int64         `json:"modified_on"`
}

func (m *Message) Format() *MessageFormated {
	if m.Model == nil || m.Model.ID == 0 {
		return nil
	}
	mf := &MessageFormated{
		ID:             m.ID,
		SenderUserID:   m.SenderUserID,
		SenderUser:     &UserFormated{},
		ReceiverUserID: m.ReceiverUserID,
		Type:           m.Type,
		Breif:          m.Breif,
		Content:        m.Content,
		PostID:         m.PostID,
		Post:           &PostFormated{},
		CommentID:      m.CommentID,
		Comment:        &Comment{},
		ReplyID:        m.ReplyID,
		Reply:          &CommentReply{},
		IsRead:         m.IsRead,
		CreatedOn:      m.CreatedOn,
		ModifiedOn:     m.ModifiedOn,
	}

	return mf
}

func (m *Message) Create(db *gorm.DB) (*Message, error) {
	err := db.Create(&m).Error

	return m, err
}

func (m *Message) Update(db *gorm.DB) error {
	return db.Model(&Message{}).Where("id = ? AND is_del = ?", m.Model.ID, 0).Save(m).Error
}

func (m *Message) Get(db *gorm.DB) (*Message, error) {
	var message Message
	if m.Model != nil && m.ID > 0 {
		db = db.Where("id = ? AND is_del = ?", m.ID, 0)
	}
	if m.ReceiverUserID > 0 {
		db = db.Where("receiver_user_id = ?", m.ReceiverUserID)
	}

	err := db.First(&message).Error
	if err != nil {
		return &message, err
	}

	return &message, nil
}

func (c *Message) List(db *gorm.DB, conditions *ConditionsT, offset, limit int) ([]*Message, error) {
	var messages []*Message
	var err error
	if offset >= 0 && limit > 0 {
		db = db.Offset(offset).Limit(limit)
	}

	for k, v := range *conditions {
		if k == "ORDER" {
			db = db.Order(v)
		} else {
			db = db.Where(k, v)
		}
	}

	if err = db.Where("is_del = ?", 0).Find(&messages).Error; err != nil {
		return nil, err
	}

	return messages, nil
}

func (m *Message) Count(db *gorm.DB, conditions *ConditionsT) (int64, error) {
	var count int64

	for k, v := range *conditions {
		if k != "ORDER" {
			db = db.Where(k, v)
		}
	}
	if err := db.Model(m).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
