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

func (s *Student) init() {
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
		query += "name = " + s.Name
		isFirst = false
	}

	if s.Phone != "" {
		if isFirst == false {
			query += ","
		}

		query += "phone = " + s.Phone
		isFirst = false
	}

	if s.Parents != "" {
		if isFirst == false {
			query += ","
		}

		query += "parents = " + s.Parents
		isFirst = false
	}

	if s.Sex != -1 {
		if isFirst == false {
			query += ","
		}

		query += "sex = " + strconv.Itoa(s.Sex)
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

	err = permCheck(request.UserID, 1)
	
	if err != nil {
		SendData(w, 401, err.Error())
		return
	}

	textJson, err := json.Marshal(request.Body)

	if err != nil {
		SendData(w, 400, err.Error())
		return
	}

	err = json.Unmarshal(textJson, &newStudent)

	if err != nil {
		SendData(w, 400, err.Error())
		return
	}

	err = newStudent.Add()

	if err !=nil {
		SendData(w, 500, err.Error())
	}
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var request Request
	var updatingStudent Student

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

	err = permCheck(request.UserID, 1)

	if err != nil {
		SendData(w, 401, err.Error())
		return
	}

	textJson, err := json.Marshal(request.Body)

	if err != nil {
		SendData(w, 400, err.Error())
		return
	}

	err = json.Unmarshal(textJson, &updatingStudent)

	if err != nil {
		SendData(w, 400, err.Error())
		return
	}

	err = updatingStudent.Update()

	if err != nil {
		SendData(w, 500, err.Error())
	}
}

func RemoveStudent (w http.ResponseWriter, r *http.Request) {
	var request Request
	var removingStudent Student

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

	err = permCheck(request.UserID, 1)

	if err != nil {
		SendData(w, 401, err.Error())
		return
	}

	textJson, err := json.Marshal(request.Body)

	if err != nil {
		SendData(w, 400, err.Error())
		return
	}

	err = json.Unmarshal(textJson, &removingStudent)

	if err != nil {
		SendData(w, 400, err.Error())
		return
	}

	err = removingStudent.Remove()

	if err != nil {
		SendData(w, 500, err.Error())
	}
}

func SelectStudent (w http.ResponseWriter, r *http.Request) {
	var request Request
	var searchedStudent Student
	var selectedStudents Students

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

	err = permCheck(request.UserID, 1)

	if err != nil {
		SendData(w, 401, err.Error())
		return
	}

	textJson, err := json.Marshal(request.Body)

	if err != nil {
		SendData(w, 400, err.Error())
		return
	}

	searchedStudent.init()
	err = json.Unmarshal(textJson, &searchedStudent)

	if err != nil {
		SendData(w, 400, "Client send a wrong or empty data")
		return
	}

	err = selectedStudents.selectStudents(searchedStudent)

	if err != nil {
		SendData(w, 500, err.Error())
	}

	SendData(w, 200, selectedStudents)
}