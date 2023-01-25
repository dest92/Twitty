package database

import "golang.org/x/crypto/bcrypt"

// EncryptPassword is a function to encrypt the password

func EncryptPassword(pass string) (string, error) {
	cost := 8 //The cost is the amount of encryptions
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)
	return string(bytes), err
}