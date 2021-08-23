package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
)

type Device struct {
	ID     int    `json:"id"`
	Hash   string `json:"hash"`
	Active int    `json:"active"`
}

func (d *Device) Init () {
	d.ID = -1
	d.Hash = ""
	d.Active = -1
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


func (c *ConfirmedDevices) Get(info interface{}) error {
	var whereState string

	if reflect.TypeOf(info).Kind() == reflect.String {
		whereState = "hash = ?"
		info = info.(string)
	} else if reflect.TypeOf(info).Kind() == reflect.Int32 {
		whereState = "id = ?"
		info = info.(int32)
	} else {
		return errors.New("wrong data type")
	}

	row := db.QueryRow("select * from confirmed_devices where "+whereState, info)

	if row.Err() != nil {
		return row.Err()
	}

	err := row.Scan(&c.ID, &c.Hash, &c.Active)
	return err
}

func (c ConfirmedDevices) Add() error {
	_, err := db.Query("insert into robotenok.confirmed_devices (hash) values (?)", c.Hash)
	return err
}

func (c ConfirmedDevices) StatusChange() error {
	_, err := db.Query("update robotenok.confirmed_devices set active = ? where id = ?", c.Active, c.ID)
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