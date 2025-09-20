package sm

import (
	"log"
	"strconv"
	"testing"
)

func TestSm2(t *testing.T) {
	const publicKey = "04b80e9cd4bc497b80367b3c3fc9f3159f90ec88e3725fea97a3b6a15b783464325fd3f4332795a52a318e01b93d08a2acc957ad5e6a60902341d8a77741c7a4b4" // 公钥
	const privateKey = "00dd70063c6d25356f6bd0044180e5f1d0d360fdc5c554b00090877a5a01339b94"
	msg := "123456"
	pubkey, err := GenPublicKey(publicKey)
	if err != nil {
		t.Fatal(err)
	}
	prikey, err := GenPrivatekey(privateKey, pubkey)
	if err != nil {
		t.Fatal(err)
	}
	enc, err := EncSM2(msg, pubkey)
	if err != nil {
		t.Fatal(err)
	}
	log.Print(enc)
	dec, err := DecSM2(enc, prikey)
	if err != nil {
		t.Fatal(err)
	}
	log.Print(dec)
}

func BenchmarkEncsm2(b *testing.B) {
	const publicKey = "04b80e9cd4bc497b80367b3c3fc9f3159f90ec88e3725fea97a3b6a15b783464325fd3f4332795a52a318e01b93d08a2acc957ad5e6a60902341d8a77741c7a4b4" // 公钥
	msg := "123456"
	pubkey, err := GenPublicKey(publicKey)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		EncSM2(msg+strconv.Itoa(i), pubkey)
	}
}
func BenchmarkDecSM2(b *testing.B) {
	const publicKey = "04b80e9cd4bc497b80367b3c3fc9f3159f90ec88e3725fea97a3b6a15b783464325fd3f4332795a52a318e01b93d08a2acc957ad5e6a60902341d8a77741c7a4b4" // 公钥
	const privateKey = "00dd70063c6d25356f6bd0044180e5f1d0d360fdc5c554b00090877a5a01339b94"
	pubkey, err := GenPublicKey(publicKey)
	if err != nil {
		b.Fatal(err)
	}
	prikey, err := GenPrivatekey(privateKey, pubkey)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		_, err := DecSM2("054e20774cd5826bc38a9e424cbfa89cd8cdf6b2ddfa365ca35c864a87f0ef77de3929a4672f1d01b2a65fd66c0f3672c09b7f0c6a7f30b6df7f5c0736f43b754a64f5200616031f38fee70cac959a08f18bd77be7cdfdba49ce11c4da5aec42071a04", prikey)
		if err != nil {
			{
				b.Fatal(err)
			}
		}
	}
}


func TestSm3(t *testing.T) {
	res := Sm3("123", "1234")
	log.Print(res)
}