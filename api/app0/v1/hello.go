package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
}
type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Content string `json:"content"`

}

type GetelloReq struct {
	g.Meta `path:"/get" tags:"Hello" method:"get" summary:"You first hello api"`
}
type GetHelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Content string `json:"content"`

}