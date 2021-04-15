package pkg

import (
	"log"

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
