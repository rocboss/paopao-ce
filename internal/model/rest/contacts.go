package rest

type RequestingFriendReq struct {
	UserId    int64  `json:"user_id" binding:"required"`
	Greetings string `json:"greetings" binding:"required"`
}

type AddFriendReq struct {
	UserId int64 `json:"user_id" binding:"required"`
}

type RejectFriendReq struct {
	UserId int64 `json:"user_id" binding:"required"`
}

type DeleteFriendReq struct {
	UserId int64 `json:"user_id"`
}

type ContactItem struct {
	UserId   int64  `json:"user_id"`
	UserName string `json:"username"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
}

type ContactsResp struct {
	Contacts []ContactItem `json:"contacts"`
	Total    int64         `json:"total"`
}
