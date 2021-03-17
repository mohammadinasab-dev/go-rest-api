package jwt

import (
	"go-rest-api/data"
	Log "go-rest-api/logwrapper"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey []byte

func JWTSetter(jwt string) {
	jwtKey = []byte(jwt)
}

//Claims is a claime structure for generating jwt token
type Claims struct {
	authorized bool
	Email      string
	jwt.StandardClaims
}

//SetJwtToken generate and sets the Authorization token
func SetJwtToken(w http.ResponseWriter, login data.User) bool {

	expTime := time.Now().Add(10 * time.Minute)

	claims := &Claims{
		authorized: true,
		Email:      login.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	validToken, err := token.SignedString(jwtKey)

	if err != nil {
		Log.STDLog.Error(err)
		return false
	}

	bearer := "Bearer " + validToken
	w.Header().Set("Authorization", bearer)
	return true
}

//IsAuthorized checks the Authorization token
func IsAuthorized(r *http.Request) (data.User, bool) {

	if r.Header["Authorization"] != nil {
		reqToken := r.Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]

		tokenString := reqToken

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if !token.Valid {
			Log.STDLog.Warn("Invalid token")
			return data.User{}, false
		}
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				Log.STDLog.Error(err)
				return data.User{}, false
			}
			Log.STDLog.Error(err)
			return data.User{}, false
		}

		return data.User{
			Email: claims.Email,
		}, true
	} else {
		return data.User{}, false

	}

}
