package main

import (
	"calendar/db"
	"calendar/model"
	"flag"
	"fmt"
	"strconv"

	"calendar/logger"
	"calendar/srv"
)

var Version string

func main() {
	/*	u1 := &model.User{Name: "user1", Pass: "dfdrf", Token: "defrefre"}
		u2 := &model.User{Name: "user2", Pass: "dfdrf2", Token: "defrefre2"}
		fmt.Println(u1)

		s1 := []*model.User{u1, u2}

		s2 := &[]model.User{*u1, *u2}

		fmt.Println(s1)
		fmt.Println(s2)
	*/
	/*s1 := make([]model.User,1)
	s2 := make([]model.User,1)
	*/

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
