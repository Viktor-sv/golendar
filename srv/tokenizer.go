package srv

import (
	"calendar/model"
	//"fmt"
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
