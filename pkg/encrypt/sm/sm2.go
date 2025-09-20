package sm

import (
	crand "crypto/rand"
	"encoding/hex"
	"math/big"
	"unsafe"

	"github.com/tjfoc/gmsm/sm2"
)

// GenPublicKey 生成公钥
func GenPublicKey(publicKey string) (*sm2.PublicKey, error) {
	d, err := hex.DecodeString(publicKey)
	if err != nil {
		return nil, err
	}
	c := sm2.P256Sm2()

	return &sm2.PublicKey{Curve: c, X: new(big.Int).SetBytes(d[1:33]), Y: new(big.Int).SetBytes(d[33:])}, nil
}

// GenPrivatekey 生成私钥
func GenPrivatekey(privateKey string, publicKey *sm2.PublicKey) (*sm2.PrivateKey, error) {
	d, err := hex.DecodeString(privateKey)
	if err != nil {
		return nil, err
	}

	return &sm2.PrivateKey{PublicKey: *publicKey, D: new(big.Int).SetBytes(d)}, nil
}

// EncSM2 SM2加密
func EncSM2(msg string, pubkey *sm2.PublicKey) (string, error) {
	res, err := sm2.Encrypt(pubkey, []byte(msg), crand.Reader, sm2.C1C3C2)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(res), nil
}

// DecSM2 SM2解密
func DecSM2(msg string, prikey *sm2.PrivateKey) (string, error) {
	// 做兼容
	if msg[:2] != "04" {
		msg = "04" + msg
	}

	b, err := hex.DecodeString(msg)
	if err != nil {
		return "", err
	}
	enc, err := sm2.Decrypt(prikey, b, sm2.C1C3C2)
	if err != nil {
		return "", err
	}
	return *(*string)(unsafe.Pointer(&enc)), nil
}
