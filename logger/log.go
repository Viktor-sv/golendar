package logger

import (
	"calendar/common"
	"fmt"
	lrus "github.com/sirupsen/logrus"
	"log"
	"os"
)

const (
	Info = iota
	Trace
	Warning
	Errors
)

type logger struct {
	file *os.File `json:"string"`
}

var L logger

// init fith func will be
//called at the init package time
func init() {
	L = logger{file: nil}
	// Only log the warning severity or above.
	lrus.SetOutput(os.Stdout)
	lrus.SetLevel(lrus.WarnLevel)

}

type tx int

func (l *logger) Info(msg ...string) {
	lrus.Warning("logrus warning ")
	lrus.Info("logrus info ")
	var x tx = 0x34
	fmt.Println(x)

	date, err := common.LocalTime("Europe/Kiev")
	if err != nil {
		//Warning.Println(":", err.Error())
	}

	path := date + "_test.log"
	l.file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		//Warning.Println("cannot open file", path)
	}
	defer l.file.Close() //todo check if it will not crash if 'f' is nil

	log.New(l.file, "INFO:\t", log.Ldate|log.Ltime).Println(msg)
}

func (l *logger) Warning(msg ...string) {
	date, err := common.LocalTime("Europe/Kiev")
	if err != nil {
		//Warning.Println(":", err.Error())
	}

	path := date + "_test.log"
	l.file, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		//Warning.Println("cannot open file", path)
	}
	defer l.file.Close() //todo check if it will not crash if 'f' is nil

	log.New(l.file, "WARNING:\t", log.Ldate|log.Ltime).Println(msg)
}

/*
func InitLogger(loglevel string) {
	date, err := common.LocalTime("Europe/Kiev")
	if err != nil {
		Warning.Println(":", err.Error())
	}

	path := date + "_test.log"
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		Warning.Println("cannot open file", path)
		f = nil
	}
	defer f.Close() //todo check if it will not crash if 'f' is nil

	Info = log.New(f, "INFO:\t", log.Ldate|log.Ltime)
	Info.Println("dfhdgfdhfghsgfhjdgfjdfg dgfdg")

	Warning = log.New(f, "WARNING:\t", log.Ldate|log.Ltime)
	Error = log.New(f, "ERROR:\t", log.Ldate|log.Ltime)
}*/
