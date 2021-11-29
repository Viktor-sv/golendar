package db

import (
	"calendar/model"
	"fmt"
)

func GetUser(name string, pass string) *model.User {

	u := GetUsers()
	for _, v := range u {
		if v.Name == name && v.Pass == pass {
			return &v
			break
		}
	}

	return nil
}

func GetUsers() []model.User {

	db, err := connectToDB()
	if err != nil {
		fmt.Println("connect to db error ", err.Error())
	}

	res, err := db.Query("SELECT name,pass FROM users")
	if err != nil {
		fmt.Println("select error ", err.Error())
	}

	var users []model.User

	for res.Next() {
		var u model.User

		err = res.Scan(&u.Name, &u.Pass)
		if err != nil {
			fmt.Println("select error ", err.Error())
			continue
		}
		users = append(users, u)
		fmt.Println("sql select res", u)
	}

	return users
}

func updateUser(u *model.User) error {
	return nil
}

func delete(u *model.User) {

}
