package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"

	"net/http"
	"time"
)


var user = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credential struct {
	Username string `json:username`
	Password string `json:password`
}
type Claims struct {
	Username string `json:username`
	jwt.StandardClaims
}
var jwtKey = []byte("secret_key")
func Login(w http.ResponseWriter, r *http.Request) {
	var credential Credential
	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expectedPass, ok := user[credential.Username]
	if !ok || expectedPass != credential.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	expireAt := time.Now().Add(time.Minute * 5)
	claim := &Claims{
		Username: credential.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireAt.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w,
		&http.Cookie{
			Name:    "token",
			Value:   tokenStr,
			Expires: expireAt,
		})
}
func Home(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenStr := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte(fmt.Sprintf("hello, %s",claims.Username)))
}
func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie,err:=r.Cookie("token")
	if err!=nil{
		if err==http.ErrNoCookie{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenStr:=cookie.Value
	claims:=&Claims{}
	tkn,err:=jwt.ParseWithClaims(tokenStr,claims,
		func (t *jwt.Token)(interface{},error){
			return jwtKey,nil
		})
	if err!=nil{
		if err==jwt.ErrSignatureInvalid{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	expireAt:=time.Now().Add(time.Minute*5)
	claims.ExpiresAt=expireAt.Unix()
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenStr,err=token.SignedString(jwtKey)
	if err!=nil{
		if err==jwt.ErrSignatureInvalid{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	http.SetCookie(w,
		&http.Cookie{
			Name: "token",
			Value: tokenStr,
			Expires: expireAt,
		})

}