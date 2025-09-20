package dto

type LoginReq struct {
	Username string `json:"username" form:"username"  validate:"required"`
	Password string `json:"password" form:"password"  validate:"required"`
	Id       string `json:"id" form:"id"`
	Captcha  string `json:"captcha" form:"captcha" validate:"required"`
}
type SSoLoginReq struct {
	Username string `form:"user_name" validate:""`
	Date     int64  `form:"date" validate:""`
	UserID   int64  `form:"user_id" validate:""`
	Sign     string `form:"sign" validate:""`
}

type LoginResp struct {
	Token string `json:"token"` //token
	Exp   int    `json:"exp"`   //过期时间
}

type RegisterReq struct {
	Username   string `json:"username" form:"username"  validate:"required"`
	Email      string `json:"email" form:"email" validate:"required"`
	Code       string `json:"code" form:"code" validate:"required"`
	Password   string `json:"password" form:"password"  validate:"required"`
	RePassword string `json:"repassword" form:"repassword" validate:"required"`
	Agreement  int    `json:"agreement" form:"agreement" validate:"required"`
}

type ForgetReq struct {
	Email      string `json:"email" form:"email" validate:"required"`
	Code       string `json:"code" form:"code" validate:"required"`
	Password   string `json:"password" form:"password"  validate:"required"`
	RePassword string `json:"repassword" form:"repassword" validate:"required"`
}

type GetCodeReq struct {
	Email   string `json:"email" form:"email"`
	Phone   string `json:"phone" form:"phone"`
	Type    int    `json:"type" form:"type"`
	Id      string `json:"id" form:"id"`
	Captcha string `json:"captcha" form:"captcha" validate:"required"`
}
