package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var db *sql.DB
var ActiveUsers []User

var WrongDataError ResponceError
var SecurityError ResponceError
var UnknownError ResponceError


func configure() {
	var err error
	var database Database

	WrongDataError.WrongDataError()
	SecurityError.SecurityError()
	UnknownError.UnknownError()

	// Local database values
	database.Username = "admin"
	database.Password = "Zero_twO*Rengyou"
	database.Database = "robotenok"

	log.SetFlags(log.LstdFlags | log.Llongfile)
	createDirectory("Logs")

	date := time.Now().Format(time.RFC1123)
	dates := strings.Split(date, ":")
	date = ""
	for i := 0; i < len(dates); i++ {
		date = date + dates[i] + "-"
	}
	f, err := os.OpenFile("logs/"+date+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 777)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)

	db, err = sql.Open("mysql", database.Username+":"+database.Password+"@tcp(localhost:3306)/"+database.Database)

	if err != nil {
		fmt.Println("[ERROR] [Database] " + err.Error())
		return
	}

	log.Println("[INFO] [Database] connect successful")
}

func DataException(w http.ResponseWriter) {
	var response Response
	response.Status = 400

	response.Response = map[string]interface{}{
		"error": "Client send a wrong or empty data",
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Println("[ERROR] Error during encode data exception.", err.Error())
	}
}

func access(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")

		var input AuthData
		var user User

		err := json.NewDecoder(r.Body).Decode(&input)

		if err != nil {
			DataException(w)
		}

		user.Login = input.Login
		err = user.Select()

		if err != nil {
			log.Println(err.Error())
		}

		if user.Password == input.Hash && user.UserType == "root" {
			err := json.NewEncoder(w).Encode(ActiveUsers)

			if err != nil {
				log.Println("[ERROR] Error during encode data.", err.Error())
			}
		}

	} else {
		log.Println(GetIP(r) + "[INFO] Visits a root of server")
		_, err := fmt.Fprintf(w, "We tracking you! Dont trust us? Your ip will be logged: "+GetIP(r))

		if err != nil {
			log.Println("Can't return message to user. ", err.Error())
		}
	}
}

func initHandlers(r *mux.Router) {
	r.HandleFunc("/", access).Methods("GET", "POST")
	r.HandleFunc("/robotenok/auth", Auth).Methods("POST")

	r.HandleFunc("/robotenok/user/add/{hash}", AddUser).Methods("POST")
	r.HandleFunc("/robotenok/user/update", UpdateUser).Methods("POST")
	r.HandleFunc("/robotenok/user/remove", RemoveUser).Methods("POST")
	r.HandleFunc("/robotenok/user/select", SelectUser).Methods("POST")

	r.HandleFunc("/robotenok/students/add", AddStudent).Methods("POST")
	r.HandleFunc("/robotenok/students/update", UpdateStudent).Methods("POST")
	r.HandleFunc("/robotenok/students/remove", RemoveStudent).Methods("POST")
	r.HandleFunc("/robotenok/students/select", SelectStudents).Methods("POST")

	r.HandleFunc("/robotenok/visits/add", AddVisit).Methods("POST")
	r.HandleFunc("/robotenok/visits/update", UpdateVisit).Methods("POST")
	r.HandleFunc("/robotenok/visits/remove", RemoveVisit).Methods("POST")
	r.HandleFunc("/robotenok/visits/select", SelectVisits).Methods("POST")

	r.HandleFunc("/robotenok/group-types/add", AddGroupType).Methods("POST")
	r.HandleFunc("/robotenok/group-types/update", UpdateGroupType).Methods("POST")
	r.HandleFunc("/robotenok/group-types/remove", RemoveGroupType).Methods("POST")
	r.HandleFunc("/robotenok/group-types/select", SelectGroupTypes).Methods("POST")

	r.HandleFunc("/robotenok/groups/add", AddGroup).Methods("POST")
	r.HandleFunc("/robotenok/groups/update", UpdateGroup).Methods("POST")
	r.HandleFunc("/robotenok/groups/remove", RemoveGroup).Methods("POST")
	r.HandleFunc("/robotenok/groups/select", SelectGroups).Methods("POST")

	r.HandleFunc("/robotenok/payments/add", AddPayment).Methods("POST")
	r.HandleFunc("/robotenok/payments/update", UpdatePayment).Methods("POST")
	r.HandleFunc("/robotenok/payments/remove", RemovePayment).Methods("POST")
	r.HandleFunc("/robotenok/payments/select", SelectPayments).Methods("POST")

	r.HandleFunc("/robotenok/group-students/add", AddGroupStudent).Methods("POST")
	r.HandleFunc("/robotenok/group-students/update", UpdateGroupStudent).Methods("POST")
	r.HandleFunc("/robotenok/group-students/remove", RemoveGroupStudent).Methods("POST")
	r.HandleFunc("/robotenok/group-students/select", SelectGroupStudents).Methods("POST")
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	configure()

	r := mux.NewRouter()

	/* TODO:
	* Realize all handlers
	 */
	initHandlers(r)

	go func() {
		UserDaemon()
	}()


	err := http.ListenAndServe(":80", r)

	if err != nil {
		log.Println("Can't start server. Reason: ", err.Error())
	}
}
