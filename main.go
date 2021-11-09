package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	_ "os"
	"reflect"
	"strconv"
	"strings"

	"time"
	"unsafe"
)

type Config struct {
	Port        int    `json:"Port"`
	LoggerLevel string `json:"LoggerLevel"`
}

type User struct {
	Name  string `json:"Username"`
	Pass  string `json:"Password"`
	Token string `json:"token"`
}

type Event struct {
	Id          string
	Title       string
	Description string
	Name        string
	Time        string
	Timezone    string
	Duration    uint32
	Notes       []string
}

type Token struct {
	UserId uint
	jwt.StandardClaims
	Token string `json:"token"`
}

const (
	configFile string = "config/config.json"
	host       string = "127.0.0.1"
)

var events = make(map[string]Event, 1000)
var users = make([]User, 0, 10000)

type jsonPersarer interface {
	Parse(b []byte) error
}

func (e *Event) Parse(b []byte) error {

	err := json.Unmarshal(b, e)
	return err
}

func (e *User) Parse(b []byte) error {

	err := json.Unmarshal(b, e)
	return err
}

type id string

func (i *id) Parse(b []byte) error {

	err := json.Unmarshal(b, i)
	return err
}
func parseReq(req io.Reader, v jsonPersarer) error {
	body, err := ioutil.ReadAll(req)
	if err != nil {
		return err
	}

	err = v.Parse(body)
	return err
}

func response(w http.ResponseWriter, req *http.Request) string {
	return "OK"
}

// a handler is by using the `http.HandlerFunc` adapter
// on functions with the appropriate signature.
func hello(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Host)
	fmt.Fprintf(w, "hello my tyest \n")
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

func removeElement(u []User, idx int) []User {
	return append(u[:idx], u[idx+1:]...)
}

func logoutHandler(w http.ResponseWriter, req *http.Request) {
	u := User{}
	err := parseReq(req.Body, &u)
	if err != nil {
		fmt.Fprintf(w, "Bad data \n")
	}

	for idx, v := range users {
		fmt.Println("idx", idx, "v name: ", v.Name, "uname: ", u.Name)
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
	if found {
		events[e.Id] = e
	}

	return
}

//eventHandler process req  from ...
func getEventsHandler(w http.ResponseWriter, req *http.Request) {
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

func postEventsHandler(w http.ResponseWriter, req *http.Request) {

}

func putEventsHandler(w http.ResponseWriter, request *http.Request) {

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

func getSliceHeader(slice *[]int) string {
	sh := (*reflect.SliceHeader)(unsafe.Pointer(slice))
	return fmt.Sprintf("%+v", sh)
}

func registerHandlers() *mux.Router {
	r := mux.NewRouter()

	//user PUT
	r.HandleFunc("/login", loginHandler)
	r.HandleFunc("/logout", auth(logoutHandler))

	//user PUT
	r.HandleFunc("/api/put", auth(eventsHandler))

	//events GET PUT POST
	r.HandleFunc("/api/events", auth(getEventsHandler)).Methods(http.MethodGet)
	r.HandleFunc("/api/events", auth(postEventsHandler)).Methods(http.MethodPost)
	r.HandleFunc("/api/events", auth(putEventsHandler)).Methods(http.MethodPut)

	return r
}

func intiLog() {
	fmt.Println("Hello world!", localTime("Chernivtsi"))
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

func parseToken(tokenS string) (string, error) {

	fmt.Printf("token not valid %s \n", tokenS)
	token, err := jwt.Parse(tokenS, func(token *jwt.Token) (interface{}, error) {

		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return string([]byte("my_token_pass")), nil
	})

	if err != nil {
		return "errr", err
	}

	if !token.Valid {
		fmt.Println("token not valid")
	}

	return "", nil
}

/*func parseToken(t string) error {
	token, err := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret"), nil
	})

	if err != nil {
		fmt.Printf("parce token err %s", err.Error())
		return err
	}

	if !token.Valid {
		fmt.Println("token not valid")
	}
	return err
}
*/

func userLoggedIn(token string) bool {
	t, err := parseToken(token)

	if err != nil {
		fmt.Printf("error parsing tokeng. %s \n", err)
		return false
	}

	for _, v := range users {
		v, _ := parseToken(v.Token)
		if t == v {
			fmt.Printf("tokens are the same. \n")
			return true
		}
	}

	fmt.Printf("tokens are not the same. \n")
	return false
}

func readConfig() Config {
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Config %s file not found. \n", configFile)
	}

	conf := Config{}

	err = json.Unmarshal(file, &conf)
	if err != nil {
		fmt.Printf("err %s", err.Error())
	}

	fmt.Println(conf)
	return conf
}

func genToken(user *User) string {
	t := &Token{UserId: 15}
	t.ExpiresAt = time.Now().Add(2 * time.Second).Unix()
	t.Subject = user.Name
	t.Issuer = "golendar"
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), t)
	tokenS, _ := token.SignedString([]byte("my_token_pass"))
	user.Token = tokenS

	return tokenS
}

func emptyIf(i interface{}) interface{} {

	/*
		v, ok := i.(int) //cast
		if !ok {
			fmt.Println("empty is not int ")
			return false
		}

		u, ok := i.(User) //cast
		if !ok {
			fmt.Println("empty is not user ")
			return false
		}

		fmt.Println("empty %d", v)
		fmt.Println("empty %d", u)
		return true*/

	/*	switch val := i.(type) {
		case int:
			fmt.Printf("empty is not int %d \n", val)
		case string:
			fmt.Printf("empty is not int %s \n", val)
		case bool:
			fmt.Printf("empty is not int %s \n", val)
		default:
			fmt.Printf("type is not defined %T \n", val)

		}*/
	return nil
}

func main() {
	//f := factory_.Factory("DB")
	//f.WriteTo("hollo")
	user := User{Name: "Andy", Pass: "123123", Token: ""}
	users = append(users, user)

	t := genToken(&user)
	fmt.Println("token: ", t)
	intiLog()

	conf := readConfig()
	srv := &http.Server{
		Handler: registerHandlers(),
		Addr:    host + ":" + strconv.Itoa(conf.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}

//auth
func auth(h http.HandlerFunc) http.HandlerFunc {

	return func(writer http.ResponseWriter, request *http.Request) {
		tokenHeader := request.Header.Get("Authorization") //get token
		token := strings.Split(tokenHeader, " ")[1]
		//logging
		//to do recovery

		//fmt.Printf("before http handler %s", token)
		if userLoggedIn(token) {
			fmt.Println(writer, "user not authorized")
			return
		}

		//to do add  id user to context
		h(writer, request)
	}

}

//cource
