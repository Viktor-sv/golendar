package main

import (
	"calendar/adder"
	"calendar/api/proto"
	"calendar/db"
	"flag"
	"fmt"
	"log"
	"net"
	//"os"
	"strconv"
	"sync"
	//"sync/atomic"
	"google.golang.org/grpc"

	"calendar/logger"
	"calendar/srv"
)

var Version string

func readCH(c chan int) {

	for {
		r, ok := <-c
		fmt.Println("r:", r, " ok:", ok)

		if ok == false {
			return
		}
	}
}

func writeCH(i ...int) chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for _, val := range i {
			ch <- val

		}
	}()

	return ch
}

func test() {
	c := writeCH(1, 2, 34, 5, 6, 6)
	readCH(c)
}

func g(wg *sync.WaitGroup) {
	defer wg.Done()

	s := grpc.NewServer()
	sr := &adder.GRPCServer{}

	proto.RegisterAdderServer(s, sr)

	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("sercer")
	if err := s.Serve(l); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	/*	cha := make(chan int)
		<-cha*/
	//c := writeCH(1, 2, 34, 5, 6, 6)
	//readCH(c)
	//os.Exit(0)

	var wg sync.WaitGroup

	wg.Add(1)
	go g(&wg)
	wg.Wait()

	fmt.Println("main")
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
