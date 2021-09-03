package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Device struct {
	ID     int    `json:"id"`
	Hash   string `json:"hash"`
	Active int    `json:"active"`
	Status int    `json:"status"`
}

func (d *Device) Init () {
	d.ID = -1
	d.Hash = ""
	d.Active = -1
	d.Status = -1
}

func (d *Device) Add() error {
	var queryValues []interface{}

	d.Hash = GenString(32)
	queryValues = append(queryValues, d.Hash)

	var query = "insert into robotenok.confirmed_devices (hash) values (?)"

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(queryValues...)

	return err
}

func (d Device) Update() error {

	if d.ID == -1 {
		return errors.New("device id has wrong data")
	}

	var queryValues []interface{}

	// Wrote it separately because goland marked it as error -_(O_O|)_-
	var query = "update robotenok.confirmed_devices" + " set "
	var isFirst = true

	if d.Hash != "" {
		query += " hash = ?"
		queryValues = append(queryValues, d.Hash)
		isFirst = false
	}

	if d.Active != -1 {
		if isFirst == false {
			query += ","
		}

		query += " active = ?"
		queryValues = append(queryValues, d.Active)
	}

	query += " where id = " + ""
	queryValues = append(queryValues, d.ID)

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	_, err = stmt.Exec(queryValues...)
	return err
}

func (d *Device) Remove() error {
	d.Active = 0
	return d.Update()
}

func (c *ConfirmedDevices) Init() {
	c.ID = -1
	c.Hash = ""
	c.Active = -1
}


func (d *Device) Get() error {
	var query = "select * from robotenok.confirmed_devices where active = 1"

	var queryValues []interface{}

	if d.ID != -1 {
		query += " and id = ?"
		queryValues = append(queryValues, d.ID)
	}

	if d.Hash != "" {
		query += " and hash = ?"
		queryValues = append(queryValues, d.Hash)
	}

	stmt, err := db.Prepare(query)

	if err != nil {
		return err
	}

	row := stmt.QueryRow(queryValues...)

	if row.Err() != nil {
		return row.Err()
	}

	err = row.Scan(&d.ID, &d.Hash, &d.Active, &d.Status)
	return err
}

func (c ConfirmedDevices) Add() error {
	_, err := db.Query("insert into robotenok.confirmed_devices (hash) values (?)", c.Hash)
	return err
}

func (d Device) StatusChange() error {
	_, err := db.Query("update robotenok.confirmed_devices set active = ? where id = ?", d.Active, d.ID)
	return err
}

func AddDevice(w http.ResponseWriter, r *http.Request) {
	var request Request
	var addingDevice Device

	defer LogHandler("device add")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &addingDevice)
	HandleError(err, w, WrongDataError)

	err = addingDevice.Add()
	HandleError(err, w, UnknownError)

	SendData(w, 200, addingDevice)
}