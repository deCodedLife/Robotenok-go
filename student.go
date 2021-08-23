package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"strconv"
)

type Student struct {
	ID      int    `json:"id"`
	Active  int    `json:"active"`
	Name    string `json:"name"`
	Age     string `json:"age"`
	Phone   string `json:"phone"`
	Parents string `json:"parents"`
	Sex     int    `json:"sex"` // 0 male, 1 female
	Image   int `json:"image"`
}

func (s *Student) Init() {
	s.ID = -1
	s.Active = -1
	s.Name = ""
	s.Age = ""
	s.Phone = ""
	s.Parents = ""
	s.Sex = -1
	s.Image = -1
}

func (s Student) Add() error {
	var queryValues []interface{}

	queryValues = append(queryValues, s.Name)
	queryValues = append(queryValues, s.Age)
	queryValues = append(queryValues, s.Phone)
	queryValues = append(queryValues, s.Parents)
	queryValues = append(queryValues, s.Sex)
	queryValues = append(queryValues, s.Image)

	var query = "insert into robotenok.students (name, age, phone, parents, sex, image) values (?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(query, queryValues...)

	return err
}

func (s Student) Update() error {
	if s.ID == -1 {
		return errors.New("user id has wrong data")
	}

	var queryValues []interface{}

	// Wrote it separately because goland marked it as error -_(O_O|)_-
	var query = "update robotenok.students" + " set "
	var isFirst = true

	if s.Name != "" {
		query += " name = ?"
		queryValues = append(queryValues, s.Name)
		isFirst = false
	}

	if s.Age != "" {
		query += " age = ?"
		queryValues = append(queryValues, s.Age)
		isFirst = false
	}

	if s.Phone != "" {
		if isFirst == false {
			query += ","
		}

		query += " phone = ?"
		queryValues = append(queryValues, s.Phone)
		isFirst = false
	}

	if s.Parents != "" {
		if isFirst == false {
			query += ","
		}

		query += " parents = ?"
		queryValues = append(queryValues, s.Parents)
		isFirst = false
	}

	if s.Sex != -1 {
		if isFirst == false {
			query += ","
		}

		query += " sex = ?"
		queryValues = append(queryValues, s.Sex)
		isFirst = false
	}

	if s.Image != -1 {
		if isFirst == false {
			query += ","
		}

		query += " image= ?"
		queryValues = append(queryValues, s.Image)
	}

	query += " where id = " + strconv.Itoa(s.ID)

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	row := stmt.QueryRow(queryValues...)
	return row.Err()
}

func (s *Student) Remove() error {
	s.Active = 0
	return s.Update()
}


type Students struct {
	Students []Student `json:"students"`
}

func (s *Students) selectStudents(q Student) error {
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

	if q.Sex != -1 {
		query += " and sex = ?"
		queryValues = append(queryValues, q.Sex)
		isSearch = true
	}

	if q.Parents != "" {
		query += " and parents like '%" + template.HTMLEscapeString(q.Parents) + "%'"
		isSearch = true
	}

	if q.Phone != "" {
		query += " and phone like '%" + template.HTMLEscapeString(q.Phone) + "%'"
		isSearch = true
	}

	if q.Name != "" {
		query += " and name like '%" + template.HTMLEscapeString(q.Name) + "%'"
		isSearch = true
	}

	if q.Age != "" {
		query += " and age = ?"
		queryValues = append(queryValues, q.Age)
		isSearch = true
	}

	if q.Image != -1 {
		query += " and image = ?"
		queryValues = append(queryValues, q.Image)
		isSearch = true
	}

	if isSearch == false {
		return errors.New("nothing to do")
	}

	stpm, err := db.Prepare(query)
	defer stpm.Close()

	if err != nil {
		return err
	}

	row, err := stpm.Query(queryValues...)

	if err != nil {
		return err
	}

	for row.Next() {
		t := Student{}
		err := row.Scan(&t.ID, &t.Active, &t.Name, &t.Age, &t.Phone, &t.Parents, &t.Sex, &t.Image)

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

	SendData(w, 200, selectedStudents.Students)
}