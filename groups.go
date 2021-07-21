package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"text/template"
)

type GroupType struct {
	ID     int  `json:"id"`
	Active int   `json:"active"`
	Name   string `json:"name"`
}

func (g *GroupType) Init () {
	g.ID = -1
	g.Active = -1
	g.Name = ""
}

func (g GroupType) Add() error {
	var query string
	var queryValues []interface{}

	queryValues = append(queryValues, g.Name)

	query = "insert into robotenok.group_types (name) values ?"

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(queryValues...)

	return err
}

func (g GroupType) Update() error {
	if g.ID == -1 {
		return errors.New("group type id has wrong data")
	}

	var query string
	var isFirst bool
	var queryValues []interface{}

	// Wrote it separately because goland marked it as error -_(O_O|)_-
	query = "update robotenok.students" + " set "
	isFirst = true

	if g.Name != "" {
		query += "name = ?"
		queryValues = append(queryValues, g.Name)
		isFirst = false
	}

	if g.Active != -1 {
		if isFirst == false {
			query += ","
		}

		query += "active = ?"
		queryValues = append(queryValues, g.Active)
	}

	query += " where id = " + "?"
	queryValues = append(queryValues, g.ID)

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(queryValues...)
	return err
}

func (g *GroupType) Remove() error {
	g.Active = 0
	return g.Update()
}

type GroupTypes struct {
	GroupTypes []GroupType
}

func (g* GroupTypes) Select(q GroupType) error {
	var query string
	var isSearch bool
	var queryValues []interface{}

	isSearch = false
	query = "select * from robotenok.students" + " where "

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

	if q.Name != "" {
		query += " and name like = '%" + template.HTMLEscapeString(q.Name) + "%'"
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
		t := GroupType{}
		err := row.Scan(&t.ID, &t.Active, &t.Name)

		if err != nil {
			return err
		}

		g.GroupTypes = append(g.GroupTypes, t)
	}

	return nil
}

func AddGroupType(w http.ResponseWriter, r *http.Request) {
	var request Request
	var newGroupType GroupType

	defer LogHandler("grouptype add")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &newGroupType)
	HandleError(err, w, WrongDataError)

	err = newGroupType.Add()
	HandleError(err, w, UnknownError)

	SendData(w, 200, newGroupType)
}

func UpdateGroupType(w http.ResponseWriter, r *http.Request) {
	var request Request
	var updatingGroupType GroupType

	defer LogHandler("grouptype update")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	updatingGroupType.Init()
	err = json.Unmarshal(textJson, &updatingGroupType)
	HandleError(err, w, WrongDataError)

	err = updatingGroupType.Update()
	HandleError(err, w, UnknownError)

	SendData(w, 200, updatingGroupType)
}

func RemoveGroupType(w http.ResponseWriter, r *http.Request) {
	var request Request
	var removingGroupType GroupType

	defer LogHandler("grouptype remove")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &removingGroupType)
	HandleError(err, w, WrongDataError)

	removingGroupType.Remove()
	HandleError(err, w, UnknownError)

	SendData(w, 200, removingGroupType)
}

func SelectGroupTypes(w http.ResponseWriter, r *http.Request) {
	var request Request
	var searchingGroupType GroupType
	var selectedGroupTypes GroupTypes

	defer LogHandler("grouptype select")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	searchingGroupType.Init()
	err = json.Unmarshal(textJson, &searchingGroupType)
	HandleError(err, w, WrongDataError)

	err = selectedGroupTypes.Select(searchingGroupType)
	HandleError(err, w, UnknownError)

	SendData(w, 200, selectedGroupTypes)
}

type GroupStudent struct {
	ID        int `json:"id"`
	Active    int `json:"active"`
	GroupID   int `json:"group_id"`
	StudentID int `json:"student_id"`
}

func (g GroupStudent) Init() {
	g.ID = -1
	g.Active = -1
	g.GroupID = -1
	g.StudentID = -1
}

func (g GroupStudent) Add() error {
	var query string
	var queryValues []interface{}

	queryValues = append(queryValues, g.GroupID)
	queryValues = append(queryValues, g.StudentID)

	query = "insert into robotenok.group_students (group_id, student_id) values (?, ?)"

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(queryValues...)

	return err
}

