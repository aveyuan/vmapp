package aes

import (
	"fmt"
	"testing"

)

func TestAes(t *testing.T) {
	encipher := NewAes([]byte("hgfedcba87654321"))
	source := []byte("test")
	encode, err := encipher.Encrypt(source)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(encode)
	decode, err := encipher.Decrypt(encode)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(decode))
}
