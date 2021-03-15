package jwt

import (
	"fmt"
	"go-rest-api/data"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_Secret_Key")

//Claims is jwt claim
type Claims struct {
	Email string
	jwt.StandardClaims
}

//SetJwtToken is jwt setter
func SetJwtToken(w http.ResponseWriter, login data.User) bool {

	expTime := time.Now().Add(10 * time.Minute)

	claims := &Claims{
		Email: login.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	stringToken, err := token.SignedString(jwtKey)

	if err != nil {
		return false
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "JWTToken",
		Expires:  expTime,
		Value:    stringToken,
		HttpOnly: true,
	})
	fmt.Fprintln(w, "cookies set")
	return true
}

//IsLogedin is is log check
func IsLogedin(r *http.Request) (data.User, bool) {

	c, err := r.Cookie("JWTToken")
	if err != nil {
		log.Println("i am here")
		log.Println(err)
		return data.User{}, false
	}

	tokenString := c.Value

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if !token.Valid {
		fmt.Println("not valid token")
		return data.User{}, false
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return data.User{}, false
		}
		return data.User{}, false
	}

	return data.User{
		Email: claims.Email,
	}, true
}