func (g GroupStudent) Update() error {
	if g.ID == -1 {
		return errors.New("group students id has wrong data")
	}

	var query string
	var isFirst bool
	var queryValues []interface{}

	// Wrote it separately because goland marked it as error -_(O_O|)_-
	query = "update robotenok.group_students" + " set "
	isFirst = true

	if g.GroupID != -1 {
		query += " group_id = ?"
		queryValues = append(queryValues, g.GroupID)
		isFirst = false
	}

	if g.StudentID != -1 {
		if isFirst == false {
			query += ","
		}

		query += " student_id = ?"
		queryValues = append(queryValues, g.StudentID)
		isFirst = false
	}

	query += " where id = " + "?"
	queryValues = append(queryValues, g.ID)

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(queryValues...)
	return err
}

func (g *GroupStudent) Remove() error {
	g.Active = 0
	return g.Update()
}

type GroupStudents struct {
	GroupStudents []GroupStudent `json:"groups_students"`
}

func (g *GroupStudents) Select(q GroupStudent) error {
	var query string
	var isSearch bool
	var queryValues []interface{}

	isSearch = false
	query = "select * from robotenok.group_students" + " where "

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

	if q.GroupID != -1 {
		query += " and group_id = ?"
		queryValues = append(queryValues, q.GroupID)
		isSearch = true
	}

	if q.StudentID != -1 {
		query += " and student_id = ?"
		queryValues = append(queryValues, q.StudentID)
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
		t := GroupStudent{}
		err := row.Scan(&t.ID, &t.Active, &t.GroupID, &t.StudentID)

		if err != nil {
			return err
		}

		g.GroupStudents = append(g.GroupStudents, t)
	}

	return nil
}

func AddGroupStudent(w http.ResponseWriter, r *http.Request) {
	var request Request
	var newGroupStudent GroupStudent

	defer LogHandler("group student add")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &newGroupStudent)
	HandleError(err, w, WrongDataError)

	err = newGroupStudent.Add()
	HandleError(err, w, UnknownError)

	SendData(w, 200, newGroupStudent)
}

func UpdateGroupStudent(w http.ResponseWriter, r *http.Request) {
	var request Request
	var updatingGroupStudent GroupStudent

	defer LogHandler("group student update")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	updatingGroupStudent.Init()
	err = json.Unmarshal(textJson, &updatingGroupStudent)
	HandleError(err, w, WrongDataError)

	err = updatingGroupStudent.Update()
	HandleError(err, w, UnknownError)

	SendData(w, 200, updatingGroupStudent)
}

func RemoveGroupStudent(w http.ResponseWriter, r *http.Request) {
	var request Request
	var removingGroupStudent GroupStudent

	defer LogHandler("group student remove")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &removingGroupStudent)
	HandleError(err, w, WrongDataError)

	err = removingGroupStudent.Remove()
	HandleError(err, w, UnknownError)

	SendData(w, 200, removingGroupStudent)
}

func SelectGroupStudents(w http.ResponseWriter, r *http.Request) {
	var request Request
	var searchingGroupStudent GroupStudent
	var selectedGroupStudents GroupStudents

	defer LogHandler("group students select")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	searchingGroupStudent.Init()
	err = json.Unmarshal(textJson, &searchingGroupStudent)
	HandleError(err, w, WrongDataError)

	err = selectedGroupStudents.Select(searchingGroupStudent)
	HandleError(err, w, UnknownError)

	SendData(w, 200, selectedGroupStudents)
}

type Group struct {
	ID        int  `json:"id"`
	Active    int   `json:"active"`
	Name      string `json:"name"`
	Time      string `json:"time"`
	Duration  int  `json:"duration"`
	Weekday   int  `json:"weekday"`
	GroupType int  `json:"group_type"`
}

func (g *Group) Init() {
	g.ID = -1
	g.Active = -1
	g.Name = ""
	g.Time = ""
	g.Duration = -1
	g.Weekday = -1
	g.GroupType = -1
}

func (g Group) Add() error {
	var query string
	var queryValues []interface{}

	queryValues = append(queryValues, g.Name)
	queryValues = append(queryValues, g.Time)
	queryValues = append(queryValues, g.Duration)
	queryValues = append(queryValues, g.Weekday)
	queryValues = append(queryValues, g.GroupType)

	query = "insert into robotenok.`groups` (name, time, duration, weekday, group_type) values (?,?,?,?,?)"

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(queryValues...)

	return err
}

