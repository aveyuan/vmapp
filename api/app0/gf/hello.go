package gf

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HelloReq struct {
	g.Meta `path:"/hello" tags:"Hello" method:"get" summary:"You first hello api"`
	Name string `json:"name" query:"name"`
}
type HelloRes struct {
	g.Meta `mime:"text/html" example:"string"`
	Content string `json:"content"`

}
