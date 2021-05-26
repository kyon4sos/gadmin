package util

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func CreatToken(username string)  (string,error){
	mySigningKey := []byte("AllYourBase")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user": username,
		})
	signedString, err := token.SignedString(mySigningKey)
	fmt.Printf("%s",err)
	fmt.Printf("token: %s\n",signedString)
	return signedString,err
}
