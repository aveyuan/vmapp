package sm

import (
	"crypto/hmac"
	"encoding/hex"

	"github.com/tjfoc/gmsm/sm3"
)

func Sm3(sm3str string, secret string) string {
	//key 传入具体的key，根据key编码转换成[]byte
	key := []byte(secret)
	//具体需要加密的数据
	message := []byte(sm3str)
	// 创建HMAC-SM3哈希算法对象
	h := hmac.New(sm3.New, key)
	// 将消息添加到哈希算法中
	h.Write(message)
	// 计算哈希值
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}