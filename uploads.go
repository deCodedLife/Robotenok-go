package main

import (
	"errors"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"os"
)

func (i *ImageData) GetImage() error {
	var selectedImages ImagesData
	var err = selectedImages.Select(*i)

	if len(selectedImages.Images) == 0 && err != nil {
		err = errors.New("hash not exist")
	}

	i.ID = selectedImages.Images[0].ID
	i.UserID = selectedImages.Images[0].UserID
	i.Date = selectedImages.Images[0].Date
	i.Time = selectedImages.Images[0].Time
	i.FileName = selectedImages.Images[0].FileName
	i.Active = selectedImages.Images[0].Active
	i.Hash = selectedImages.Images[0].Hash

	return err
}

func UploadImage(w http.ResponseWriter, r *http.Request) {
	var param = mux.Vars(r)
	var hash = param["hash"]

	var searchingImage ImageData
	searchingImage.Init()
	searchingImage.Hash = hash

	//r.ParseMultipartForm(0)

	defer LogHandler("image upload")
	var err = searchingImage.GetImage()

	HandleError(err, w, WrongDataError)

	image, err := ioutil.ReadAll(r.Body)
	HandleError(err, w, WrongDataError)

	err = ioutil.WriteFile(ImagesFolder + searchingImage.FileName, image, 0777)
	HandleError(err, w, UnknownError)

	SendData(w, 200, searchingImage)
}

/*
var param = mux.Vars(r)
	var hash = param["hash"]

	var searchingImage ImageData
	searchingImage.Init()
	searchingImage.Hash = hash

	r.ParseMultipartForm(0)

	defer LogHandler("image upload")
	var err = searchingImage.GetImage()

	HandleError(err, w, WrongDataError)

	for _, h := range r.MultipartForm.Value[0] {
		image, err := h.Open()
		HandleError(err, w, WrongDataError)

		defer image.Close()

		file, err := os.Create(ImagesFolder + "/" + hash)
		defer file.Close()

		HandleError(err, w, UnknownError)

		_, err = io.Copy(file, image)
		HandleError(err, w, UnknownError)
	}

	SendData(w, 200, searchingImage)
 */

func RemoveImageFile(w http.ResponseWriter, r *http.Request) {
	var param = mux.Vars(r)
	var hash = param["hash"]

	var searchingImage ImageData
	searchingImage.Init()
	searchingImage.Hash = hash

	defer LogHandler("image upload")
	var err = searchingImage.GetImage()

	HandleError(err, w, WrongDataError)

	err = os.Remove(ImagesFolder + searchingImage.FileName)
	HandleError(err, w, UnknownError)

	err = searchingImage.Remove()
	HandleError(err, w, UnknownError)

	SendData(w, 200, searchingImage)
}