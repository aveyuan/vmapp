package sso

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

// SSO 标准格式
type SSO struct {
	URL      string
	Username string
	Date     int64
	UserID   int64
}

// GetSignature 签名生成
func (c *SSO) GetSignature(token string) string {
	toSing := fmt.Sprintf("%v%v%v", c.Username, c.UserID, c.Date)
	byteSing := []byte(toSing)
	bas := base64.StdEncoding.EncodeToString(byteSing)
	mac := hmac.New(sha1.New, []byte(token))
	mac.Write([]byte(bas))
	ssoEncode := fmt.Sprintf("%x", mac.Sum(nil))
	return string(ssoEncode)
}
