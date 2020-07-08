package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		catchtoken := r.Header.Get("Authorization")
		if len(catchtoken) == 0 {
			catchtoken = ""
		} else {
			tokenVal := strings.Split(catchtoken, "Bearer ")
			catchtoken = strings.Trim(tokenVal[1], "")
			fmt.Println(catchtoken)
			parsetoken, err := jwt.Parse(catchtoken, func(token *jwt.Token) (interface{}, error) {
				if jwt.GetSigningMethod("HS256") != token.Method {
					return  nil,fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return []byte("secret"), nil
			})
			if parsetoken != nil  && err == nil {
				next.ServeHTTP(w, r)
			}else {
				// Call the next handler, which can be another middleware in the chain, or the final handler.
				desc := "token tidak valid "
				w.Write([]byte(desc))
			}
		}


	})
}



