package main

import (
	"calendar/common"
	"calendar/srv"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	_ "os"
	"strconv"
	"time"
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
	configFile    string = "config/config.json"
	host          string = "127.0.0.1"
	my_token_pass string = "paswdwdwdwdwdwdwd"
)

var events = make(map[string]Event, 1000)
var users = make([]User, 0, 10000)

func removeElement(u []User, idx int) []User {
	return append(u[:idx], u[idx+1:]...)
}

func intiLog() {
	fmt.Println("Hello world!", common.LocalTime("Chernivtsi"))
}


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
		fmt.Printf("Config %s file not found. \n",  configFile)
	}

	conf := Config{}

	err = json.Unmarshal(file, &conf)
	if err != nil {
		fmt.Printf("err %s", err.Error())
	}

	fmt.Println(conf)
	return conf
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

	intiLog()
	conf := readConfig()
	r := registerHandlers()

	/*
	if(arv != nil)
		conf.port = arv

	*/

	srv.Start()
	//./golendar.exe --port=2323

	srv := &http.Server{
		Handler: r,
		Addr:    host + ":" + strconv.Itoa(conf.Port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}



func parseReq(req io.Reader, v jsonPersarer) error {
	body, err := ioutil.ReadAll(req)
	if err != nil {
		return err
	}

	err = v.Parse(body)
	return err
}

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
