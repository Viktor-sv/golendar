package db

import (
	"calendar/model"
	"fmt"
)

func init() {
	//addTwoNum(2, 5)
}

func addTwoNum(x int, y int) int {
	fmt.Println(x + y)
	return x + y

}

func GetEvent(id string) *model.Event {
	db, err := connectToDB()
	if err != nil {
		fmt.Println("connect to db error ", err.Error())
	}

	res, err := db.Query("SELECT * FROM evenets where id=?", id)
	if err != nil {
		fmt.Println("select error ", err.Error())
	}

	var event model.Event
	for res.Next() {
		err = res.Scan(&event)
		if err != nil {
			fmt.Println("select error ", err.Error())
			continue
		}

		if event.Id == id {
			return &event
		}
	}
	return &event
}

func AddE(e model.Event) error {
	db, err := connectToDB()
	if err != nil {
		fmt.Println("connect to db error ", err.Error())
	}

	_, err = db.Exec(`insert into users(name, pass) values(?, ?,?,,,)`, e.Id, e.Name)
	if err != nil {
		fmt.Println("connect to db error ", err.Error())
	}

	return err
}
