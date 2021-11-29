package main

import (
	"calendar/db"
	"flag"
	"fmt"
	"strconv"

	"calendar/logger"
	"calendar/srv"
)

var Version string

func main() {
	fmt.Print(Version)

	port := flag.Int("port", -1, "server port")
	m := flag.Bool("migrate", false, "run migrations")
	flag.Parse()

	if *m == true {
		err := db.Run()

		if err != nil {
			fmt.Println("run migrations fails ", err.Error())
			logger.L.Warning("run migrations fails ", err.Error())
		}

		logger.L.Info("Migration run was successful")
	}

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
