package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
)

type Course struct {
	ID      int    `json:"id"`
	Active  int    `json:"active"`
	Name    string `json:"name"`
	Payment int    `json:"payment"`
	Lessons int    `json:"lessons"`
	Image   string `json:"image"`
}

func (c Course) Init() {
	c.ID = -1
	c.Active = -1
	c.Name = ""
	c.Payment = -1
	c.Lessons = -1
	c.Image = ""
}

func (c Course) Add() error {
	var queryValues []interface{}

	queryValues = append(queryValues, c.Name)
	queryValues = append(queryValues, c.Payment)
	queryValues = append(queryValues, c.Lessons)
	queryValues = append(queryValues, c.Image)

	var query = "insert into robotenok.courses (name, payment, lessons, image) values (?, ?, ?, ?)"

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(queryValues...)

	return err
}

func (c Course) Update() error {
	if c.ID == -1 {
		return errors.New("course id has wrong data")
	}

	var queryValues []interface{}

	// Wrote it separately because goland marked it as error -_(O_O|)_-
	var query = "update robotenok.courses" + " set "
	var isFirst = true

	if c.Name != "" {
		query += " name like %" + template.HTMLEscapeString(c.Name) + "%"
		isFirst = false
	}

	if c.Payment != -1 {
		if isFirst == false {
			query += ","
		}

		query += " payment = ?"
		queryValues = append(queryValues, c.Payment)
		isFirst = false
	}

	if c.Lessons != -1 {
		if isFirst == false {
			query += ","
		}

		query += " lessons = ?"
		queryValues = append(queryValues, c.Lessons)
		isFirst = false
	}

	if c.Image != "" {
		if isFirst == false {
			query += ","
		}

		query += " image = ?"
		queryValues = append(queryValues, c.Image)
		isFirst = false
	}

	// JetBrains marks: "where id = ?" as error
	query += " where id = " + "?"
	queryValues = append(queryValues, c.ID)

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(queryValues...)
	return err
}

func (c *Course) Remove() error {
	c.Active = 0
	return c.Update()
}

type Courses struct {
	Courses []Course `json:"courses"`
}

func (c *Courses) Select(q Course) error {
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

	if q.Name != "" {
		query += " and name like '%" + template.HTMLEscapeString(q.Name) + "%'"
		isSearch = true
	}

	if q.Payment != -1 {
		query += " and payment = ?"
		queryValues = append(queryValues, q.Payment)
		isSearch = true
	}

	if q.Lessons != -1 {
		query += " and lessons = ?"
		queryValues = append(queryValues, q.Lessons)
		isSearch = true
	}

	if q.Image != "" {
		query += " and image = ?"
		queryValues = append(queryValues, q.Image)
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
		t := Course{}
		err := row.Scan(&t.ID, &t.Active, &t.Name, &t.Payment, &t.Lessons, &t.Image)

		if err != nil {
			return err
		}

		c.Courses = append(c.Courses, t)
	}

	return nil
}

func AddCourse(w http.ResponseWriter, r *http.Request) {
	var request Request
	var newCourse Course

	defer LogHandler("course add")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &newCourse)
	HandleError(err, w, WrongDataError)

	err = newCourse.Add()
	HandleError(err, w, UnknownError)

	SendData(w, 200, newCourse)
}

func UpdateCourse(w http.ResponseWriter, r * http.Request) {
	var request Request
	var updatingCourse Course

	defer LogHandler("course update")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	updatingCourse.Init()
	err = json.Unmarshal(textJson, &updatingCourse)
	HandleError(err, w, WrongDataError)

	err = updatingCourse.Update()
	HandleError(err, w, UnknownError)

	SendData(w, 200, updatingCourse)
}

func RemoveCourse(w http.ResponseWriter, r *http.Request) {
	var request Request
	var removingCourse Course

	defer LogHandler("course remove")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &removingCourse)
	HandleError(err, w, WrongDataError)

	err = removingCourse.Remove()
	HandleError(err, w, UnknownError)

	SendData(w, 200, removingCourse)
}

func SelectCourses(w http.ResponseWriter, r *http.Request) {
	var request Request
	var searchingCourse Course
	var selectedCourses Courses

	defer LogHandler("courses select")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	searchingCourse.Init()
	err = json.Unmarshal(textJson, &searchingCourse)
	HandleError(err, w, WrongDataError)

	err = selectedCourses.Select(searchingCourse)
	HandleError(err, w, UnknownError)

	SendData(w, 200, selectedCourses.Courses)
}

