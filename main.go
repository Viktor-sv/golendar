package main

import (
	"calendar/adder"
	"calendar/api/proto"
	"calendar/db"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"path/filepath"
	"time"

	"bufio"
	"google.golang.org/grpc"
	"os"
	"strconv"
	"sync"

	"calendar/logger"
	"calendar/srv"
	"io/ioutil"
	"runtime"
	"strings"
)

var Version string

type myS string
type myMapType map[string]int32

var fileWordCnt myMapType

type sl []map[string]int32

func (s sl) Read([]byte) (int, error) {

	for _, v := range s {
		fmt.Println(v)
	}

	return 1, nil
}

type getter interface {
	GetOne(f string) (int32, error)
}

type getterAll interface {
	GetOne(f string) (int32, error)
	GetAll() ([]map[string]int32, error)
}

func (s *myMapType) GetOne(f string) (int32, error) {
	return 0, nil
}

func (s *myMapType) GetAll() ([]map[string]int32, error) {
	return nil, nil
}

func (s myS) GetOne(f string) (int32, error) {

	return 0, nil
}

func (s myS) Read([]byte) (int, error) {

	return 0, nil
}

func (s myS) GetAll() ([]map[string]int32, error) {
	return nil, nil
}

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

func gRPC(wg *sync.WaitGroup) {
	defer wg.Done()

	s := grpc.NewServer()
	sr := &adder.GRPCServer{}

	proto.RegisterAdderServer(s, sr)

	l, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("gRPC server started!")
	if err := s.Serve(l); err != nil {
		log.Fatalln(err)
	}
}

func apiServer(w *sync.WaitGroup) {
	defer w.Done()
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

func getNews() []int {
	var tmpSlice []int
	//sl.connect
	//sql query ("select * from tebale")
	for i := 0; i < 10; i++ {
		tmpSlice = append(tmpSlice, i)

	}
	return tmpSlice
}

func f4() {
	sl := getNews()
	fmt.Println(sl)

}

func main() {
	//client(sl)
	//server()
	fmt.Println("main")
	fmt.Print(Version)
	fmt.Println("threads:", runtime.GOMAXPROCS(-1))
	//runTwoGo()

	var wg sync.WaitGroup
	wg.Add(1)
	go gRPC(&wg)
	wg.Add(1)
	go apiServer(&wg)
	wg.Wait()

	/*	var int getter
		var s myS
		var mmap *myMapType

		int = mmap
		int = s

		if int != nil {

		}*/

	/*println("main started")
	logWc()
	println("main finished")*/

	// ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// Cancel even if everything goes fine without an error to release resources.
	// defer cancel()
	// ch := make(chan int)
	// go doSomething(ch)

	/*	var wg sync.WaitGroup
		wg.Add(1)
		go f1(&wg) //wg.Done()
		go f2()    //wg.Done() -1
		go f3()    //wg.Done() -1

		wg.Wait()
	*/
	/*	go func() {
			defer wg.Done() //-1
			select {
			case <-ctx.Done():
				fmt.Println("TIMEOUT:", ctx.Err())
			case t := <-ch:
				fmt.Printf("JOB DONE in %d seconds\n", t)
			}
		}()
		time.Sleep(time.Duration(2) * time.Second)
		cancel()
		time.Sleep(time.Duration(2) * time.Second)
	*/
	/*
		var dialAddr = flag.String("dial", "localhost:8000", "host:port to dial")

		type Message struct {
			Body string
		}

		flag.Parse()

		conn, err := net.Dial("tcp", *dialAddr)

		if err != nil {
			fmt.Println("connan connect to ", conn.RemoteAddr().String())
		}

		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			m := Message{Body: s.Text()}
			bytes, err := json.Marshal(m)
			if err != nil {
				log.Fatal(err)
			}

			conn.Write(bytes)
		}

		if err := s.Err(); err != nil {
			log.Fatal(err)
		}
	*/
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(" *wg go go go")
}

func f2() {
	//defer wg.Done()
	fmt.Println("wg go go go")
}

func f3() {
	//defer wg.Done()
	fmt.Println("wg go go go")
}

func doSomething(ch chan<- int) {
	// Prevent picking up the same random number all the time for sleeping.
	rand.Seed(time.Now().UnixNano())

	// Pick a random number to simulate time it takes to finish the job.
	delay := rand.Intn(5)
	fmt.Printf("RUN: %d seconds\n", delay)
	time.Sleep(time.Duration(8) * time.Second)

	ch <- delay
}

func logWc() {
	fileWordCnt = make(map[string]int32, 0)
	_ = getFilesCnt(".")

	var wg sync.WaitGroup
	for i, _ := range fileWordCnt {
		wg.Add(1)
		go countWords(&wg, i)
	}
	wg.Wait()

	for i, v := range fileWordCnt {
		fmt.Println(v, " wc[", i, "]")
	}

	var s sl = []map[string]int32{fileWordCnt, fileWordCnt}
	r := bufio.NewReader(s)
	r.ReadByte()
	//sort.Strings(fileWordCnt)
}

var mx sync.Mutex

func countWords(wg *sync.WaitGroup, path string) error {
	defer wg.Done()

	t, err := os.Open(path)
	if err != nil {
		fmt.Printf("Could not open file '%v': %v", path, err)
		return err
	}

	reader := bufio.NewReader(t)
	line, _ := reader.ReadString('\\')
	words := strings.Fields(line)

	mx.Lock()
	fileWordCnt[path] = int32(len(words))
	mx.Unlock()

	//println("-", line)

	return nil
}

func getFilesCnt(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println("err", err.Error())
		return err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".log" {
			fileWordCnt[file.Name()] = -1
		}
	}

	return nil
}

var c = make(chan *int32, 2)
var data = []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func runTwoGo() {

	var wg sync.WaitGroup
	wg.Add(2)
	go save(&wg)
	go read(&wg)
	wg.Wait()
}

func save(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("readdddd:", &data[0])
	fmt.Println("readdddd:", &data[1])

	for _, val := range data {
		//	tmp := val
		c <- &val //2
		//fmt.Println("read:")

		go func(v *int32) {
			fmt.Println(*v) //2

		}(&val)
	}
}

func read(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < len(data); i++ {
		val := <-c
		fmt.Println("read:", *val)
	}
}
