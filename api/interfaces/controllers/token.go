package controllers

import (
	"fmt"
	"strings"
	"github.com/dgrijalva/jwt-go"
)

func checkToken(id int, tokenString string) bool {
	if tokenString == "" {
		return false
	}

	if len(tokenString) < 7 || strings.ToLower(tokenString[0:6]) != "bearer" {
		return false
	}
	  
	tokenString = tokenString[7:]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Invalid token")
		}

		return []byte("secret"), nil
	})

	if err != nil {
		return false
	}
  
	if !token.Valid {
	  	return false
	}
  
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
	  	return false
	}
  
	checkId := int(claims["id"].(float64))
	if checkId != id {
		return false
	}

	return true
}