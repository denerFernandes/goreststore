package middlewares

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/denerFernandes/goreststore/models"
	"github.com/denerFernandes/goreststore/utils"
	"github.com/dgrijalva/jwt-go"
)

// isAuth - validates the token - gives access to protected endpoints
func IsAuth(next http.HandlerFunc) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var errObj models.Error
		authHeader := r.Header.Get("Authorization")
		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) == 2 {
			authToken := bearerToken[1]
			token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte(os.Getenv("APP_SECRET")), nil
			})

			if err != nil {
				errObj.Message = err.Error()
				utils.RespondWithError(w, http.StatusUnauthorized, errObj)
				return
			}

			if token.Valid {
				// invoke function getting called on
				// in this case the protected endpoint function handler
				next.ServeHTTP(w, r)
			} else {
				errObj.Message = err.Error()
				utils.RespondWithError(w, http.StatusUnauthorized, errObj)
				return
			}
		} else {
			errObj.Message = "Invalid Token"
			utils.RespondWithError(w, http.StatusUnauthorized, errObj)
			return
		}
	})
}
