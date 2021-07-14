package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type Student struct {
	ID      int    `json:"id"`
	Active  int    `json:"active"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Parents string `json:"parents"`
	Sex     int    `json:"sex"`
}

func (s *Student) Init() {
	s.ID = -1
	s.Active = -1
	s.Sex = -1
}

func (s Student) Add() error {
	var query string

	query = "insert into robotenok.students (name, phone, parents, sex) values (?,?,?,?)"
	_, err := db.Exec(query, s.Name, s.Phone, s.Parents, s.Sex)

	return err
}

func (s Student) Update() error {
	if s.ID == -1 {
		return errors.New("user id has wrong data")
	}

	var query string
	var isFirst bool

	// Wrote it separately because goland marked it as error -_(O_O|)_-
	query = "update robotenok.students" + " set "
	isFirst = true

	if s.Name != "" {
		query += " name like %" + s.Name + "%"
		isFirst = false
	}

	if s.Phone != "" {
		if isFirst == false {
			query += ","
		}

		query += " phone like %" + s.Phone + "%"
		isFirst = false
	}

	if s.Parents != "" {
		if isFirst == false {
			query += ","
		}

		query += " parents = " + s.Parents
		isFirst = false
	}

	if s.Sex != -1 {
		if isFirst == false {
			query += ","
		}

		query += " sex = " + strconv.Itoa(s.Sex)
		isFirst = false
	}

	query += " where id = " + strconv.Itoa(s.ID)

	_, err := db.Exec(query)
	return err
}

func (s *Student) Remove() error {
	s.Active = 0
	return s.Update()
}


type Students struct {
	Students []Student `json:"students"`
}

func (s *Students) selectStudents(q Student) error {
	var query string
	var isSearch bool

	isSearch = false
	query = "select * from robotenok.students" + " where "

	if q.Active != -1 {
		query += "active = " + strconv.Itoa(q.Active)
		isSearch = true
	} else {
		query += "active = 1"
	}

	if q.ID != -1 {
		query += " and id = " + strconv.Itoa(q.ID)
		isSearch = true
	}

	if q.Sex != -1 {
		query += " and sex = " + strconv.Itoa(q.Sex)
		isSearch = true
	}

	if q.Parents != "" {
		query += " and parents like '%" + q.Parents + "%'"
		isSearch = true
	}

	if q.Phone != "" {
		query += " and phone like '%" + q.Phone + "%'"
		isSearch = true
	}

	if q.Name != "" {
		query += " and name like '%" + q.Name + "%'"
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
		t := Student{}
		err := row.Scan(&t.ID, &t.Active, &t.Name, &t.Phone, &t.Parents, &t.Sex)

		if err != nil {
			return err
		}

		s.Students = append(s.Students, t)
	}

	return nil
}

func AddStudent(w http.ResponseWriter, r *http.Request) {
	var request Request
	var newStudent Student

	defer LogHandler("student add")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &newStudent)
	HandleError(err, w, WrongDataError)

	err = newStudent.Add()
	HandleError(err, w, UnknownError)

	SendData(w, 200, newStudent)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var request Request
	var updatingStudent Student

	defer LogHandler("student update")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	updatingStudent.Init()
	err = json.Unmarshal(textJson, &updatingStudent)
	HandleError(err, w, WrongDataError)

	err = updatingStudent.Update()
	HandleError(err, w, UnknownError)

	SendData(w, 200, updatingStudent)
}

func RemoveStudent (w http.ResponseWriter, r *http.Request) {
	var request Request
	var removingStudent Student

	defer LogHandler("student remove")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &removingStudent)
	HandleError(err, w, WrongDataError)

	err = removingStudent.Remove()
	HandleError(err, w, UnknownError)

	SendData(w, 200, removingStudent)
}

func SelectStudents (w http.ResponseWriter, r *http.Request) {
	var request Request
	var searchingStudent Student
	var selectedStudents Students

	defer LogHandler("student select")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	searchingStudent.Init()
	err = json.Unmarshal(textJson, &searchingStudent)
	HandleError(err, w, WrongDataError)

	err = selectedStudents.selectStudents(searchingStudent)
	HandleError(err, w, UnknownError)

	SendData(w, 200, selectedStudents)
}