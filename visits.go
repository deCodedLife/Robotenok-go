package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Visit struct {
	ID        int  `json:"id"`
	Active    int   `json:"active"`
	StudentID int  `json:"student_id"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	Type      string `json:"type"`
}

func (v *Visit) Init() {
	v.ID = -1
	v.Active = -1
	v.StudentID = -1
	v.Date = ""
	v.Time = ""
	v.Type = ""
}

func (v Visit) Add() error {
	var query string
	var queryValues []interface{}

	queryValues = append(queryValues, v.StudentID)
	queryValues = append(queryValues, GetDate())
	queryValues = append(queryValues, GetTime())
	queryValues = append(queryValues, v.Type)

	query = "insert into robotenok.visits (student_id, date, time, type) values (?, ?, ?, ?)"

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(queryValues...)

	return err
}

func (v Visit) Update() error {
	if v.ID == -1 {
		return errors.New("visit id has wrong data")
	}

	var query string
	var isFirst bool
	var queryValues []interface{}

	// Wrote it separately because goland marked it as error -_(O_O|)_-
	query = "update robotenok.visits" + " set "
	isFirst = true

	if v.Active != -1 {
		query += "active = ?"
		queryValues = append(queryValues, v.Active)
		isFirst = false
	}

	if v.StudentID != -1 {
		if isFirst == false {
			query += ","
		}

		query += " student_id = ?"
		queryValues = append(queryValues, v.StudentID)
		isFirst = false
	}

	if v.Date != "" {
		if isFirst == false {
			query += ","
		}

		query += " date = ?"
		queryValues = append(queryValues, v.Date)
		isFirst = false
	}

	if v.Time != "" {
		if isFirst == false {
			query += ","
		}

		query += " time = ?"
		queryValues = append(queryValues, v.Time)
		isFirst = false
	}

	if v.Type != "" {
		if isFirst == false {
			query += ","
		}

		query += " type = ?"
		queryValues = append(queryValues, v.Type)
		isFirst = false
	}

	query += " where id = " + "?"
	queryValues = append(queryValues, v.ID)

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(queryValues...)
	return err
}

func (v *Visit) Remove() error {
	v.Active = 0
	return v.Update()
}

type Visits struct {
	Visits []Visit `json:"visits"`
}

func (v* Visits) Select(q Visit) error {
	var query string
	var isSearch bool
	var queryValues []interface{}

	isSearch = false
	query = "select * from robotenok.visits" + " where "

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

	if q.StudentID != -1 {
		query += " and student_id = ?"
		queryValues = append(queryValues, q.StudentID)
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

	if q.Type != "" {
		query += " and type = ?"
		queryValues = append(queryValues, q.Type)
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

	row, err := stmt.Query(queryValues...)

	if err != nil {
		return err
	}

	for row.Next() {
		t := Visit{}
		err := row.Scan(&t.ID, &t.Active, &t.StudentID, &t.Date, &t.Time, &t.Type)

		if err != nil {
			return err
		}

		v.Visits = append(v.Visits, t)
	}

	return nil
}

func AddVisit(w http.ResponseWriter, r *http.Request) {
	var request Request
	var newVisit Visit

	defer LogHandler("visit add")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &newVisit)
	HandleError(err, w, WrongDataError)

	err = newVisit.Add()
	HandleError(err, w, UnknownError)

	SendData(w, 200, newVisit)
}

func UpdateVisit(w http.ResponseWriter, r *http.Request) {
	var request Request
	var updatingVisit Visit

	defer LogHandler("visit update")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	updatingVisit.Init()
	err = json.Unmarshal(textJson, &updatingVisit)
	HandleError(err, w, WrongDataError)

	err = updatingVisit.Update()
	HandleError(err, w, UnknownError)

	SendData(w, 200, updatingVisit)
}

func RemoveVisit(w http.ResponseWriter, r *http.Request) {
	var request Request
	var removingVisit Visit

	defer LogHandler("visit update")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &removingVisit)
	HandleError(err, w, WrongDataError)

	err = removingVisit.Remove()
	HandleError(err, w, UnknownError)

	SendData(w, 200, removingVisit)
}

func SelectVisits(w http.ResponseWriter, r *http.Request) {
	var request Request
	var searcherVisits Visit
	var selectedVisits Visits

	defer LogHandler("student select")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	searcherVisits.Init()
	err = json.Unmarshal(textJson, &searcherVisits)
	HandleError(err, w, WrongDataError)

	err = selectedVisits.Select(searcherVisits)
	HandleError(err, w, UnknownError)

	SendData(w, 200, selectedVisits)
}