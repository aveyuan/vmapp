package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

type PubKey string
type PriKey string

func newPubKey(str string) PubKey {
	return PubKey(fmt.Sprintf(`
-----BEGIN PUBLIC KEY-----
%v
-----END PUBLIC KEY-----`, str))
}

func newPriKey(str string) PriKey {
	return PriKey(fmt.Sprintf(`
-----BEGIN PRIVATE KEY-----
%v
-----END PRIVATE KEY-----`, str))
}

type Pcks string

const Pcks1 Pcks = "Pcks1"
const Pcks8 Pcks = "Pcks8"

type Rsa struct {
	pub  *rsa.PublicKey
	pri  *rsa.PrivateKey
	pcks Pcks
}

type RsaOption func(o *Rsa)

func WithPcks(pcks Pcks) RsaOption {
	return func(o *Rsa) { o.pcks = pcks }
}

func NewCecRSA(pubKey string, priKey string, opt ...RsaOption) (*Rsa, error) {
	cecrsa := &Rsa{}
	for _, v := range opt {
		v(cecrsa)
	}

	if cecrsa.pcks == "" {
		cecrsa.pcks = Pcks1
	}

	if pubKey != "" {
		pub := newPubKey(pubKey)
		block, _ := pem.Decode([]byte(pub))
		if block == nil {
			return nil, errors.New("priv key error")
		}
		if cecrsa.pcks == Pcks1 {
			pubv, err := x509.ParsePKCS1PublicKey(block.Bytes)
			if err != nil {
				return nil, err
			}
			cecrsa.pub = pubv

		}

		if cecrsa.pcks == Pcks8 {
			pubv, err := x509.ParsePKIXPublicKey(block.Bytes)
			if err != nil {
				return nil, err
			}
			if p, ok := pubv.(*rsa.PublicKey); ok {
				cecrsa.pub = p
			}
		}

	}

	if priKey != "" {
		pri := newPriKey(priKey)
		block, _ := pem.Decode([]byte(pri))
		if block == nil {
			return nil, errors.New("priv key error")
		}
		if cecrsa.pcks == Pcks1 {
			priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
			if err != nil {
				return nil, err
			}
			cecrsa.pri = priv
		}

		if cecrsa.pcks == Pcks8 {
			priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
			if err != nil {
				return nil, err
			}
			if p, ok := priv.(*rsa.PrivateKey); ok {
				cecrsa.pri = p
			}
		}

	}

	return cecrsa, nil

}

func (t *Rsa) Encrypt(str string) (string, error) {
	if t.pub == nil {
		return "", errors.New("公钥为空")
	}
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, t.pub, []byte(str))
	if err != nil {
		return "", err
	}
	// base64
	return base64.StdEncoding.EncodeToString(ciphertext), nil

}

func (t *Rsa) Decrypt(str string) (string, error) {
	if t.pri == nil {
		return "", errors.New("私钥为空")
	}
	// decode
	dstr, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return "", err
	}

	plainText, err := rsa.DecryptPKCS1v15(rand.Reader, t.pri, dstr)
	if err != nil {
		return "", err
	}
	return string(plainText), nil
}

// RsaSignSha1 加签
func RsaSignSha1(ciphertext []byte, privKey PriKey) (string, error) {
	block, _ := pem.Decode([]byte(privKey))
	if block == nil {
		return "", errors.New("priv key error")
	}
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	privateKey, ok := priv.(*rsa.PrivateKey)
	if !ok {
		return "", errors.New("非法私钥文件，检查私钥文件")
	}

	h := crypto.Hash.New(crypto.SHA1)
	_, err = h.Write(ciphertext)
	if err != nil {
		return "", err
	}
	hashed := h.Sum(nil)
	signatureByte, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA1, hashed)
	if err != nil {
		return "", err
	}
	//一般签名需要转一下base64
	return base64.StdEncoding.EncodeToString(signatureByte), nil
}

// 验签
func RsaVerifySha1(pubkey PubKey, sign string, data []byte) error {
	block, _ := pem.Decode([]byte(pubkey))
	if block == nil {
		return errors.New("priv key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	pub, ok := pubInterface.(*rsa.PublicKey)
	if !ok {
		return errors.New("非法公钥文件，检查公钥文件")
	}

	// 解密签名
	s, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}

	h := crypto.Hash.New(crypto.SHA1)
	_, err = h.Write(data)
	if err != nil {
		return err
	}
	hashed := h.Sum(nil)

	return rsa.VerifyPKCS1v15(pub, crypto.SHA1, hashed, s)

}
