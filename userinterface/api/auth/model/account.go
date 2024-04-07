package model

type PostAccountRequest struct {
	Username string `json:"username" binding:"required"` // 帳號
	Password string `json:"password" binding:"required"` // 密碼
}

type PostAccountResponse struct {
	UID string `json:"uid"` // 使用者UID
}

type PutAccountRequestUri struct {
	UID string `uri:"uid" binding:"required"` // 使用者UID
}

type PutAccountRequest struct {
	Status       uint `json:"status" binding:"required"`       // 狀態
	IsEnabledOTP bool `json:"isEnabledOTP" binding:"required"` // 是否啟用OTP
}

type PutAccountResponse struct {
}

type PutAccountPasswordRequestUri struct {
	UID string `uri:"uid" binding:"required"` // 使用者UID
}

type PutAccountPasswordRequest struct {
	OldPassword string `json:"oldPassword" binding:"required"` // 舊密碼
	NewPassword string `json:"newPassword" binding:"required"` // 新密碼
}

type PutAccountPasswordResponse struct {
}

type GetAccountPermissionRequestUri struct {
	UID string `uri:"uid" binding:"required"` // 使用者UID
}

type GetAccountPermissionResponse struct {
	PermissionUIDs []string `json:"permissionUIDs"` // 權限UID
}

type GetPermissionRequest struct {
}

type GetPermissionResponse struct {
	Permissions []Permission `json:"permissions"` // 權限
}

type Permission struct {
	UID   string `json:"uid"`   // 權限UID
	Group string `json:"group"` // 群組
	Name  string `json:"name"`  // 名稱
}
