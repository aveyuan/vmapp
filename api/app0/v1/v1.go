package v1

type HelloReq struct {
	Name string `json:"name" form:"name" description:"name"`
}

type HelloRes struct {
	Content string `json:"content" description:"hello world"`
}
