package v1

type HelloReq struct {
	Name string `json:"name" form:"name" description:"姓名"`
}

type HelloRes struct {
	Content string `json:"content" description:"返回内容"`
}
