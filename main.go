package main

import (
	"calendar/srv"
	"flag"
	"fmt"
	"unsafe"
)

var Version string

func main() {
	port := flag.Int("port", -1, "server port")
	flag.Parse()

	fmt.Print(Version)

	conf := srv.ReadConfig()
	if *port != -1 {
		conf.Port = int32(*port)
	}

	//todo error create own error

	if err := srv.Start(conf.Port); err != nil {
		//panic(err.Error())
	}

}
