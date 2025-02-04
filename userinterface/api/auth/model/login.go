package model

type PostLoginRequest struct {
	Username string `json:"username" binding:"required"` // 帳號
	Password string `json:"password" binding:"required"` // 密碼
}

type PostLoginResponse struct {
	AuthToken string `json:"authToken"` // 登入Token
}
