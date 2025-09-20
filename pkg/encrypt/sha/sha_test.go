package sha

import (
	"log"
	"testing"
)

func TestSha256(t *testing.T) {
	s := Sha256("abc")
	log.Print(s)
	log.Print(len(s))
}