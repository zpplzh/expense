package catalog

import (
	"crypto/rand"
	"math/big"
	"strings"
)

//func RandSessionid() string {
//	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
//var letterRunes = []rune("abcde")

//	b := make([]rune, 35) //length of the random character is 35
//	for i := range b {
//		b[i] = letterRunes[rand.Int(len(letterRunes))]
//	}
//	return string(b)
//}

func RandSessionid() string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"

	leng := 35
	ret := make([]byte, leng)
	for i := 0; i < leng; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
		}
		ret = append(ret, letters[num.Int64()])
	}
	h := strings.Replace(string(ret), "\u0000", "", -1)

	return string(h)
}
