package main

import (
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)


func GenString(length int) string {
	var output string

	for i := 0; i < length; i++ {
		rand.Seed(time.Now().UnixNano())

		randChoice := rand.Intn(2)

		if randChoice == 0 {
			output = output + strconv.Itoa(rand.Intn(10))
		} else {
			output = output + string('a'+rune(rand.Intn(26)))
		}

		time.Sleep(time.Microsecond * 1)
	}

	return output
}

func ToSHA512(input string) string {
	data := sha512.Sum512([]byte(input))
	return hex.EncodeToString(data[:])
}

func Auth(w http.ResponseWriter, r *http.Request) {
	var input AuthData
	var user User

	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&input)

	if err != nil {
		log.Println("[" + GetIP(r) + "] " + "[ERROR] [Auth] " + err.Error())
		SendData(w, 400, "Client send a wrong or empty data")
		return
	}

	user.Login = input.Login
	err = user.Select()

	if SearchUser(user).ID == user.ID {
		var err Error
		err.create("User already in use")

		SendData(w, 400, err)
		return
	}

	if err != nil {
		SendData(w, 400, err.Error())
		return
	}

	if user.Login == "" {
		var err Error
		err.create("User is not existing")

		SendData(w, 400, err)
		return
	}

	if input.Hash != user.Password {
		var err Error
		err.create("Password wrong")

		SendData(w, 400, err)
		return
	}

	user.Secret = ToSHA512(GenString(64))
	user.Online = time.Now()

	log.Println("[INFO] [" + GetIP(r) + "] " + user.Name + " signed up")
	ActiveUsers = append(ActiveUsers, user)

	SendData(w, 200, user.Secret)
}


func (r *Request) checkToken () error {
	var user User

	user.Secret = r.Token
	user = SearchUser(user)

	if user.ID == -1 {
		return errors.New("security error")
	}

	if r.Token != user.Secret {
		return errors.New("security error")
	}

	r.UserID = user.ID
	user.Online = time.Now()

	return nil
}