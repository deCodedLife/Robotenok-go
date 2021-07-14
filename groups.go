package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
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

	query = "insert into robotenok.group_types (name) values ?"
	_, err := db.Exec(query, g.Name)

	return err
}

func (g GroupType) Update() error {
	if g.ID == -1 {
		return errors.New("group type id has wrong data")
	}

	var query string
	var isFirst bool

	// Wrote it separately because goland marked it as error -_(O_O|)_-
	query = "update robotenok.students" + " set "
	isFirst = true

	if g.Name != "" {
		query += "name = " + g.Name
		isFirst = false
	}

	if g.Active != -1 {
		if isFirst == false {
			query += ","
		}

		query += "active = " + strconv.Itoa(g.Active)
	}

	query += " where id = " + strconv.Itoa(g.ID)

	_, err := db.Exec(query)
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

	if q.Name != "" {
		query += " and name like = %" + q.Name + "%"
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

	query = "insert into robotenok.`groups` (name, time, duration, weekday, group_type) values (?,?,?,?,?)"
	_, err := db.Exec(query, g.Name, g.Time, g.Duration, g.Weekday, g.GroupType)

	return err
}

func (g Group) Update() error {
	if g.ID == -1 {
		return errors.New("group id has wrong data")
	}

	var query string
	var isFirst bool

	// Wrote it separately because goland marked it as error -_(O_O|)_-
	query = "update robotenok.`groups`" + " set "
	isFirst = true

	if g.Active != -1 {
		query += "active = " + strconv.Itoa(g.Active)
		isFirst = false
	}

	if g.Name != "" {
		if isFirst == false {
			query += ","
		}

		query += " name = " + g.Name
		isFirst = false
	}

	if g.Time != "" {
		if isFirst == false {
			query += ","
		}

		query += " time = " + g.Time
		isFirst = false
	}

	if g.Duration != -1 {
		if isFirst == false {
			query += ","
		}

		query += " duration = " + strconv.Itoa(g.Duration)
		isFirst = false
	}

	if g.Weekday != -1 {
		if isFirst == false {
			query += ","
		}

		query += " weekday = " + strconv.Itoa(g.Weekday)
		isFirst = false
	}

	if g.GroupType != -1 {
		if isFirst == false {
			query += ","
		}

		query += " group_type = " + strconv.Itoa(g.GroupType)
	}

	query += " where id = " + strconv.Itoa(g.ID)

	_, err := db.Exec(query)
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

	isSearch = false
	query = "select * from robotenok.`groups`" + " where "

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

	if q.Time != "" {
		query += " and time = " + q.Time
		isSearch = true
	}

	if q.Duration != -1 {
		query += " and duration = " + strconv.Itoa(q.Duration)
		isSearch = true
	}

	if q.Weekday != -1 {
		query += " and weekday = " + strconv.Itoa(q.Weekday)
		isSearch = true
	}

	if q.GroupType != -1 {
		query += " and group_type = " + strconv.Itoa(q.GroupType)
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