package srv

import (
	"calendar/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

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

func eventsHandler(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	e := model.Event{}
	err = json.Unmarshal(body, &e)
	if err != nil {
		fmt.Fprintf(w, "Unmarshal err %s \n", err.Error())
		return
	}

	fmt.Println(string(body))

	if !UserLoggedIn(e.Name) {
		fmt.Fprintf(w, "User was not logged in %s \n", e.Name)
		return
	}

	_, found := model.Events[e.Id]
	if found {
		model.Events[e.Id] = e
	}

	return
}

func getEventHandler(w http.ResponseWriter, req *http.Request) {
	var eventId model.Id
	model.ParseReq(req.Body, &eventId)

	eVents := make(map[string]model.Event)

	for k, v := range model.Events {
		if v.Id == string(eventId) {
			eVents[k] = v
		}
	}

	fmt.Fprintf(w, "Num of events {%d}  \n", len(eVents))
}

func postEventHandler(w http.ResponseWriter, req *http.Request) {

}

func putEventHandler(w http.ResponseWriter, request *http.Request) {
	c := context.WithValue(request.Context(), "id", "user_id")
	request.WithContext(c)

	//context.WithTimeout(c, (2 * time.Duration())
	//updateDB(c, data))
	v := request.Context().Value("w")
	if v != nil {
		s := v.(string)
		fmt.Println(s)
	}
}

func logoutHandler(w http.ResponseWriter, req *http.Request) {
	v := req.Context().Value("id")

	if v != nil {

		s := v.(string)
		fmt.Println(s)
	} else {

		log.Println()
	}

	u := model.User{}
	err := model.ParseReq(req.Body, &u)
	if err != nil {
		fmt.Fprintf(w, "Bad data \n")
	}

	for idx, v := range model.Users {
		fmt.Println("idx", idx, "v name: ", v.Name, "uname: ", u.Name)
		if v.Name == u.Name {
			RemoveElement(model.Users, idx)
			fmt.Fprintf(w, "User {%s} logged out \n", u.Name)
			return
		}
	}

	fmt.Fprintf(w, "User {%s} was not logged in. \n", u.Name)
	return
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	u := model.User{}
	model.ParseReq(req.Body, &u)

	u.Token = GenToken(&u)
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Fprintf(w, "err %s \n", err.Error())
		return
	}

	w.Header().Add("Content-Type", " application/json")
	fmt.Fprintf(w, string(b))

	//fmt.Fprintf(w, "User {%s} need to create an account firt. \n", u.Name)
	return
}

func userHandler(w http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	u := model.User{}
	json.Unmarshal(body, &u)

	fmt.Println(string(body))

	if !UserLoggedIn(u.Name) {
		fmt.Fprintf(w, "User {%s} not logged in. \n", u.Name)
	}

	//todo update events
	return
}

func registerHandlers() *mux.Router {
	r := mux.NewRouter()

	//user PUT
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", auth(logoutHandler))

	//user PUT
	r.HandleFunc("/api/user", auth(userHandler))

	//user PUT
	r.HandleFunc("/api/events", auth(eventsHandler))

	//events GET PUT POST
	r.HandleFunc("/api/event", auth(getEventHandler)).Methods(http.MethodGet)
	r.HandleFunc("/api/event/{id}", auth(postEventHandler)).Methods(http.MethodPost)
	r.HandleFunc("/api/event/{id}", auth(putEventHandler)).Methods(http.MethodPut)

	return r
}

func RemoveElement(u []model.User, idx int) []model.User {
	return append(u[:idx], u[idx+1:]...)
}
