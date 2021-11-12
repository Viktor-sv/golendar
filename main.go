package main

import (
	"calendar/srv"
)

func main() {
	var arv int32 = 2
	conf := srv.ReadConfig()
	if arv <= 0 {
		conf.Port = arv
	}

	srv.Start(conf.Port)
}
