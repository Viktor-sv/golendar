package common

import (
	"calendar/db"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	myTokenPass string = "paswdwdwdwdwdwdwd"
)

const (
	fileTimeStampFormat = "2006-01-02-15_04_05"
)

// LocalTime this func will return time in a format,
// YYYY_MM_DD_HH_MM_SS
func LocalTime(location string) (string, error) {
	loc, err := time.LoadLocation(location)
	if err != nil {
		return "", err
	}

	t := time.Now().In(loc)
	fmt.Println(t.Format(fileTimeStampFormat))

	//fmt.Println(t.Format(fileTimeStampFormat))
	return t.Format(fileTimeStampFormat), nil
}

func UserLoggedIn(token string) bool {
	fmt.Printf("UserLoggedIn: entering. %s \n", token)
	t, err := ParseToken(token)
	fmt.Printf("parsed tokeng. %s \n", t)
	if err != nil {
		fmt.Printf("error parsing tokeng. %s \n", err)
		return false
	}

	users := db.GetUsers()
	for _, v := range users {

		/*		v, err := ParseToken(v.Token)
				if err != nil {
					fmt.Printf("err\n", err.Error())
					return false
				}*/

		if token == v.Token {
			fmt.Printf("UserLoggedIn: tokens are the same. \n")
			return true
		}
	}

	fmt.Printf("UserLoggedIn: tokens are not the same. \n")
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
		return "", err
	}

	if !token.Valid {
		fmt.Println("token not valid: ")
	}

	return token.Raw, nil
}
