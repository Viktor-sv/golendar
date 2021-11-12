package srv

import (
	"calendar/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	UserId uint
	jwt.StandardClaims
	Token string `json:"token"`
}

const (
	myTokenPass string = "paswdwdwdwdwdwdwd"
)

func GenToken(user *model.User) string {
	t := &Token{UserId: 15}
	//t.ExpiresAt = time.Now().Add(72 * time.Hour).Unix()
	t.Subject = user.Name
	t.Issuer = "golendar"

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), t.StandardClaims)

	tokenS, _ := token.SignedString([]byte(myTokenPass))

	user.Token = tokenS

	return tokenS
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
