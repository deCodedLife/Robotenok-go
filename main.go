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

	defer LogHandler("configs")

	db, err = sql.Open("mysql", database.Username+":"+database.Password+"@tcp(localhost:3306)/"+database.Database)

	if err != nil {
		panic("[ERROR] [Database] " + err.Error())
		return
	}

	log.Println("[INFO] [Database] connect successful")
}

func access(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")

		var input AuthData
		var user User

		defer LogHandler("access")

		err := json.NewDecoder(r.Body).Decode(&input)
		HandleError(err, w, WrongDataError)

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

	r.HandleFunc("/robotenok/users/{hash}", AddUser).Methods("POST")
	r.HandleFunc("/robotenok/users", UpdateUser).Methods("PUT")
	r.HandleFunc("/robotenok/users", RemoveUser).Methods("DELETE")
	r.HandleFunc("/robotenok/select-users", SelectUser).Methods("POST")

	r.HandleFunc("/robotenok/students", AddStudent).Methods("POST")
	r.HandleFunc("/robotenok/students", UpdateStudent).Methods("PUT")
	r.HandleFunc("/robotenok/students", RemoveStudent).Methods("DELETE")
	r.HandleFunc("/robotenok/select-students", SelectStudents).Methods("POST")

	r.HandleFunc("/robotenok/visits", AddVisit).Methods("POST")
	r.HandleFunc("/robotenok/visits", UpdateVisit).Methods("PUT")
	r.HandleFunc("/robotenok/visits", RemoveVisit).Methods("DELETE")
	r.HandleFunc("/robotenok/select-visits", SelectVisits).Methods("POST")

	r.HandleFunc("/robotenok/group-types", AddGroupType).Methods("POST")
	r.HandleFunc("/robotenok/group-types", UpdateGroupType).Methods("PUT")
	r.HandleFunc("/robotenok/group-types", RemoveGroupType).Methods("DELETE")
	r.HandleFunc("/robotenok/select-group-types", SelectGroupTypes).Methods("POST")

	r.HandleFunc("/robotenok/groups", AddGroup).Methods("POST")
	r.HandleFunc("/robotenok/groups", UpdateGroup).Methods("PUT")
	r.HandleFunc("/robotenok/groups", RemoveGroup).Methods("DELETE")
	r.HandleFunc("/robotenok/select-groups", SelectGroups).Methods("POST")

	r.HandleFunc("/robotenok/group-curators", AddGroupCurator).Methods("POST")
	r.HandleFunc("/robotenok/group-curators", UpdateGroupCurator).Methods("PUT")
	r.HandleFunc("/robotenok/group-curators", RemoveGroupCurator).Methods("DELETE")
	r.HandleFunc("/robotenok/select-group-curators", SelectGroupCurators).Methods("POST")

	r.HandleFunc("/robotenok/payments", AddPayment).Methods("POST")
	r.HandleFunc("/robotenok/payments", UpdatePayment).Methods("PUT")
	r.HandleFunc("/robotenok/payments", RemovePayment).Methods("DELETE")
	r.HandleFunc("/robotenok/select-payments", SelectPayments).Methods("POST")

	r.HandleFunc("/robotenok/group-students", AddGroupStudent).Methods("POST")
	r.HandleFunc("/robotenok/group-students", UpdateGroupStudent).Methods("PUT")
	r.HandleFunc("/robotenok/group-students", RemoveGroupStudent).Methods("DELETE")
	r.HandleFunc("/robotenok/select-group-students", SelectGroupStudents).Methods("POST")

	r.HandleFunc("/robotenok/courses", AddCourse).Methods("POST")
	r.HandleFunc("/robotenok/courses", UpdateCourse).Methods("PUT")
	r.HandleFunc("/robotenok/courses", RemoveCourse).Methods("DELETE")
	r.HandleFunc("/robotenok/select-courses", SelectCourses).Methods("POST")

	r.HandleFunc("/robotenok/costs", AddCost).Methods("POST")
	r.HandleFunc("/robotenok/costs", UpdateCost).Methods("PUT")
	r.HandleFunc("/robotenok/costs", RemoveCost).Methods("DELETE")
	r.HandleFunc("/robotenok/select-costs", SelectCosts).Methods("POST")

	r.HandleFunc("/robotenok/classes", AddClass).Methods("POST")
	r.HandleFunc("/robotenok/classes", UpdateClass).Methods("PUT")
	r.HandleFunc("/robotenok/classes", RemoveClass).Methods("DELETE")
	r.HandleFunc("/robotenok/select-classes", SelectClasses).Methods("POST")
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	r := mux.NewRouter()

	configure()
	initHandlers(r)

	go func() {
		UsersTimeout()
	}()

	err := http.ListenAndServe(":80", r)

	if err != nil {
		log.Println("Can't start server. Reason: ", err.Error())
	}
}
