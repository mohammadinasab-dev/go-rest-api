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
	UserID     int
	AuthUUID   string
	jwt.StandardClaims
}

//SetJwtToken generate and sets the Authorization token
func SetJwtToken(w http.ResponseWriter, auth data.Authentication) bool {

	expTime := time.Now().Add(10 * time.Minute)

	claims := &Claims{
		authorized: true,
		UserID:     auth.UserID,
		AuthUUID:   auth.AuthUUID,
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

//IsValid checks the Authorization token
func IsValid(r *http.Request) (data.Authentication, bool) {

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
			return data.Authentication{}, false
		}
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				Log.STDLog.Error(err)
				return data.Authentication{}, false
			}
			Log.STDLog.Error(err)
			return data.Authentication{}, false
		}
		var auth data.Authentication
		auth.UserID = claims.UserID
		auth.AuthUUID = claims.AuthUUID

		return auth, true
	}

	return data.Authentication{}, false

}
