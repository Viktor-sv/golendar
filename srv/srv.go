package srv

import (
	"calendar/common"
	"calendar/merror"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

const (
	host       string = "127.0.0.1"
	configFile string = "config/config.json"
	Identity   string = "srv"
)

type Config struct {
	Port        int32  `json:"Port"`
	LoggerLevel string `json:"LoggerLevel"`
}

func ReadConfig() (Config, error) {
	date, err := common.LocalTime("Europe/Kiev")
	if err != nil {
		fmt.Printf("failed to get time")
	}

	conf := Config{}
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		msg := "Config" + configFile + "file not found."
		fmt.Printf(msg)
		return conf, merror.E(date, Identity, msg)
	}
	/*
		runtime.ReadMemStats(&mem)
		//runtime.ReadMemStats(&m1)

		merror.E(date, Identity, "messgase egtrtryrtyrtdfgdgfgre5465")
		//t3 := map[int]string{1: "x"}

		runtime.ReadMemStats(&m1)

		fmt.Println(m1.HeapAlloc - mem.HeapAlloc)
	*/
	err = json.Unmarshal(file, &conf)
	if err != nil {
		msg := "unable to unmarshal config file: " + configFile
		fmt.Printf("err %s", err.Error())
		errors.New(msg)
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

func memUsage(m1, m2 *runtime.MemStats) {
	p("Alloc:", m2.Alloc-m1.Alloc,
		"TotalAlloc:", m2.TotalAlloc-m1.TotalAlloc,
		"HeapAlloc:", m2.HeapAlloc-m1.HeapAlloc)
}

var p = fmt.Println

var m1, m2, mem runtime.MemStats
