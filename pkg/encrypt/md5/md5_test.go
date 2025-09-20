package md5

import (
	"log"
	"testing"
)

func TestMD5(t *testing.T) {
	res := Md5("123")
	log.Print(res)
}
