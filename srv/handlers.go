package srv

import (
	"calendar/common"
	"calendar/db"
	"calendar/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func eventsHandler(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {

	}

	e := model.Event{}
	err = json.Unmarshal(body, &e)
	if err != nil {
		fmt.Fprintf(w, "Unmarshal err %s \n", err.Error())
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
	event := model.Event{}
	model.ParseReq(req.Body, &event)

	e := db.GetEvent(event.Id)
	if nil == e {
		db.AddE(event)
		w.Header().Add("Content-Type", " application/json")
		fmt.Fprintf(w, "db.Add(user) \n")
		return
	}

}

func putEventHandler(w http.ResponseWriter, request *http.Request) {
	c := context.WithValue(request.Context(), "id", "user_id")
	request.WithContext(c)

	v := request.Context().Value("w")
	if v != nil {
		s := v.(string)
		fmt.Println(s)
	}
}

func logoutHandler(w http.ResponseWriter, req *http.Request) {

	tokenHeader := req.Header.Get("Authorization") //get token
	token := strings.Split(tokenHeader, " ")[1]

	v := req.Context().Value("id")

	if v != nil {
		s := v.(string)
		fmt.Println(s)
	} else {

		log.Println()
	}

	user := model.User{}
	user.Token = token

	err := model.ParseReq(req.Body, &user)
	if err != nil {
		fmt.Fprintf(w, "Bad data \n")
	}

	err = db.LogOut(&user)
	if err != nil {
		fmt.Println("log out error", err.Error())
	}

	fmt.Fprintf(w, "User {%s} was not logged in. \n", user.Name)
	return
}

func registrationHandlers(w http.ResponseWriter, req *http.Request) {
	user := model.User{}
	model.ParseReq(req.Body, &user)

	u := db.GetUser(user.Name, user.Pass)
	if u == nil {
		db.Add(user)
		w.Header().Add("Content-Type", " application/json")
		fmt.Fprintf(w, "db.Add(user) \n")
		return
	}

	if u.Name == user.Name && user.Pass == u.Pass {
		fmt.Println("user already logged in")
	}

	w.Header().Add("Content-Type", " application/json")

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

	bu := db.GetUser(u.Name, u.Pass)
	if bu == nil {
		msg := "user not registered"
		fmt.Println(msg)
		fmt.Fprintf(w, msg)
		return
	}

	if u.Name == bu.Name && bu.Pass == u.Pass {
		fmt.Println("user already logged in")
	}

	err = db.LoginUser(u)
	if err != nil {
		fmt.Println("error", err.Error())
	}

	msg := "db.LoginUser" + u.Name
	fmt.Println(msg)
	w.Header().Add("Content-Type", " application/json")

	fmt.Fprintf(w, string(b))
	fmt.Fprintf(w, msg)
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

	if !common.UserLoggedIn(u.Name) {
		fmt.Fprintf(w, "User {%s} not logged in. \n", u.Name)
	}

	return
}

func registerHandlers() *mux.Router {
	r := mux.NewRouter()

	//user POST
	r.HandleFunc("/registration", registrationHandlers)
	r.HandleFunc("/login", loginHandler)

	//user GET
	r.HandleFunc("/logout", auth(logoutHandler))

	//user PUT
	r.HandleFunc("/api/user", auth(userHandler))

	//user GET
	r.HandleFunc("/api/events", auth(eventsHandler))

	//events GET PUT POST
	r.HandleFunc("/api/event/{id}", auth(getEventHandler)).Methods(http.MethodGet)
	r.HandleFunc("/api/event/{id}", auth(putEventHandler)).Methods(http.MethodPut)
	r.HandleFunc("/api/event", auth(postEventHandler)).Methods(http.MethodPost)

	return r
}

func RemoveElement(u []model.User, idx int) []model.User {
	return append(u[:idx], u[idx+1:]...)
}
