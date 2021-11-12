package main

import (
	"calendar/srv"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

func eventsHandler(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	e := Event{}
	err = json.Unmarshal(body, &e)
	if err != nil {
		fmt.Fprintf(w, "Unmarshal err %s \n", err.Error())
		return
	}

	fmt.Println(string(body))

	if !userLoggedIn(e.Name) {
		fmt.Fprintf(w, "User was not logged in %s \n", e.Name)
		return
	}

	_, found := events[e.Id]
	if found {
		events[e.Id] = e
	}

	return
}

func getEventHandler(w http.ResponseWriter, req *http.Request) {
	var eventId id
	parseReq(req.Body, &eventId)

	eVents := make(map[string]Event)

	for k, v := range events {
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

	u := User{}
	err := parseReq(req.Body, &u)
	if err != nil {
		fmt.Fprintf(w, "Bad data \n")
	}

	for idx, v := range users {
		fmt.Println("idx", idx, "v name: ", v.Name, "uname: ", u.Name)
		if v.Name == u.Name {
			removeElement(users, idx)
			fmt.Fprintf(w, "User {%s} logged out \n", u.Name)
			return
		}
	}

	fmt.Fprintf(w, "User {%s} was not logged in. \n", u.Name)
	return
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	u := User{}
	parseReq(req.Body, &u)
	fmt.Println(len(users))

	u.Token = genToken(&u)
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

	u := User{}
	json.Unmarshal(body, &u)

	fmt.Println(string(body))

	if !userLoggedIn(u.Name) {
		fmt.Fprintf(w, "User {%s} not logged in. \n", u.Name)
	}

	//todo update events
	return
}


func registerHandlers() *mux.Router {
	r := mux.NewRouter()

	//user PUT
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", srv.auth(logoutHandler))

	//user PUT
	r.HandleFunc("/api/user", srv.auth(userHandler))

	//user PUT
	r.HandleFunc("/api/events", srv.auth(eventsHandler))

	//events GET PUT POST
	r.HandleFunc("/api/event", srv.auth(getEventHandler)).Methods(http.MethodGet)
	r.HandleFunc("/api/event/{id}", srv.auth(postEventHandler)).Methods(http.MethodPost)
	r.HandleFunc("/api/event/{id}", srv.auth(putEventHandler)).Methods(http.MethodPut)

	return r
}