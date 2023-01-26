package database

import (
	"github.com/dest92/Twitty/models"
	"golang.org/x/crypto/bcrypt"
)

// TryLogin compares the email and password with the database
func TryLogin(email string, password string) (models.User, bool) {
	user, found, _ := UserExists(email)

	if !found {
		return user, false
	}

	//Compare password

	passwordBytes := []byte(password) //Slice of bytes
	passwordDB := []byte(user.Password)

	//Bcrypt works with slices of bytes
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return user, false
	}

	return user, true
}
