package main

import (
	"calendar/db"
	"flag"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"

	"calendar/logger"
	"calendar/srv"
)

var Version string

func main() {
	var wg sync.WaitGroup
	var data int32

	//var mx sync.Mutex
	//	var rwmx sync.RWMutex

	for i := 1; i <= 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			/*rwmx.Lock()// 0x45454 l
			defer rwmx.Unlock()
			data = data + 1*/
			atomic.AddInt32(&data, 1)

		}()
	}

	wg.Wait()

	if data == 0 {
		fmt.Printf("----------------------------the value is %v.\n", data)
	} else {
		fmt.Printf("----------------------------the value is %v.\n", data)

	}

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
