package main

import (
	"errors"
	"strconv"
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
	var query string

	query = "insert into robotenok.courses (name, payment, lessons, image) values (?, ?, ?, ?)"
	_, err := db.Exec(query, c.Name, c.Payment, c.Lessons, c.Image)

	return err
}

func (c Course) Update() error {
	if c.ID == -1 {
		return errors.New("course id has wrong data")
	}

	var query string
	var isFirst bool

	// Wrote it separately because goland marked it as error -_(O_O|)_-
	query = "update robotenok.students" + " set "
	isFirst = true

	if c.Name != "" {
		query += " name like %" + c.Name + "%"
		isFirst = false
	}

	if c.Payment != -1 {
		if isFirst == false {
			query += ","
		}

		query += " payment = " + strconv.Itoa(c.Payment)
		isFirst = false
	}

	if c.Lessons != -1 {
		if isFirst == false {
			query += ","
		}

		query += " lessons = " + strconv.Itoa(c.Lessons)
		isFirst = false
	}

	if c.Image != "" {
		if isFirst == false {
			query += ","
		}

		query += " image = " + c.Image
		isFirst = false
	}

	query += " where id = " + strconv.Itoa(c.ID)

	_, err := db.Exec(query)
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
		query += " and name like %" + q.Name + "%"
		isSearch = true
	}

	if q.Payment != -1 {
		query += " and payment = " + strconv.Itoa(q.Payment)
		isSearch = true
	}

	if q.Lessons != -1 {
		query += " and lessons = " + strconv.Itoa(q.Lessons)
		isSearch = true
	}

	if q.Image != "" {
		query += " and image = " + q.Image
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
		t := Course{}
		err := row.Scan(&t.ID, &t.Active, &t.Name, &t.Payment, &t.Lessons, &t.Image)

		if err != nil {
			return err
		}

		c.Courses = append(c.Courses, t)
	}

	return nil
}

