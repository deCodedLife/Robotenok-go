package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type Class struct {
	ID          int    `json:"id"`
	Active      int    `json:"active"`
	UserID      int    `json:"user_id"`
	GroupID     int    `json:"group_id"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	HostAddress string `json:"host_address"`
}

func (c Class) Init() {
	c.ID = -1
	c.Active = -1
	c.UserID = -1
	c.GroupID = -1
	c.Date = ""
	c.Time = ""
	c.HostAddress = ""
}

func (c Class) Add() error {
	var queryValues []interface{}

	queryValues = append(queryValues, c.UserID)
	queryValues = append(queryValues, c.GroupID)
	queryValues = append(queryValues, GetDate())
	queryValues = append(queryValues, GetTime())
	queryValues = append(queryValues, c.HostAddress)

	var query = "insert into robotenok.classes (user_id, group_id, date, time, host_address) values (?, ?, ?, ?, ?)"

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(query, c.UserID, c.GroupID, GetDate(), GetTime(), c.HostAddress)

	return err
}

func (c Class) Update() error {
	if c.ID == -1 {
		return errors.New("class id has wrong data")
	}

	var queryValues []interface{}

	// Wrote it separately because goland marked it as error -_(O_O|)_-
	var query = "update robotenok.classes" + " set "
	var isFirst = true

	if c.UserID != -1 {
		query += " user_id = ?"
		queryValues = append(queryValues, c.UserID)
		isFirst = false
	}

	if c.GroupID != -1 {
		if isFirst == false {
			query += ","
		}

		query += " group_id = ?"
		queryValues = append(queryValues, c.GroupID)
		isFirst = false
	}

	if c.Date != "" {
		if isFirst == false {
			query += ","
		}

		query += " date = ?"
		queryValues = append(queryValues, c.Date)
		isFirst = false
	}
	if c.Time != "" {
		if isFirst == false {
			query += ","
		}

		query += " time = ?"
		queryValues = append(queryValues, c.Time)
		isFirst = false
	}
	if c.HostAddress != "" {
		if isFirst == false {
			query += ","
		}

		query += " host_address = ?"
		queryValues = append(queryValues, c.HostAddress)
		isFirst = false
	}

	query += " where id = " + strconv.Itoa(c.ID)

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(queryValues...)
	return err
}

func (c *Class) Remove() error {
	c.Active = 0
	return c.Update()
}

type Classes struct {
	Classes []Class `json:"classes"`
}

func (c *Classes) Select(q Class) error {
	var queryValues []interface{}

	var isSearch = false
	var query = "select * from robotenok.students" + " where "

	if q.Active != -1 {
		query += "active = ?"
		queryValues = append(queryValues, q.Active)
		isSearch = true
	} else {
		query += "active = 1"
	}

	if q.ID != -1 {
		query += " and id = ?"
		queryValues = append(queryValues, q.ID)
		isSearch = true
	}

	if q.UserID != -1 {
		query += " and user_id = ?"
		queryValues = append(queryValues, q.UserID)
		isSearch = true
	}

	if q.GroupID != -1 {
		query += " and group_id = ?"
		queryValues = append(queryValues, q.GroupID)
		isSearch = true
	}

	if q.Date != "" {
		query += " and date = ?"
		queryValues = append(queryValues, q.Date)
		isSearch = true
	}

	if q.Time != "" {
		query += " and time = ?"
		queryValues = append(queryValues, q.Time)
		isSearch = true
	}

	if q.HostAddress != "" {
		query += " and host_address = ?"
		queryValues = append(queryValues, q.HostAddress)
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
		t := Class{}
		err := row.Scan(&t.ID, &t.Active, &t.UserID, &t.GroupID, &t.Date, &t.Time, &t.HostAddress)

		if err != nil {
			return err
		}

		c.Classes = append(c.Classes, t)
	}

	return nil
}

func AddClass(w http.ResponseWriter, r *http.Request) {
	var request Request
	var newClass Class

	defer LogHandler("class add")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &newClass)
	HandleError(err, w, WrongDataError)

	err = newClass.Add()
	HandleError(err, w, UnknownError)

	SendData(w, 200, newClass)
}

func UpdateClass(w http.ResponseWriter, r *http.Request) {
	var request Request
	var updatingClass Class

	defer LogHandler("class update")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &updatingClass)
	HandleError(err, w, WrongDataError)

	err = updatingClass.Update()
	HandleError(err, w, UnknownError)

	SendData(w, 200, updatingClass)
}

func RemoveClass(w http.ResponseWriter, r *http.Request) {
	var request Request
	var removingClass Class

	defer LogHandler("class remove")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &removingClass)
	HandleError(err, w, WrongDataError)

	err = removingClass.Remove()
	HandleError(err, w, UnknownError)

	SendData(w, 200, removingClass)
}

func SelectClasses(w http.ResponseWriter, r *http.Request) {
	var request Request
	var searchingClass Class
	var selectedClasses Classes

	defer LogHandler("classes select")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	searchingClass.Init()
	err = json.Unmarshal(textJson, &searchingClass)
	HandleError(err, w, WrongDataError)

	err = selectedClasses.Select(searchingClass)
	HandleError(err, w, UnknownError)

	SendData(w, 200, selectedClasses.Classes)
}
