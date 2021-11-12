package srv

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

//auth
func auth(h http.HandlerFunc) http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {

		tokenHeader := request.Header.Get("Authorization") //get token
		token := strings.Split(tokenHeader, " ")[1]
		//logging
		//to do recovery

		c := context.WithValue(request.Context(), "id", "user_id")
		request = request.WithContext(c)

		//fmt.Printf("before http handler %s", token)
		if UserLoggedIn(token) {
			fmt.Println(writer, "user not authorized")
			return
		}
		//to do add  id user to context
		h(writer, request)
	}
}
