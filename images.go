package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type ImageData struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Date   string `json:"date"`
	Time   string `json:"time"`
	FileName   string `json:"filename"`
	Active int    `json:"active"`
	Hash   string `json:"hash"`
}

func (i *ImageData) Init() {
	i.ID = -1
	i.UserID = -1
	i.Date = ""
	i.Time = ""
	i.FileName = ""
	i.Active = -1
	i.Hash = ""
}

func (i ImageData) Add() error {
	var queryValues []interface{}

	queryValues = append(queryValues, i.UserID)
	queryValues = append(queryValues, GetDate())
	queryValues = append(queryValues, GetTime())
	queryValues = append(queryValues, i.FileName)
	queryValues = append(queryValues, i.Hash)

	var query = "insert into robotenok.images (user_id, date, time, filename, hash) values (?, ?, ?, ?, ?)"

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	if err != nil {
		return err
	}

	_, err = stmt.Exec(queryValues...)

	return err
}

func (i ImageData) Update() error {
	if i.ID == -1 {
		return errors.New("payment id has wrong data")
	}

	var queryValues []interface{}

	// Wrote it separately because goland marked it as error -_(O_O|)_-
	var query = "update robotenok.images" + " set "
	var isFirst = true

	if i.UserID != -1 {
		query += "user_id = ?"
		queryValues = append(queryValues, i.UserID)
		isFirst = false
	}

	if i.Date != "" {
		if isFirst == false {
			query += ","
		}

		query += " date = ?"
		queryValues = append(queryValues, i.Date)
		isFirst = false
	}

	if i.Time != "" {
		if isFirst == false {
			query += ","
		}

		query += " time = ?"
		queryValues = append(queryValues, i.Time)
		isFirst = false
	}

	if i.FileName != "" {
		if isFirst == false {
			query += ","
		}

		query += " path = ?"
		queryValues = append(queryValues, i.FileName)
		isFirst = false
	}

	if i.Active != -1 {
		if isFirst == false {
			query += ","
		}

		query += " active = ?"
		queryValues = append(queryValues, i.Active)
		isFirst = false
	}

	if i.Hash != "" {
		if isFirst == false {
			query += ","
		}

		query += " hash = ?"
		queryValues = append(queryValues, i.Hash)
		isFirst = false
	}

	query += " where id = " + ""
	queryValues = append(queryValues, i.ID)

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	_, err = stmt.Exec(queryValues...)
	return err
}

func (i *ImageData) Remove() error {
	i.Active = -1
	return i.Update()
}

type ImagesData struct {
	Images []ImageData `json:"images"`
}

func (i *ImagesData) Select (q ImageData) error {
	var queryValues []interface{}

	var isSearch = false
	var query = "select * from robotenok.images" + " where "

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

	if q.UserID != -1 {
		query += " and user_id = ?"
		queryValues = append(queryValues, q.UserID)
		isSearch = true
	}

	if q.Hash != "" {
		query += " and hash = ?"
		queryValues = append(queryValues, q.Hash)
		isSearch = true
	}

	if q.FileName != "" {
		query += " and filename = ?"
		queryValues = append(queryValues, q.FileName)
		isSearch = true
	}

	if isSearch == false {
		return errors.New("nothing to do")
	}

	stmt, err := db.Prepare(query)
	defer stmt.Close()

	row, err := stmt.Query(queryValues...)

	if err != nil {
		return err
	}

	for row.Next() {
		t := ImageData{}
		err := row.Scan(&t.ID, &t.UserID, &t.Date, &t.Time, &t.FileName, &t.Active, &t.Hash)

		if err != nil {
			return err
		}

		i.Images = append(i.Images, t)
	}

	return nil
}

func (i *ImageData) GenRandom() {
	var output ImageData
	var selectedImage ImagesData

	output.Init()
	output.FileName = GenString(32)
	output.Hash = GenString(42)

	selectedImage.Select(output)

	if len(selectedImage.Images) > 0 {
		output.GenRandom()
	}

	i.FileName = output.FileName
	i.Hash = output.Hash
}

func AddImage(w http.ResponseWriter, r *http.Request) {
	var request Request
	var newImage ImageData

	defer LogHandler("image add")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &newImage)
	HandleError(err, w, WrongDataError)

	newImage.GenRandom()

	err = newImage.Add()
	HandleError(err, w, UnknownError)

	SendData(w, 200, newImage)
}

func UpdateImage (w http.ResponseWriter, r *http.Request) {
	var request Request
	var updatingImage ImageData

	defer LogHandler("image update")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	updatingImage.Init()
	err = json.Unmarshal(textJson, &updatingImage)
	HandleError(err, w, WrongDataError)

	err = updatingImage.Update()
	HandleError(err, w, UnknownError)

	SendData(w, 200, updatingImage)
}

func RemoveImage(w http.ResponseWriter, r *http.Request) {
	var request Request
	var removingImage ImageData

	defer LogHandler("image remove")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	err = permCheck(request.UserID, 1)
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	err = json.Unmarshal(textJson, &removingImage)
	HandleError(err, w, WrongDataError)

	err = removingImage.Remove()
	HandleError(err, w, UnknownError)

	SendData(w, 200, removingImage)
}

func SelectImages(w http.ResponseWriter, r *http.Request) {
	var request Request
	var searchingImage ImageData
	var selectedImages ImagesData

	defer LogHandler("image select")

	err := requestHandler(&request, r)
	HandleError(err, w, WrongDataError)

	err = request.checkToken()
	HandleError(err, w, SecurityError)

	textJson, err := json.Marshal(request.Body)
	HandleError(err, w, WrongDataError)

	searchingImage.Init()
	err = json.Unmarshal(textJson, &searchingImage)
	HandleError(err, w, WrongDataError)

	err = selectedImages.Select(searchingImage)
	HandleError(err, w, UnknownError)

	SendData(w, 200, selectedImages.Images)
}