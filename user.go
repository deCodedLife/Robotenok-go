package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"
)

var userTypes = []string{"admin", "teacher", "client"}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Image    string `json:"image"`
	UserType string `json:"user_type"`
	Active   int    `json:"active"`
	Secret   string `json:"secret"`
	Online   time.Time
}

func (u *User) init() {
	u.ID = -1
}

func (u *User) Select() error {
	var query string
	var isSearch bool

	isSearch = false
 	query = "select * from users where active = 1"

	if u.ID != -1 && u.ID != 0 {
		query += " and id = " + strconv.Itoa(u.ID)
		isSearch = true
	}

	if u.Name != "" {
		query += " and name like '%" + u.Name + "%'"
		isSearch = true
	}

	if u.Login != "" {
		query += " and login = '" + u.Login + "'"
		isSearch = true
	}

	if isSearch == false {
		return errors.New("nothing to do")
	}

	row := db.QueryRow(query)

	if row.Err() != nil {
		return row.Err()
	}

	err := row.Scan(&u.ID, &u.Name, &u.Login, &u.Password, &u.Image, &u.UserType, &u.Active)
	return err
}

func (u User) Add(hash ConfirmedDevices) error {
	if hash.Active != 1 {
		return errors.New("hash was already used before")
	}

	hash.Active = 0
	err := hash.StatusChange()

	if err != nil {
		return err
	}

	var query string

	query = "insert into robotenok.users (name, login, password) values (?, ?, ?)"
	_, err = db.Query(query, u.Name, GenString(16), ToSHA512(u.Password))

	if err != nil {
		return err
	}

	return nil
}

func (u User) Remove() error {
	_, err := db.Query("delete from robotenok.users where id = ?", u.ID)
	return err
}

func (u User) Update() error {
	if u.ID == -1 {
		return errors.New("user id has wrong data")
	}

	var query string

	//noinspection SqlWithoutWhere
	query = "update robotenok.users set active = 1"

	if u.UserType != "" {
		query += ", user_type = " + u.UserType
	}

	if u.Password != "" {
		query += ", password = " + u.Password
	}

	if u.Login != "" {
		query += ", login = " + u.Login
	}

	if u.Name != "" {
		query += ", name = " + u.Name
	}

	if u.Secret != "" {
		query += ", secret = " + u.Secret
	}

	if u.Image != "" {
		query += ", image = " + u.Image
	}

	query += " where id = " + strconv.Itoa(u.ID)
	_, err := db.Exec(query)

	return err
}

type Users struct {
	Users []User
}

func (u *Users) Select(q User) error {
	var query string
	var isSearch bool

	isSearch = false
	query = "select * from robotenok.users" + " where "

	if q.Active != -1 {
		query += "active = " + strconv.Itoa(q.Active)
		isSearch = true
	} else {
		query += "active = 1"
	}

	if q.Name != "" {
		query += " and name like '%" + q.Name + "%'"
		isSearch = true
	}

	if q.Login != "" {
		query += " and login like '%" + q.Login + "%'"
		isSearch = true
	}

	if q.UserType != "" {
		query += " and user_type = " + q.UserType
		isSearch = true
	}

	if q.ID != -1 {
		query += " and id = " + strconv.Itoa(q.ID)
		isSearch = true
	}

	if isSearch == false {
		return errors.New("nothing to do")
	}

	row, err := db.Query(query)

	if err != nil {
		return err
	}

	for row.Next() {
		t := User{}
		err := row.Scan(&t.ID, &t.Name, &t.Login, &t.Password, &t.Image, &t.UserType, &t.Active)

		if err != nil {
			return err
		}

		u.Users = append(u.Users, t)
	}

	return nil
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	var request Request
	var newUser User
	var device ConfirmedDevices

	hash := r.URL.Query().Get("hash")

	if len(hash) > 13 {
		err := device.Get(hash)

		if err != nil {
			SendData(w, 401, "security error")
		}
	} else {
		id, err := strconv.Atoi(hash)

		if err != nil {
			SendData(w, 400, "client sent wrong data")
		}

		err = device.Get(id)

		if err != nil {
			SendData(w, 401, "security error")
		}
	}

	if hash == "" {
		SendData(w, 400, "can't read hash value")
		return
	}

	err := requestHandler(&request, r)

	if err != nil {
		SendData(w, 400, "Client send a wrong or empty data")
		return
	}

	err = request.checkToken()

	if err != nil {
		SendData(w, 401, err.Error())
		return
	}

	textJson, err := json.Marshal(request.Body)

	if err != nil {
		SendData(w, 400, err.Error())
		return
	}

	err = json.Unmarshal(textJson, &newUser)

	if err != nil {
		SendData(w, 400, "Client send a wrong or empty data")
		return
	}

	err = newUser.Add(device)

	if err != nil {
		log.Println(err.Error())
		SendData(w, 500, "Something went wrong. " + err.Error())
	}
}

func UpdateUser (w http.ResponseWriter, r *http.Request) {
	var request Request
	var updateUser User

	err := requestHandler(&request, r)

	if err != nil {
		SendData(w, 400, "Client send a wrong or empty data")
		return
	}

	err = request.checkToken()

	if err != nil {
		SendData(w, 401, err.Error())
		return
	}

	textJson, err := json.Marshal(request.Body)

	if err != nil {
		SendData(w, 400, err.Error())
		return
	}

	err = json.Unmarshal(textJson, &updateUser)

	if err != nil {
		SendData(w, 400, "Client send a wrong or empty data")
		return
	}

	if request.UserID != updateUser.ID {
		SendData(w, 401, "you can't change other user data")
		return
	}

	err = updateUser.Update()

	if err != nil {
		SendData(w, 500, "can't update user. Reason: " + err.Error())
		return
	}
}

func RemoveUser (w http.ResponseWriter, r *http.Request) {
	var request Request
	var removingUser User

	err := requestHandler(&request, r)

	if err != nil {
		SendData(w, 400, "Client send a wrong or empty data")
		return
	}

	err = request.checkToken()

	if err != nil {
		SendData(w, 401, err.Error())
		return
	}

	err = permCheck(request.UserID, 0)

	if err != nil {
		SendData(w, 401, err.Error())
		return
	}

	textJson, err := json.Marshal(request.Body)

	if err != nil {
		SendData(w, 400, err.Error())
		return
	}

	err = json.Unmarshal(textJson, &removingUser)

	if err != nil {
		SendData(w, 400, "Client send a wrong or empty data")
		return
	}

	err = removingUser.Remove()

	if err != nil {
		SendData(w, 500, err.Error())
	}
}

func SelectUser(w http.ResponseWriter, r *http.Request) {
	var request Request
	var searchedUser User
	var selectedUsers Users

	err := requestHandler(&request, r)

	if err != nil {
		SendData(w, 400, "Client send a wrong or empty data")
		return
	}

	err = request.checkToken()

	if err != nil {
		SendData(w, 401, err.Error())
		return
	}

	textJson, err := json.Marshal(request.Body)

	if err != nil {
		SendData(w, 400, err.Error())
		return
	}

	err = json.Unmarshal(textJson, &searchedUser)

	if err != nil {
		SendData(w, 400, "Client send a wrong or empty data")
		return
	}

	requestLevel := 2

	if searchedUser.ID != request.UserID {
		requestLevel = 0
	}

	err = permCheck(request.UserID, requestLevel)

	if err != nil {
		SendData(w, 401, err.Error())
		return
	}

	err = selectedUsers.Select(searchedUser)

	if err != nil {
		SendData(w, 500, err.Error())
	}

	SendData(w, 200, selectedUsers)
}