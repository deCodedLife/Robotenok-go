package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"strconv"
)

type CostReceipt struct {
	ID int `json:"id"`
	PaymentID int `json:"payment_id"`
	ImageID int `json:"image_id"`
	Active int `json:"active"`
}

func (p *CostReceipt) Init() {
	p.ID = -1
	p.PaymentID = -1
	p.ImageID = -1
	p.Active = -1
}

func (p CostReceipt) Add () error {
	var queryValues []interface{}

	queryValues = append(queryValues, p.PaymentID)
	queryValues = append(queryValues, p.ImageID)

	var query = "insert into robotenok.costs_receipts (payment_id, image_id) values (?, ?)"

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(queryValues)

	return err
}

func (p CostReceipt) Update() error {
	if p.ID == -1 {
		return errors.New("payment receipt has wrong data")
	}

	var queryValues []interface{}

	// Wrote it separately because goland marked it as error -_(O_O|)_-
	var query = "update robotenok.costs_receipts" + " set "
	var isFirst = true

	if p.Active != -1 {
		query += " active = ?"
		queryValues = append(queryValues, p.Active)
		isFirst = true
	}

	if p.PaymentID != -1 {
		if isFirst == false {
			query += ","
		}

		query += " payment_id = ?"
		queryValues = append(queryValues, p.PaymentID)
		isFirst = false
	}

	if p.ImageID != -1 {
		if isFirst == false {
			query += ","
		}

		query += " image_id = ?"
		queryValues = append(queryValues, p.ImageID)
	}

	query += " where id = " + ""
	queryValues = append(queryValues, p.ID)

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	_, err = stmt.Exec(queryValues)
	return err
}

func (p *CostReceipt) Remove() error {
	p.Active = 0
	return p.Update()
}

type CostsReceipts struct {
	PaymentReceipts []CostReceipt `json:"payment_receipts"`
}

func (p *CostsReceipts) Select(q CostReceipt) error {
	var queryValues []interface{}

	var isSearch = false
	var query = "select * from robotenok.costs_receipts" + " where "

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

	if q.PaymentID != -1 {
		query += " and payment_id = ?"
		queryValues = append(queryValues, q.PaymentID)
		isSearch = true
	}

	if q.ImageID != -1 {
		query += "and image_id = ?"
		queryValues = append(queryValues, q.ImageID)
		isSearch = true
	}

	if isSearch == false {
		return errors.New("nothing to do")
	}

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	row, err := stmt.Query(queryValues)

	if err != nil {
		return err
	}

	for row.Next() {
		t := CostReceipt{}
		err := row.Scan(&t.ID, &t.PaymentID, &t.ImageID, &t.Active)

		if err != nil {
			return err
		}

		p.PaymentReceipts = append(p.PaymentReceipts, t)
	}

	return nil
}

func AddCostsReceipt(w http.ResponseWriter, r *http.Request) {
	var request Request
	var newPaymentReceipt CostReceipt

	defer LogHandler("payment receipts add")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 0)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &newPaymentReceipt)
	HandleError(err, w, WrongDataError)

	err = newPaymentReceipt.Add()
	HandleError(err, w, UnknownError)

	SendData(w, 200, newPaymentReceipt)
}

func UpdateCostsReceipt(w http.ResponseWriter, r *http.Request) {
	var request Request
	var updatingPaymentReceipt CostReceipt

	defer LogHandler("payment receipts update")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 0)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	updatingPaymentReceipt.Init()
	err = json.Unmarshal(textJson, &updatingPaymentReceipt)
	HandleError(err, w, WrongDataError)

	err = updatingPaymentReceipt.Update()
	HandleError(err, w, UnknownError)

	SendData(w, 200, updatingPaymentReceipt)
}

func RemoveCostsReceipt(w http.ResponseWriter, r *http.Request) {
	var request Request
	var removingPaymentReceipt CostReceipt

	defer LogHandler("payment receipts remove")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 0)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &removingPaymentReceipt)
	HandleError(err, w, WrongDataError)

	err = removingPaymentReceipt.Remove()
	HandleError(err, w, UnknownError)

	SendData(w, 200, removingPaymentReceipt)
}

func SelectCostsReceipts(w http.ResponseWriter, r *http.Request) {
	var request Request
	var searchingPaymentReceipt CostReceipt
	var selectedPaymentReceipts CostsReceipts

	defer LogHandler("payment receipts select")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	searchingPaymentReceipt.Init()
	err = json.Unmarshal(textJson, &searchingPaymentReceipt)
	HandleError(err, w, WrongDataError)

	err = selectedPaymentReceipts.Select(searchingPaymentReceipt)
	HandleError(err, w, UnknownError)

	SendData(w, 200, selectedPaymentReceipts.PaymentReceipts)
}

type Cost struct {
	ID      int    `json:"id"`
	Active  int    `json:"active"`
	Product string `json:"product"`
	Cost    int    `json:"cost"`
	Date    string `json:"date"`
	Time    string `json:"time"`
	PaymentObject int `json:"payment_object"`
}

func (c *Cost) Init() {
	c.ID = -1
	c.Active = -1
	c.Product = ""
	c.Cost = -1
	c.Date = ""
	c.Time = ""
	c.PaymentObject = -1
}

func (c Cost) Add() error {
	var queryValues []interface{}

	queryValues = append(queryValues, c.Product)
	queryValues = append(queryValues, c.Cost)
	queryValues = append(queryValues, GetDate())
	queryValues = append(queryValues, GetTime())
	queryValues = append(queryValues, c.PaymentObject)

	var query = "insert into robotenok.costs (product, cost, date, time, payment_object) values (?, ?, ?, ?, ?)"

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

	var queryValues []interface{}

	// Wrote it separately because goland marks it as error -_(O_O|)_-
	var query = "update robotenok.costs" + " set "
	var isFirst = true

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

	if c.PaymentObject != -1 {
		if isFirst == false {
			query += ","
		}

		query += " payment_object = ?"
		queryValues = append(queryValues, c.PaymentObject)
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
	var queryValues []interface{}

	var isSearch = false
	var query = "select * from robotenok.costs" + " where "

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

	if q.PaymentObject != -1 {
		query += " and payment_object = ?"
		queryValues = append(queryValues, q.PaymentObject)
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
		err := row.Scan(&t.ID, &t.Active, &t.Cost, &t.Product, &t.Date, &t.Time, &t.PaymentObject)

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

	SendData(w, 200, selectedCosts.Costs)
}