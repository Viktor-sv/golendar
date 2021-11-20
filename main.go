package main

import (
	"calendar/logger"
	"calendar/srv"
	"flag"
	"fmt"
	"strconv"
)

var Version string

func main() {
	port := flag.Int("port", -1, "server port")
	flag.Parse()
	fmt.Print(Version)

	conf, err := srv.ReadConfig()
	if err != nil {
		fmt.Println("Error", err.Error()) //todo here should be my error
	}

	logger.L.Info("main test")

	if *port != -1 {
		conf.Port = int32(*port)
		logger.L.Warning("port not set", strconv.Itoa(*port))
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	srv.Start(conf.Port)

}
