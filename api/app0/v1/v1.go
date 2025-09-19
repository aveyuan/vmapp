package v1

// @Title 获取Hello
// @Param  .  query  HelloReq  false  "请求参数"
// @Success  200  object  HelloRes  "返回结果"
// @Tag users
// @Route /api/v1/hello [get]
type HelloReq struct {
	Name string `json:"name" form:"name" description:"姓名"`
}

type HelloRes struct {
	Content string `json:"content" description:"返回内容"`
}

type Page struct {
	Total       int `json:"total"`
	Count       int `json:"count"`
	PerPage     int `json:"perPage"`
	CurrentPage int `json:"currentPage"`
	TotalPages  int `json:"totalPages"`
}

type LoginReq struct {
	Username  string `json:"username" form:"username" description:"用户名"`
	Password  string `json:"password" form:"password" description:"密码"`
	Captcha   string `json:"captcha" form:"captcha" description:"验证码"`
	CaptchaID string `json:"captchaId" form:"captchaId" description:"验证码ID"`
}

type LoginRes struct {
	Token  string `json:"token" description:"token"`
	Expire int64  `json:"expire" description:"过期时间"`
}
