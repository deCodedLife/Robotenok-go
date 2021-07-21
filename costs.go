package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"strconv"
)

type Cost struct {
	ID      int    `json:"id"`
	Active  int    `json:"active"`
	Product string `json:"product"`
	Cost    int    `json:"cost"`
	Date    string `json:"date"`
	Time    string `json:"time"`
}

func (c *Cost) Init() {
	c.ID = -1
	c.Active = -1
	c.Product = ""
	c.Cost = -1
	c.Date = ""
	c.Time = ""
}

func (c Cost) Add() error {
	var query string
	var queryValues []interface{}

	queryValues = append(queryValues, c.Product)
	queryValues = append(queryValues, c.Cost)
	queryValues = append(queryValues, GetDate())
	queryValues = append(queryValues, GetTime())

	query = "insert into robotenok.costs (product, cost, date, time) values (?, ?, ?, ?)"

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(queryValues...)

	return err
}

func (c Cost) Update() error {
	if c.ID == -1 {
		return errors.New("costs id has wrong data")
	}

	var query string
	var isFirst bool
	var queryValues []interface{}

	// Wrote it separately because goland marks it as error -_(O_O|)_-
	query = "update robotenok.costs" + " set "
	isFirst = true

	if c.Product != "" {
		query += " product like '%" + template.HTMLEscapeString(c.Product) + "%'"
		isFirst = false
	}

	if c.Active != -1 {
		if isFirst == false {
			query += ","
		}

		query += " active = ?"
		queryValues = append(queryValues, c.Active)
		isFirst = false
	}

	if c.Cost != -1 {
		if isFirst == false {
			query += ","
		}

		query += " cost = ?"
		queryValues = append(queryValues, c.Cost)
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

	query += " where id = " + strconv.Itoa(c.ID)

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(queryValues...)

	return err
}

func (c *Cost) Remove() error {
	c.Active = 0
	return c.Update()
}
type Costs struct {
	Costs []Cost `json:"costs"`
}

func (c *Costs) Select(q Cost) error {
	var query string
	var isSearch bool
	var queryValues []interface{}

	isSearch = false
	query = "select * from robotenok.costs" + " where "

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

	if q.Product != "" {
		query += " and product like '%" + template.HTMLEscapeString(q.Product) + "%'"
		isSearch = true
	}

	if q.Cost != -1 {
		query += " and cost = ?"
		queryValues = append(queryValues, q.Cost)
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
		t := Cost{}
		err := row.Scan(&t.ID, &t.Active, &t.Cost, &t.Product, &t.Date, &t.Time)

		if err != nil {
			return err
		}

		c.Costs = append(c.Costs, t)
	}

	return nil
}

func AddCost(w http.ResponseWriter, r *http.Request) {
	var request Request
	var newCost Cost

	defer LogHandler("cost add")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &newCost)
	HandleError(err, w, WrongDataError)

	err = newCost.Add()
	HandleError(err, w, UnknownError)

	SendData(w, 200, newCost)
}

func UpdateCost(w http.ResponseWriter, r *http.Request) {
	var request Request
	var updatingCost Cost

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	updatingCost.Init()
	err = json.Unmarshal(textJson, &updatingCost)
	HandleError(err, w, WrongDataError)

	err = updatingCost.Update()
	HandleError(err, w, UnknownError)

	SendData(w, 200, updatingCost)
}

func RemoveCost(w http.ResponseWriter, r *http.Request) {
	var request Request
	var removingCost Cost

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &removingCost)
	HandleError(err, w, WrongDataError)

	err = removingCost.Remove()
	HandleError(err, w, WrongDataError)

	SendData(w, 200, removingCost)
}

func SelectCosts(w http.ResponseWriter, r *http.Request) {
	var request Request
	var searchingCost Cost
	var selectedCosts Costs

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	searchingCost.Init()
	err = json.Unmarshal(textJson, &searchingCost)
	HandleError(err, w, WrongDataError)

	err = selectedCosts.Select(searchingCost)
	HandleError(err, w, WrongDataError)

	SendData(w, 200, selectedCosts)
}