func (g Group) Update() error {
	if g.ID == -1 {
		return errors.New("group id has wrong data")
	}

	var query string
	var isFirst bool
	var queryValues []interface{}

	// Wrote it separately because goland marked it as error -_(O_O|)_-
	query = "update robotenok.`groups`" + " set "
	isFirst = true

	if g.Active != -1 {
		query += "active = ?"
		queryValues = append(queryValues, g.Active)
		isFirst = false
	}

	if g.Name != "" {
		if isFirst == false {
			query += ","
		}

		query += " name = ?"
		queryValues = append(queryValues, g.Name)
		isFirst = false
	}

	if g.Time != "" {
		if isFirst == false {
			query += ","
		}

		query += " time = ?"
		queryValues = append(queryValues, g.Time)
		isFirst = false
	}

	if g.Duration != -1 {
		if isFirst == false {
			query += ","
		}

		query += " duration = ?"
		queryValues = append(queryValues, g.Duration)
		isFirst = false
	}

	if g.Weekday != -1 {
		if isFirst == false {
			query += ","
		}

		query += " weekday = ?"
		queryValues = append(queryValues, g.Duration)
		isFirst = false
	}

	if g.GroupType != -1 {
		if isFirst == false {
			query += ","
		}

		query += " group_type = ?"
		queryValues = append(queryValues, g.GroupType)
	}

	query += " where id = " + "?"
	queryValues = append(queryValues, g.ID)

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(queryValues...)
	return err
}

func (g *Group) Remove() error {
	g.Active = 0
	return g.Update()
}

type Groups struct {
	Groups []Group `json:"groups"`
}

func (g *Groups) Select(q Group) error {
	var query string
	var isSearch bool
	var queryValues []interface{}

	isSearch = false
	query = "select * from robotenok.`groups`" + " where "

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

	if q.Time != "" {
		query += " and time = ?"
		queryValues = append(queryValues, q.Time)
		isSearch = true
	}

	if q.Duration != -1 {
		query += " and duration = ?"
		queryValues = append(queryValues, q.Duration)
		isSearch = true
	}

	if q.Weekday != -1 {
		query += " and weekday = ?"
		queryValues = append(queryValues, q.Weekday)
		isSearch = true
	}

	if q.GroupType != -1 {
		query += " and group_type = ?"
		queryValues = append(queryValues, q.GroupType)
		isSearch = true
	}

	if q.Name != "" {
		query += " and name like '%" + template.HTMLEscapeString(q.Name) + "%'"
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
		t := Group{}
		err := row.Scan(&t.ID, &t.Active, &t.Name, &t.Time, &t.Duration, &t.Weekday, &t.GroupType)

		if err != nil {
			return err
		}

		g.Groups = append(g.Groups, t)
	}

	return nil
}

func AddGroup(w http.ResponseWriter, r *http.Request) {
	var request Request
	var newGroup Group

	defer LogHandler("group add")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &newGroup)
	HandleError(err, w, WrongDataError)

	err = newGroup.Add()
	HandleError(err, w, UnknownError)

	SendData(w, 200, newGroup)
}

func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	var request Request
	var updatingGroup Group

	defer LogHandler("group update")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &updatingGroup)
	HandleError(err, w, WrongDataError)

	err = updatingGroup.Update()
	HandleError(err, w, UnknownError)

	SendData(w, 200, updatingGroup)
}

func RemoveGroup(w http.ResponseWriter, r *http.Request) {
	var request Request
	var removingGroup Group

	defer LogHandler("group remove")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &removingGroup)
	HandleError(err, w, WrongDataError)

	err = removingGroup.Remove()
	HandleError(err, w, UnknownError)

	SendData(w, 200, removingGroup)
}

func SelectGroups(w http.ResponseWriter, r *http.Request) {
	var request Request
	var searchingGroup Group
	var selectedGroups Groups

	defer LogHandler("groups select")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	searchingGroup.Init()
	err = json.Unmarshal(textJson, &searchingGroup)
	HandleError(err, w, WrongDataError)

	err = selectedGroups.Select(searchingGroup)
	HandleError(err, w, UnknownError)

	SendData(w, 200, selectedGroups)
}