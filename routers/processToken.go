package routers

import (
	"errors"
	"github.com/dest92/Twitty/database"
	"github.com/dest92/Twitty/models"
	jwt "github.com/dgrijalva/jwt-go"
	"strings"
)

var Email string  //Global variable to store the email of the user who made the request
var UserID string //Global variable to store the ID of the user who made the request

// Extract values from the token
func ProcessToken(tk string) (*models.Claim, bool, string, error) {

	const myKey = "Twitty" //Constant key for the token
	myKeyByte := []byte(myKey)
	claims := &models.Claim{} //Check token needs a pointer to a Claim struct

	splitToken := strings.Split(tk, "Bearer") //Split the standard "bearer" from the token to get the value

	//If the token is not in the correct format, return an error

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("invalid format token")

	}

	tk = strings.TrimSpace(splitToken[1]) //Remove spaces from the token

	//Parse the token and check if it is valid

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) { //Parse the token and check if it is valid
		return myKeyByte, nil
	})

	//If the token is valid, check if the user exists in the database
	if err == nil {
		_, found, _ := database.UserExists(claims.Email)
		if found {
			Email = claims.Email
			UserID = claims.ID.Hex()
			return claims, found, UserID, nil
		}
		//If the user does not exist, return an error
		return claims, false, string(""), errors.New("user not found")
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}

	return claims, false, string(""), err
}
