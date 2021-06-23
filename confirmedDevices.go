package main

import (
	"errors"
	"reflect"
)

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