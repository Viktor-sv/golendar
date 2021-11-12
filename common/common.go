package common

import (
	"fmt"
	"time"
)

func LocalTime(location string) error {
	var Tmp int32 = 0
	fmt.Println(Tmp)
	loc, err := time.LoadLocation(location)
	if err != nil {
		//panic(err)
		return err
	}

	t := time.Now().In(loc)
	fmt.Println("%t", t)
	fmt.Println(t)
	return nil
}
