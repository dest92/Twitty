package middlew

import (
	"github.com/dest92/Twitty/database"
	"net/http"
)

// CheckDB is a middleware that checks the connection to the database
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if database.CheckConnection() == 0 {
			http.Error(w, "Lost connection to database", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
