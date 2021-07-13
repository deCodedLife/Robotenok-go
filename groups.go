package main

import (
	"errors"
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