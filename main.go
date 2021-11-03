package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	_ "os"
	"reflect"
	"time"
	"unsafe"
)

type User struct{
	Name string `json:"Username"`
	Pass string `json:"Password"`
}

type Event struct{
	Id string
	Title string
	Description string
	Name string
	Time string
	Timezone string
	Duration uint32
	Notes []string
}

var data = make(map[string]*Event, 1000)
var users = make([]User,0,10000)

func parseReq(b []byte) (*Event, error) {
	//We Read the response body on the line below.

	var ev *Event = &Event{}
	json.Unmarshal(b, ev)
	return ev, nil
}

func proccessReq(e *Event){

}

func response() string{
	return "OK"
}

// a handler is by using the `http.HandlerFunc` adapter
// on functions with the appropriate signature.
func hello(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Host)
	fmt.Fprintf(w, "hello my tyest \n")
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	u := User{}
	json.Unmarshal(body, &u)

	fmt.Println(string(body))
	fmt.Println(len(users))

	for k, v := range users {
		fmt.Println("asd",k, "v name: ",v.Name, "uname: ", u.Name)

		if v.Name == u.Name {
			fmt.Fprintf(w, "User {%s} already logged in1 \n", u.Name)
			return
		}

		tmp := User{u.Name, u.Pass}
		users = append(users, tmp)
		fmt.Fprintf(w, "User {%s} logged in. \n", u.Name)
		return
	}

	tmp := User{u.Name, u.Pass}
	users = append(users, tmp)
	fmt.Fprintf(w, "User {%s} logged in. \n", u.Name)
	return
}

func removeElement(u []User, idx int) []User {
	return append(u[:idx], u[idx+1:]...)
}

func logoutHandler(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	u := User{}
	json.Unmarshal(body, &u)

	//fmt.Println(string(body))
	//fmt.Println(len(users))

	for idx, v := range users {
		if v.Name == u.Name {
			//delete  from map
			removeElement(users, idx)
			fmt.Fprintf(w, "User {%s} logged out \n", u.Name)
			return
		}

		fmt.Fprintf(w, "User {%s} was not logged in. \n", u.Name)
		return
	}

	fmt.Fprintf(w, "User {%s} was not logged in. \n", u.Name)
	return
}

// localTime return time based on location
func localTime(location string) int {
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		panic(err)
	}

	t := time.Now().In(loc)
	fmt.Println("%t", t)
	fmt.Println(t)
	return 0
}

func eventsHandler(w http.ResponseWriter, req *http.Request) {
	b := make([]byte, 10)
	data, err := parseReq(b)
	if err != nil {

	}

	proccessReq(data)
	fmt.Fprintf(w, response())
}

//eventHandler process req  from ...
func eventHandler(w http.ResponseWriter, req *http.Request) {

}

func getSliceHeader(slice *[]int) string {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(slice))
	return fmt.Sprintf("%+v", sh)
}

func registerHandlers() {
	r := mux.NewRouter()

	//author POST GET
	r.HandleFunc("/hello", hello)

	//user PUT
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", logoutHandler)

	//user PUT
	r.HandleFunc("/api/put", eventsHandler)

	//events GET
	r.HandleFunc("/events", eventsHandler).Methods(http.MethodGet)
}

func intiLog() {

}


func main() {
	fmt.Println("Hello world!", localTime("Chernivtsi"))
	//to do init log
	intiLog()

	//registerHandlers()

	r := mux.NewRouter()

	//author POST GET
	r.HandleFunc("/hello", hello)

	//user PUT
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", logoutHandler)

	//user PUT
	r.HandleFunc("/api/put", eventsHandler)

	//events GET
	r.HandleFunc("/events", eventsHandler).Methods(http.MethodGet)

	//http.ListenAndServe(":8090", nil)
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8090",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
