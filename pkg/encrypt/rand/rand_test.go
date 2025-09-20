package rand

import (
	"fmt"
	"log"
	"testing"
	"time"
)

func TestApproach8(t *testing.T) {
	for i := 0; i < 64; i++ {
		fmt.Println(RandStr(32))
	}
}
func TestRandInt(t *testing.T) {
	log.Print(time.Now().UnixNano())
	log.Print(time.Now().UnixMicro())
	log.Print(time.Now().UnixMilli())
	log.Print(time.Now().Unix())
	for i := 0; i < 64; i++ {
		fmt.Println(RandInt(6))
		fmt.Println(RandStr16(32))
	}
}

func BenchmarkApproach8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := RandStr(10)
		log.Printf("%s,====%v", s, len(s))
	}
}
