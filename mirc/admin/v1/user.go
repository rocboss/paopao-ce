package v1

import (
	. "github.com/alimy/mir/v5"
)

type AgentInfo struct {
	Platform  string `json:"platform"`
	UserAgent string `json:"user_agent"`
}

type ServerInfo struct {
	ApiVer string `json:"api_ver"`
}

type UserInfo struct {
	Name string `json:"name"`
}

type LoginReq struct {
	AgentInfo AgentInfo `json:"agent_info"`
	Name      string    `json:"name"`
	Passwd    string    `json:"passwd"`
}

type LoginResp struct {
	UserInfo
	ServerInfo ServerInfo `json:"server_info"`
	JwtToken   string     `json:"jwt_token"`
}

type User struct {
	Schema `mir:"m/v1,chain"`

	Login  func(Post, LoginReq) LoginResp `mir:"user/login"`
	Logout func(Post)                     `mir:"user/logout"`
}
