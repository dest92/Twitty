package jwt

import (
	"github.com/dest92/Twitty/models"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

//GenerateJWT generates the JWT 

func GenerateJWT(t models.User) (string, error) {

	myKey := []byte(os.Getenv("PRIVATE_KEY")) //Private key

	//List of privileges

	//Payload is the data that will be encoded in the token
	payload := jwt.MapClaims{
		"email":     t.Email,
		"name":      t.Name,
		"lastName":  t.LastName,
		"birthDate": t.BirthDate,
		"biography": t.Biography,
		"location":  t.Location,
		"webSite":   t.WebSite,
		"_id":       t.ID.Hex(),                            //Convert the id to string
		"exp":       time.Now().Add(time.Hour * 24).Unix(), //Expire in 24 hours //Unix saves fast
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload) //Create the token

	tokenStr, err := token.SignedString(myKey) //Sign the token

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}
