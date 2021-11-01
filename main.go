package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	_ "os"
	"reflect"
	"unsafe"

	//"io"
	"net/http"
	"time"
)

type Event struct{
	id string
	title string
	description string
	name string
	time string
	timezone string
	duration uint32
	notes []string
}

var data = make(map[string]*Event)

func parseReq([]byte) (*Event, error) {
	ev := Event{}

	return &ev, nil
}

func proccessReq(e *Event){

}

func response() string{
	return "OK"
}

// A fundamental concept in `net/http` servers is
// *handlers*. A handler is an object implementing the
// `http.Handler` interface. A common way to write
// a handler is by using the `http.HandlerFunc` adapter
// on functions with the appropriate signature.
func hello(w http.ResponseWriter, req *http.Request) {

	// Functions serving as handlers take a
	// `http.ResponseWriter` and a `http.Request` as
	// arguments. The response writer is used to fill in the
	// HTTP response. Here our simple response is just
	// "hello\n".
	log.Println(req.Host)
	fmt.Fprintf(w, "hello my tyest \n")
}

func loginHandler(w http.ResponseWriter, req *http.Request) {
	log.Println(req.Host)
	fmt.Fprintf(w, "hello my login \n")
}

func logoutHandler(w http.ResponseWriter, req *http.Request) {

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

	//events POST GET PUT
	//r.HandleFunc("/api/events", eventHandler)
	//	r.HandleFunc("/login", loginHandler)
	//	r.HandleFunc("/logout", logoutHandler)
	//r.HandleFunc("/products/{key}", eventHandler).Methods(http.MethodGet)
	//r.HandleFunc("/articles/{category}/", logoutHandler)
	//r.HandleFunc("/articles/{id}/{id:[0-9]+}", eventHandler)
}

func main() {
	fmt.Println("Hello world!", localTime("Chernivtsi"))
	registerHandlers()

	http.ListenAndServe(":8090", nil)
}
