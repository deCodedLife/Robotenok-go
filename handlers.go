package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

func GetDate() string {
	currentTime := time.Now()
	return currentTime.Format("2006-01-02")
}

func GetTime() string {
	currentTime := time.Now()
	dateStamp := currentTime.Format("2006.01.02 15:04:05")
	dateArray := strings.Fields(dateStamp)
	return dateArray[1]
}

func (e *Error) create(description interface{}) {
	e.Error = description
}

func LogHandler (sender string) {
	if r := recover(); r != nil {
		LogData(sender, r)
	}
}

func HandleError(err error, w http.ResponseWriter, r ResponceError) {
	if err != nil {
		SendData(w, r.Status, r.Description)
		panic(r.Description + " cause: " + err.Error())
	}
}

func requestHandler (request *Request, r *http.Request) error {
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		return err
	}

	return err
}

func typeFromString(t string) int {
	for result, currentType := range userTypes {
		if t == currentType {
			return result
		}
	}

	return -1
}

func createDirectory(dirName string) bool {
	src, err := os.Stat(dirName)

	if os.IsNotExist(err) {

		errDir := os.MkdirAll(dirName, 0755)
		if errDir != nil {
			panic(err)
		}

		return true
	}

	if src.Mode().IsRegular() {
		fmt.Println(dirName, "already exist as a file!")
		return false
	}

	return false
}

func permCheck(userID int, perm int) error {
	var user User

	user.Init()
	user.ID	= userID
	err := user.Select()

	if err != nil {
		return err
	}

	userType := typeFromString(user.UserType)

	if userType > perm || userType == -1 {
		return errors.New("you have no permissions")
	}

	return nil
}

func UsersTimeout() {
	for {
		var currentTime = time.Now()

		if len(ActiveUsers) == 0 {
			return
		}

		for index, user := range ActiveUsers {
			var onlineTime = user.Online.Local().Add(time.Minute * 15)

			if currentTime.Sub(onlineTime) > 1 {
				log.Printf("[INFO] Timeout. %s was disconnected\n", user.Name)
				sliceUser(index)
			}
		}

		time.Sleep(time.Minute * 1)
	}
}

func sliceUser (index int) {
	if len(ActiveUsers) == 1 {
		ActiveUsers = ActiveUsers[:len(ActiveUsers)-1]
	}

	ActiveUsers[index] = ActiveUsers[len(ActiveUsers)-1]
	ActiveUsers = ActiveUsers[:len(ActiveUsers)-1]
}

func SearchUser(u User) User {
	if len(ActiveUsers) == 0 {
		return User{ID: -1}
	}

	for _, user := range ActiveUsers {
		if u.Secret == user.Secret {
			return user
		}
	}

	return User{ID: -1}
}

func GetIP(req *http.Request) string {
	ip, _, err := net.SplitHostPort(req.RemoteAddr)

	if err != nil {
		return req.RemoteAddr
	}

	userIP := net.ParseIP(ip)

	if userIP == nil {
		return req.RemoteAddr
	}

	forward := req.Header.Get("X-Forwarded-For")

	if forward != "" {
		forward = "Forwarded[" + forward + "]"
	}

	return forward + "IP " + req.RemoteAddr
}

func LogData (sender string, data interface{}) {
	log.Println("[" + strings.ToUpper(sender) + "]", data)
}

func SendData(w http.ResponseWriter, status int, data interface{}) {
	var response Response

	response.Status = status
	response.Response = data

	if status != 200 {
		w.WriteHeader(http.StatusInternalServerError)
		err := Error{data}
		data = err
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		LogData("response writer", "Can't send data to user. Reason: " + err.Error())
	}
}