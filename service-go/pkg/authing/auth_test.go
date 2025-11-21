package authing

import (
	"encoding/base32"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"math"
	"testing"
	"time"
)

func TestAuth(t *testing.T) {
	s, _ := base32.StdEncoding.DecodeString("WFJRVSFVU3F7KSKJCDVXTN56ANBU7OX2")

	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      "github.com/ncraft-io/armory/auth",
		AccountName: "34xVOKC0zN37k21ihLsOjVS7bQS",
		Secret:      s,
	})

	counter := int64(math.Floor(float64(time.Now().Unix()) / float64(30)))

	i := -3
	for count := counter - 3; count < counter+3; count++ {
		code, _ := totp.GenerateCodeCustom(key.Secret(), time.Unix(count*30, 0), totp.ValidateOpts{
			Period:    30,
			Skew:      1,
			Digits:    otp.DigitsSix,
			Algorithm: otp.AlgorithmSHA1,
		})
		println(i, " :code: ", code, "\n")
		i++
	}

	code, _ := totp.GenerateCode(key.Secret(), time.Now())
	println("\n", code)

	//code, _ := totp.GenerateCode(key.Secret(), time.Now())
	//time.Sleep(35 * time.Second)
	//valid, _ := totp.ValidateCustom(
	//	"365725",
	//	key.Secret(),
	//	time.Now().UTC(),
	//	totp.ValidateOpts{
	//		Period:    30,
	//		Skew:      2,
	//		Digits:    otp.DigitsSix,
	//		Algorithm: otp.AlgorithmSHA1,
	//	})
	//
	//assert.True(t, valid)
}
