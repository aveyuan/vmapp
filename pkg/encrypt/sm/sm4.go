package sm

import (
	"github.com/tjfoc/gmsm/sm4"
)

type sm4Encrypt struct {
	key []byte
}

func NewAes(key []byte) *sm4Encrypt {
	return &sm4Encrypt{
		key: key,
	}
}

// DecSM4 解码
func (t *sm4Encrypt) DecSM4(encbyte []byte) ([]byte, error) {
	ecbDec, err := sm4.Sm4Cbc(t.key, encbyte, false)
	if err != nil {
		return nil, err
	}

	return ecbDec, nil
}

// EncSM4 编码
func (t *sm4Encrypt) EncSM4(msg []byte) ([]byte, error) {
	ecbMsg, err := sm4.Sm4Cbc(t.key, msg, true)
	if err != nil {
		return nil, err
	}

	return ecbMsg, nil
}
