package middlew

import (
	"github.com/dest92/Twitty/routers" //Executes the function in the routers package
	"net/http"
)

/*The CheckJWT function is an http middleware that is responsible for validating the JWT token sent in the request header.
It uses the ProcessToken function to get the token from the request header and verify it.
If the token is valid, the next.ServeHTTP function is called to continue executing the request.
If the token is invalid or has some error, a "Bad Request" HTTP status error is sent and execution stops.*/

func CheckJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization")) //Get the token from the header

		if err != nil {
			http.Error(w, "Error in token: "+err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}
