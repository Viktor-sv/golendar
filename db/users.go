package db

import (
	"calendar/model"
	"fmt"
)

func GetUser(name string, pass string) *model.User {

	u := GetUsers()
	for _, v := range u {
		fmt.Println("GetUsers: ", v.Name, v.Pass, name, pass)
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

	res, err := db.Query("SELECT name,pass,token FROM users")
	if err != nil {
		fmt.Println("select error ", err.Error())
	}

	var users []model.User
	for res.Next() {
		var u model.User
		err = res.Scan(&u.Name, &u.Pass, &u.Token)
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

func LogOut(u *model.User) error {
	db, err := connectToDB()
	if err != nil {
		fmt.Println("connect to db error ", err.Error())
	}

	fmt.Println("user: ", u)
	res, err1 := db.Exec("update users set  token=? where name=? and token=?", "", u.Name, u.Token)
	fmt.Println(res)
	if err1 != nil {
		fmt.Println("connect to db error ", err.Error())
	}

	return err
}

func LoginUser(u model.User) error {
	db, err := connectToDB()
	if err != nil {
		fmt.Println("connect to db error ", err.Error())
	}

	_, err = db.Exec("update users set token=? where name=? and pass=?", u.Token, u.Name, u.Pass)
	if err != nil {
		fmt.Println("connect to db error ", err.Error())
	}

	fmt.Printf("use[%s] was logged in with token:[%s]\n", u.Name, u.Token)
	return nil
}

func Add(u model.User) error {
	db, err := connectToDB()
	if err != nil {
		fmt.Println("connect to db error ", err.Error())
	}

	_, err = db.Exec(`insert into users(name, pass) values(?, ?)`, u.Name, u.Pass)
	if err != nil {
		fmt.Println("connect to db error ", err.Error())
	}

	return err
}
