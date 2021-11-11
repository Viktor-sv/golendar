package factory_

import (
	"fmt"
	"time"
)

type myIntreface interface {
	Run(time time.Duration) error
}

type worker struct {
	name string
}

func (s *worker) Error() string {
	return ""
}
func (s *worker) String() string {
	return ""
}

func (s *worker) Run(duration time.Duration) error {
	fmt.Printf("worker run \n")
	time.Sleep(duration)
	fmt.Println(s.name)
	return nil
}

func (s *worker) myfucn(i int) error {
	return s
}

type DataWriter interface {
	WriteTo(string2 string)
}

type Map struct {
	size    int
	storage map[string]string
}

type DB struct {
}

func (m *Map) WriteTo(string2 string) {

	m.storage["k"] = string2

}

func (m *DB) WriteTo(string2 string) {
	//connect to DB
	//write to db

}

func do(w myIntreface) {

}

type dd int

func (i dd) Run(duration time.Duration) error {
	fmt.Printf("dd run \n")
	time.Sleep(time.Duration(i) * time.Second)
	return nil
}

func Factory(t string) DataWriter {

	switch t {
	case "DB":
		return &DB{}
	case "MAP":
		return &Map{storage: make(map[string]string)}
	}

	return nil
}

//example of using
//f := factory_.Factory("DB")
//f.WriteTo("hollo")
