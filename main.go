package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	_ "os"
	"reflect"
	"strconv"
	"time"
	"unsafe"
)

type Config struct {
	Port int `json:"Port"`
	LoggerLevel string `json:"LoggerLevel"`
}

type User struct {
	Name string `json:"Username"`
	Pass string `json:"Password"`
}

type Event struct {
	Id string
	Title string
	Description string
	Name string
	Time string
	Timezone string
	Duration uint32
	Notes []string
}

const configFile string = "config/config.json"
//var data = make(map[string]*Event, 1000)
var events = make(map[string]Event, 1000)
var users = make([]User,0,10000)

type jsonPersarer interface{
	Parse(b []byte) error
}

func (e *Event) Parse(b []byte) error{

	err := json.Unmarshal(b, e)
	return err
}

func (e *User) Parse(b []byte) error{

	err := json.Unmarshal(b, e)
	return err
}

func parseReq(req io.Reader, v jsonPersarer) ( error) {
	body, err := ioutil.ReadAll(req)
	if err != nil {
		return err
	}

	err = v.Parse(body)
	return err
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

	//req.Body.Read()
	u := User{}
	err = json.Unmarshal(body, &u)
	if err != nil {
		fmt.Fprintf(w, "Unmarshal err %s \n", err.Error())
		return
	}

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
	/*body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalln(err)
	}

	u := User{}
	json.Unmarshal(body, &u)

	*/
	u := User{}
	e := Event{Name: "Andy"}
	e1 := Event{}
	err := parseReq(req.Body, &u)

	if err != nil {
		fmt.Fprintf(w, "Bad data \n")
	}

	for idx, v := range users {
		fmt.Println("idx", idx, "v name: ",v.Name, "uname: ", u.Name)
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
	if  found {
		events[e.Id] = e
	}

	return
}

//eventHandler process req  from ...
func eventHandler(w http.ResponseWriter, req *http.Request) {

}

func getSliceHeader(slice *[]int) string {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(slice))
	return fmt.Sprintf("%+v", sh)
}

func registerHandlers() *mux.Router {
	r := mux.NewRouter()

	//author POST GET
	r.HandleFunc("/hello", hello)

	//user PUT
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", logoutHandler)

	//user PUT
	r.HandleFunc("/api/put", eventsHandler)

	//events GET PUT POST
	r.HandleFunc("/api/events", eventsHandler).Methods(http.MethodGet, http.MethodPost, http.MethodPut)
	return r
}

func intiLog()  {
	fmt.Println("Hello world!", localTime("Chernivtsi"))
}

func userHandler(w http.ResponseWriter, req *http.Request){

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

func userLoggedIn(user string) bool{

	for _, v := range users {
		//	fmt.Println("asd",k, "v name: ",v.Name, "uname: ", u.Name)
		if v.Name == user {
			return true
		}
	}

	return false
}

func readConfig() Config {
	file, err := ioutil.ReadFile(configFile)
	if err != nil{
		fmt.Printf( "Config %s file not found. \n", configFile)
	}

	conf := Config{}

	err = json.Unmarshal(file, &conf)
	if err != nil{
		 fmt.Printf("err %s", err.Error())
	}

	fmt.Println(conf)
	return conf
}

const host string = "127.0.0.1"
func main() {
    //f := factory_.Factory("DB")
	//f.WriteTo("hollo")

	intiLog()
    conf := readConfig()

	r := registerHandlers()
	srv := &http.Server{
		Handler: r,
		Addr:    host + ":"+ strconv.Itoa(conf.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