type CourseGroup struct {
	ID       int `json:"id"`
	CourseID int `json:"course_id"`
	GroupID  int `json:"group_id"`
	Active   int `json:"active"`
}

func (c *CourseGroup) Init() {
	c.ID = -1
	c.CourseID = -1
	c.GroupID = -1
	c.Active = -1
}

func (c CourseGroup) Add() error {
	var queryValues []interface{}

	queryValues = append(queryValues, c.GroupID)
	queryValues = append(queryValues, c.GroupID)

	var query = "insert into robotenok.course_groups (course_id, group_id) values (?, ?)"

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	_, err = stmt.Exec(queryValues...)

	return err
}

func (c CourseGroup) Update() error {
	if c.ID == -1 {
		return errors.New("course group id has wrong data")
	}

	var queryValues []interface{}

	// Wrote it separately because goland marked it as error -_(O_O|)_-
	var query = "update robotenok.course_groups" + " set "
	var isFirst = true

	if c.CourseID != -1 {
		query += " course_id = ?"
		queryValues = append(queryValues, c.CourseID)
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

	if c.Active != -1 {
		if isFirst == false {
			query += ","
		}

		query += " active = ?"
		queryValues = append(queryValues, c.Active)
	}

	// JetBrains marks: "where id = ?" as error
	query += " where id = " + "?"
	queryValues = append(queryValues, c.ID)

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(queryValues...)

	return err
}

func (c *CourseGroup) Remove() error {
	c.Active = -1
	return c.Update()
}

type CourseGroups struct {
	CourseGroups []CourseGroup `json:"course_groups"`
}

func (c *CourseGroups) Select(q CourseGroup) error {
	var queryValues []interface{}

	var isSearch = false
	var query = "select * from robotenok.course_groups" + " where "

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

	if q.CourseID != -1 {
		query += " and course_id = ?"
		queryValues = append(queryValues, q.CourseID)
		isSearch = true
	}

	if q.GroupID != -1 {
		query += " group_id = ?"
		queryValues = append(queryValues, q.GroupID)
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

	row, err := stmt.Query(queryValues)

	if err != nil {
		return err
	}

	for row.Next() {
		t := CourseGroup{}
		err := row.Scan(&t.ID, &t.CourseID, &t.GroupID, &t.Active)

		if err != nil {
			return err
		}

		c.CourseGroups = append(c.CourseGroups, t)
	}

	return nil
}

func AddCourseGroup(w http.ResponseWriter, r *http.Request) {
	var request Request
	var newCourseGroup CourseGroup

	defer LogHandler("course group add")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &newCourseGroup)
	HandleError(err, w, WrongDataError)

	err = newCourseGroup.Add()
	HandleError(err, w, UnknownError)

	SendData(w, 200, newCourseGroup)
}

func UpdateCourseGroup(w http.ResponseWriter, r *http.Request) {
	var request Request
	var updatingCourseGroup CourseGroup

	defer LogHandler("course group add")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	updatingCourseGroup.Init()
	err = json.Unmarshal(textJson, &updatingCourseGroup)
	HandleError(err, w, WrongDataError)

	err = updatingCourseGroup.Update()
	HandleError(err, w, UnknownError)

	SendData(w, 200, updatingCourseGroup)
}

func RemoveCourseGroup(w http.ResponseWriter, r *http.Request) {
	var request Request
	var removingCourseGroup CourseGroup

	defer LogHandler("course group remove")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &removingCourseGroup)
	HandleError(err, w, WrongDataError)

	err = removingCourseGroup.Remove()
	HandleError(err, w, UnknownError)

	SendData(w, 200, removingCourseGroup)
}

func SelectCourseGroups(w http.ResponseWriter, r *http.Request) {
	var request Request
	var searchingCourseGroup CourseGroup
	var selectedCourseGroups CourseGroups

	defer LogHandler("course groups select")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	searchingCourseGroup.Init()
	err = json.Unmarshal(textJson, &searchingCourseGroup)
	HandleError(err, w, WrongDataError)

	err = selectedCourseGroups.Select(searchingCourseGroup)
	HandleError(err, w, UnknownError)

	SendData(w, 200, selectedCourseGroups)
}