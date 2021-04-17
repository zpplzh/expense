package catalog

import (
	"crypto/rand"
	"log"
	"math/big"
	"regexp"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type (
	Hash struct{}
)

func (h *Hash) HashandSalt(pwd string) string {

	ha, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	k := string(ha)

	return k
}

func (h *Hash) CheckPass(pwd string, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	if err != nil {
		log.Println(err)
	}

	return err == nil
}

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

func checkEmail(email string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	if len(email) < 3 && len(email) > 254 {
		return false
	}

	return emailRegex.MatchString(email)
}

func checkInput(in string) bool {
	var inpstr = regexp.MustCompile("[a-zA-Z0-9]")

	return inpstr.MatchString(in)
}
