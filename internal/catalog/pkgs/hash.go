package pkgs

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
