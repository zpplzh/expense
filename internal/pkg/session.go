package pkg

import "math/rand"

func RandSessionid() string {

	//var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	var letterRunes = []rune("ab")

	b := make([]rune, 1) //length of the random character is 35
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
