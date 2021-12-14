package model

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

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

var myslice []int32

func GetNews() error {
	for 10000000000 {
		myslice = append(myslice, 1)
	}
	return nil
}

func GetSlice() *[]int32 {
	return &myslice
}

var Events = make(map[string]Event, 1000)
var Users = make([]User, 0, 10000)

type jsonPersarer interface {
	Parse(b []byte) error
}

type Id string

func (e *Event) Parse(b []byte) error {
	err := json.Unmarshal(b, e)
	return err
}

func (u *User) Parse(b []byte) error {
	err := json.Unmarshal(b, u)
	return err
}

func (i *Id) Parse(b []byte) error {
	err := json.Unmarshal(b, i)
	return err
}

func ParseReq(req io.Reader, v jsonPersarer) error {
	body, err := ioutil.ReadAll(req)
	if err != nil {
		return err
	}

	err = v.Parse(body)
	return err
}
