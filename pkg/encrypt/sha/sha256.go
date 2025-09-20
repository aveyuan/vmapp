package sha

import (
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(str string) string {
	data := []byte(str)
	hash := sha256.New()
	hash.Write(data)
	hashValue := hash.Sum(nil)
	return hex.EncodeToString(hashValue)
}

func Sha256Byte(data []byte) string {
	hash := sha256.New()
	hash.Write(data)
	hashValue := hash.Sum(nil)
	return hex.EncodeToString(hashValue)

}
