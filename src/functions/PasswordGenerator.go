package functions

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

//GenerateMD5 transforma a senha para hash md5
func GenerateMD5(pass string) string {
	hash := md5.Sum([]byte(pass))
	return hex.EncodeToString(hash[:])
}

// GeneratePassword faz o hash da senha
func GeneratePassword(passwordSend string) (string, error) {
	password := []byte(passwordSend)

	// Hashing the password with the default cost of 10
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePasswords Faz a comparação da senha
func ComparePasswords(passwordSend string, hashedPassword []byte) bool {
	password := []byte(passwordSend)

	// Comparing the password with the hash
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	fmt.Println(err) // nil means it is a match
	if err != nil {
		return true
	} else {
		return false
	}
}
