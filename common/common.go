package common

import (
	"calendar/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	myTokenPass string = "paswdwdwdwdwdwdwd"
)

func LocalTime(location string) error {
	var Tmp int32 = 0
	fmt.Println(Tmp)
	loc, err := time.LoadLocation(location)
	if err != nil {
		//panic(err)
		return err
	}

	t := time.Now().In(loc)
	fmt.Println("%t", t)
	fmt.Println(t)
	return nil
}

func UserLoggedIn(token string) bool {
	t, err := ParseToken(token)
	if err != nil {
		fmt.Printf("error parsing tokeng. %s \n", err)
		return false
	}

	for _, v := range model.Users {
		v, _ := ParseToken(v.Token)
		if t == v {
			fmt.Printf("tokens are the same. \n")
			return true
		}
	}

	fmt.Printf("tokens are not the same. \n")
	return false
}

func ParseToken(tokenS string) (string, error) {

	fmt.Printf("token not valid %s \n", tokenS)
	token, err := jwt.Parse(tokenS, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(myTokenPass), nil
	})

	if err != nil {
		return "error: ", err
	}

	if !token.Valid {
		fmt.Println("token not valid: ")
	}

	return "", nil
}
