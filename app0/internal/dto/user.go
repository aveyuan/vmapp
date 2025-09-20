package dto

// GetUserReq
type GetUserReq struct {
	UID int64 `form:"uid" json:"uid" validate:"required"`
}

// DeleteUserReq
type DeleteUserReq struct {
	UID int64 `form:"uid" json:"uid" validate:"required"`
}

type CreateUserReq struct {
	Username   string `json:"username" form:"username" validate:"required"`
	Avatar     string `json:"avatar" form:"avatar"`
	Email      string `json:"email" form:"email"`
	Url        string `json:"url" form:"url"`
	Nickname   string `json:"nick_name" form:"nick_name"`
	RoleCodes  string `json:"role_codes" form:"role_codes"`
	IsAdmin    int8   `json:"is_admin" form:"is_admin"`
	Password   string `json:"password" form:"password" validate:"required"`
	Repassword string `json:"repassword" form:"repassword" validate:"required"`
	Status     int8   `json:"status" form:"status"`
}

type UpdateUserReq struct {
	Uid         int64  `json:"uid" form:"uid" validate:"required"`
	Avatar      string `json:"avatar" form:"avatar"`
	Email       string `json:"email" form:"email"`
	Url         string `json:"url" form:"url"`
	Nickname    string `json:"nick_name" form:"nick_name"`
	RoleCodes   string `json:"role_codes" form:"role_codes"`
	IsAdmin     int8   `json:"is_admin" form:"is_admin"`
	Status      int8   `json:"status" form:"status"`
	CheckStatus int8   `json:"check_status" form:"check_status"`
}

type ListUserReq struct {
	Uid      int64  `json:"uid" form:"uid"`
	Username string `json:"username" form:"username"`
	Nickname string `json:"nick_name" form:"nick_name"`
	Mail     string `json:"mail" form:"mail"`
	Page     int    `json:"page" form:"page"`
	Limit    int    `json:"limit" form:"limit"`
}

type RepassUserReq struct {
	Uid        int64  `json:"uid" form:"uid" validate:"required"`
	Password   string `json:"password" form:"password" validate:"required"`
	RePassword string `json:"repassword" form:"repassword" validate:"required"`
}

type UpdateUserPassReq struct {
	OldPassword string `json:"old_password" form:"old_password" validate:"required"`
	Password    string `json:"password" form:"password" validate:"required"`
	RePassword  string `json:"re_password" form:"re_password" validate:"required"`
}

type UpdateUserProfileReq struct {
	Avatar      string `json:"avatar" form:"avatar"`
	Nickname    string `json:"nickname" form:"nickname" validate:"required"`
	Description string `json:"description" form:"description"`
}
