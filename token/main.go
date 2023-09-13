package token

import (

	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)
var secretKey = []byte("https://www.postman.com/descent-module-technologist-70995397/workspace/jwt") // Replace with your secret key

func createToken() (string, error) {
    claims := jwt.MapClaims{
        "username": "exampleuser",
        "exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(secretKey)
}

func handler(w http.ResponseWriter, r *http.Request) string{
    tokenString, err := createToken()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return "error"
    }
    return tokenString
}