package main

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
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

func (u *User) Init() {
	u.ID = -1
	u.Active = -1
}

func (u *User) Select() error {
	var queryValues []interface{}

	var isSearch = false
 	var query = "select * from users" + " where "

	if u.Active != -1 {
		query += "active = ?"
		queryValues = append(queryValues, u.Active)
		isSearch = true
	} else {
		query += "active = 1"
	}

	if u.ID != -1 && u.ID != 0 {
		query += " and id = ?"
		queryValues = append(queryValues, u.ID)
		isSearch = true
	}

	if u.Name != "" {
		query += " and name like '%" + template.HTMLEscapeString(u.Name) + "%'"
		isSearch = true
	}

	if u.Login != "" {
		query += " and login = ?"
		queryValues = append(queryValues, u.Login)
		isSearch = true
	}

	if isSearch == false {
		return errors.New("nothing to do")
	}

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	row := stmt.QueryRow(queryValues...)

	if row.Err() != nil {
		return row.Err()
	}

	err = row.Scan(&u.ID, &u.Name, &u.Login, &u.Password, &u.Image, &u.UserType, &u.Active)

	return err
}

func (u User) Add(hash Device) error {
	if hash.Active != 1 {
		return errors.New("hash was already used before")
	}

	hash.Active = 0
	err := hash.StatusChange()

	if err != nil {
		return err
	}

	var query string
	var queryValues []interface{}

	queryValues = append(queryValues, u.Name)
	queryValues = append(queryValues, GenString(8))
	queryValues = append(queryValues, ToSHA512(u.Password))
	queryValues = append(queryValues, u.UserType)

	query = "insert into robotenok.users (name, login, password, user_type) values (?, ?, ?, ?)"
	_, err = db.Query(query, queryValues...)

	return err
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
	var queryValues []interface{}

	query = "update robotenok.users set active = 1"

	if u.UserType != "" {
		query += ", user_type = ?"
		queryValues = append(queryValues, u.UserType)
	}

	if u.Password != "" {
		query += ", password = ?"
		queryValues = append(queryValues, u.Password)
	}

	if u.Login != "" {
		query += ", login = ?"
		queryValues = append(queryValues, u.Login)
	}

	if u.Name != "" {
		query += ", name = ?"
		queryValues = append(queryValues, u.Name)
	}

	if u.Secret != "" {
		query += ", secret = ?"
		queryValues = append(queryValues, u.Secret)
	}

	if u.Image != "" {
		query += ", image = ?"
		queryValues = append(queryValues, u.Image)
	}

	query += " where id = " + "?"
	queryValues = append(queryValues, u.ID)

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(queryValues...)

	return err
}

type Users struct {
	Users []User
}

func (u *Users) Select(q User) error {
	var queryValues []interface{}

	var isSearch = false
	var query = "select * from robotenok.users" + " where "

	if q.Active != -1 {
		query += "active = ?"
		queryValues = append(queryValues, q.Active)
		isSearch = true
	} else {
		query += "active = 1"
	}

	if q.Name != "" {
		query += " and name like '%" + template.HTMLEscapeString(q.Name) + "%'"
		isSearch = true
	}

	if q.Login != "" {
		query += " and login like '%" + template.HTMLEscapeString(q.Login) + "%'"
		isSearch = true
	}

	if q.UserType != "" {
		query += " and user_type = ?"
		queryValues = append(queryValues, q.UserType)
		isSearch = true
	}

	if q.ID != -1 {
		query += " and id = ?"
		queryValues = append(queryValues, q.ID)
		isSearch = true
	}

	if isSearch == false {
		return errors.New("nothing to do")
	}

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	row, err := stmt.Query(query)

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
	var device Device

	defer LogHandler("User add")

	hash := mux.Vars(r)["hash"]

	if hash == "" {
		SendData(w, 400, WrongDataError.Description)
		return
	}

	device.Init()
	device.Hash = hash

	err := device.Get()
	HandleError(err, w, SecurityError)

	err = requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &newUser)
	HandleError(err, w, WrongDataError)

	err = newUser.Add(device)
	HandleError(err, w, UnknownError)

	SendData(w, 200, newUser)
}

func UpdateUser (w http.ResponseWriter, r *http.Request) {
	var request Request
	var updateUser User

	defer LogHandler("user update")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	updateUser.Init()
	err = json.Unmarshal(textJson, &updateUser)
	HandleError(err, w, WrongDataError)

	if request.UserID != updateUser.ID {
		SendData(w, 401, "You can't change other user data")
		return
	}

	err = updateUser.Update()
	HandleError(err, w, UnknownError)

	SendData(w, 200, updateUser)
}

func RemoveUser (w http.ResponseWriter, r *http.Request) {
	var request Request
	var removingUser User

	defer LogHandler("user remove")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 0)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &removingUser)
	HandleError(err, w, WrongDataError)

	err = removingUser.Remove()
	HandleError(err, w, UnknownError)

	SendData(w, 200, removingUser)
}

func SelectUser(w http.ResponseWriter, r *http.Request) {
	var request Request
	var searchedUser User
	var selectedUsers Users

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &searchedUser)
	HandleError(err, w, WrongDataError)

	requestLevel := 2

	if searchedUser.ID != request.UserID {
		requestLevel = 0
	}

	err = permCheck(request.UserID, requestLevel)
	HandleError(err, w, SecurityError)

	err = selectedUsers.Select(searchedUser)
	HandleError(err, w, UnknownError)

	SendData(w, 200, selectedUsers.Users)
}