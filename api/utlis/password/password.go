package password

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, 12)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}
