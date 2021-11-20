package srv

import (
	//"calendar/common"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	//"log"
	"net/http"
	//"os"
	"strconv"
	"time"
)

const (
	host       string = "127.0.0.1"
	configFile string = "config/config.json"
)

type Config struct {
	Port        int32  `json:"Port"`
	LoggerLevel string `json:"LoggerLevel"`
}

func ReadConfig() (Config, error) {
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Printf("Config %s file not found. \n", configFile)

	}

	conf := Config{}

	err = json.Unmarshal(file, &conf)
	if err != nil {
		fmt.Printf("err %s", err.Error())
		errors.New("unable to unmarshal config file: " + configFile)
	}

	fmt.Println(conf)
	return conf, nil
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

	err := srv.ListenAndServe()
	if err != nil || err != http.ErrServerClosed {
		msg := "serve is dead!"

		panic(msg)
	}
	return nil
}
