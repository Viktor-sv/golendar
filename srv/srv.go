package srv

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	//"log
	"net/http"
	"strconv"
	"time"
)

const (
	host string = "127.0.0.1"
)

type Config struct {
	Port        int32  `json:"Port"`
	LoggerLevel string `json:"LoggerLevel"`
}

const (
	configFile string = "config/config.json"
)

func intiLog() {
	//fmt.Println("Hello world!", common.LocalTime("Chernivtsi"))
}

func ReadConfig() Config {
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
func Start(port int32) error {

	r := registerHandlers()
	srv := &http.Server{
		Handler: r,
		Addr:    host + ":" + strconv.Itoa(int(port)),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//log.Fatal(srv.ListenAndServe())
	err := srv.ListenAndServe()
	if err != nil || err != http.ErrServerClosed {

		return err
	}
	return nil
}
