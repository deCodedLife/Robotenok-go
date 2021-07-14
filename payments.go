package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

type Payment struct {
	ID        int    `json:"id"`
	Active    int    `json:"active"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	StudentID int    `json:"student_id"`
	Credit    int    `json:"credit"`
	Type      string `json:"type"`
	UserID    int    `json:"user_id"`
}

func (p *Payment) Init() {
	p.ID = -1
	p.Active = -1
	p.Date = ""
	p.Time = ""
	p.StudentID = -1
	p.Credit = -1
	p.Type = ""
	p.UserID = -1
}

func (p Payment) Add() error {
	var query string

	query = "insert into robotenok.payments (date, time, student_id, credit, type, user_id) values (?,?,?,?,?,?)"
	_, err := db.Exec(query, GetDate(), GetTime(), p.StudentID, p.Credit, p.Type, p.UserID)

	return err
}

func (p Payment) Update() error {
	if p.ID == -1 {
		return errors.New("payment id has wrong data")
	}

	var query string
	var isFirst bool

	// Wrote it separately because goland marked it as error -_(O_O|)_-
	query = "update robotenok.students" + " set "
	isFirst = true

	if p.Date != "" {
		query += "date = " + p.Date
		isFirst = false
	}

	if p.Time != "" {
		if isFirst == false {
			query += ","
		}

		query += " time = " + p.Time
		isFirst = false
	}

	if p.StudentID != -1 {
		if isFirst == false {
			query += ","
		}

		query += " student_id = " + strconv.Itoa(p.StudentID)
		isFirst = false
	}

	if p.Credit != -1 {
		if isFirst == false {
			query += ","
		}

		query += " credit = " + strconv.Itoa(p.Credit)
		isFirst = false
	}

	if p.Type != "" {
		if isFirst == false {
			query += ","
		}

		query += " type = " + p.Type
		isFirst = false
	}

	if p.UserID != -1 {
		if isFirst == false {
			query += ","
		}

		query += " user_id = " + strconv.Itoa(p.UserID)
	}

	query += " where id = " + strconv.Itoa(p.ID)

	_, err := db.Exec(query)
	return err
}

func (p *Payment) Remove() error {
	p.Active = 0
	return p.Update()
}

type Payments struct {
	Payments []Payment `json:"payments"`
}

func (p *Payments) Select(q Payment) error {
	var query string
	var isSearch bool

	isSearch = false
	query = "select * from robotenok.payments" + " where "

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

	if q.Date != "" {
		query += " and date = " + q.Date
		isSearch = true
	}

	if q.Time != "" {
		query += " and time = " + q.Time
		isSearch = true
	}

	if q.StudentID != -1 {
		query += " and student_id = " + strconv.Itoa(q.StudentID)
		isSearch = true
	}

	if q.Credit != -1 {
		query += " and credit = " + strconv.Itoa(q.Credit)
		isSearch = true
	}

	if q.Type != "" {
		query += " and type = " + q.Type
		isSearch = true
	}

	if q.UserID != -1 {
		query += " and user_id = " + strconv.Itoa(q.UserID)
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
		t := Payment{}
		err := row.Scan(&t.ID, &t.Active, &t.Date, &t.Time, &t.StudentID, &t.Credit, &t.Type, &t.UserID)

		if err != nil {
			return err
		}

		p.Payments = append(p.Payments, t)
	}

	return nil
}

func AddPayment(w http.ResponseWriter, r *http.Request) {
	var request Request
	var newPayment Payment

	defer LogHandler("payment add")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &newPayment)
	HandleError(err, w, WrongDataError)

	err = newPayment.Add()
	HandleError(err, w, UnknownError)

	SendData(w, 200, newPayment)
}

func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	var request Request
	var updatingPayment Payment

	defer LogHandler("payment update")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	updatingPayment.Init()
	err = json.Unmarshal(textJson, &updatingPayment)
	HandleError(err, w, WrongDataError)

	err = updatingPayment.Update()
	HandleError(err, w, UnknownError)

	SendData(w, 200, updatingPayment)
}

func RemovePayment(w http.ResponseWriter, r *http.Request) {
	var request Request
	var removingPayment Payment

	defer LogHandler("payment remove")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &removingPayment)
	HandleError(err, w, WrongDataError)

	err = removingPayment.Remove()
	HandleError(err, w, UnknownError)

	SendData(w, 200, removingPayment)
}

func SelectPayments(w http.ResponseWriter, r *http.Request) {
	var request Request
	var searchingPayment Payment
	var selectedPayments Payments

	defer LogHandler("payments select")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	searchingPayment.Init()
	err = json.Unmarshal(textJson, &searchingPayment)
	HandleError(err, w, WrongDataError)

	err = selectedPayments.Select(searchingPayment)
	HandleError(err, w, UnknownError)

	SendData(w, 200, selectedPayments)
}