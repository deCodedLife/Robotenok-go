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
var ImagesFolder string
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

	ImagesFolder = "images"

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
	r.PathPrefix("/robotenok/images/").Handler(
		http.StripPrefix("/robotenok/images/",
		http.FileServer(http.Dir(ImagesFolder)))).Methods("GET")

	r.HandleFunc("/robotenok/addDevice", AddDevice).Methods("POST")

	r.HandleFunc("/robotenok/image", AddImage).Methods("POST")
	r.HandleFunc("/robotenok/image", UpdateImage).Methods("PUT")
	r.HandleFunc("/robotenok/image", RemoveImage).Methods("DELETE")
	r.HandleFunc("/robotenok/images", SelectImages).Methods("POST")

	r.HandleFunc("/robotenok/uploads/{hash}", UploadImage).Methods("POST")
	r.HandleFunc("/robotenok/uploads/{hash}", RemoveImageFile).Methods("DELETE")

	r.HandleFunc("/", access).Methods("GET", "POST")
	r.HandleFunc("/robotenok/auth", Auth).Methods("POST")

	r.HandleFunc("/robotenok/user/{hash}", AddUser).Methods("POST")
	r.HandleFunc("/robotenok/user", UpdateUser).Methods("PUT")
	r.HandleFunc("/robotenok/user", RemoveUser).Methods("DELETE")
	r.HandleFunc("/robotenok/users", SelectUser).Methods("POST")

	r.HandleFunc("/robotenok/student", AddStudent).Methods("POST")
	r.HandleFunc("/robotenok/student", UpdateStudent).Methods("PUT")
	r.HandleFunc("/robotenok/student", RemoveStudent).Methods("DELETE")
	r.HandleFunc("/robotenok/students", SelectStudents).Methods("POST")

	r.HandleFunc("/robotenok/visit", AddVisit).Methods("POST")
	r.HandleFunc("/robotenok/visit", UpdateVisit).Methods("PUT")
	r.HandleFunc("/robotenok/visit", RemoveVisit).Methods("DELETE")
	r.HandleFunc("/robotenok/visits", SelectVisits).Methods("POST")

	r.HandleFunc("/robotenok/group-type", AddGroupType).Methods("POST")
	r.HandleFunc("/robotenok/group-type", UpdateGroupType).Methods("PUT")
	r.HandleFunc("/robotenok/group-type", RemoveGroupType).Methods("DELETE")
	r.HandleFunc("/robotenok/group-types", SelectGroupTypes).Methods("POST")

	r.HandleFunc("/robotenok/group", AddGroup).Methods("POST")
	r.HandleFunc("/robotenok/group", UpdateGroup).Methods("PUT")
	r.HandleFunc("/robotenok/group", RemoveGroup).Methods("DELETE")
	r.HandleFunc("/robotenok/groups", SelectGroups).Methods("POST")

	r.HandleFunc("/robotenok/group-curator", AddGroupCurator).Methods("POST")
	r.HandleFunc("/robotenok/group-curator", UpdateGroupCurator).Methods("PUT")
	r.HandleFunc("/robotenok/group-curator", RemoveGroupCurator).Methods("DELETE")
	r.HandleFunc("/robotenok/group-curators", SelectGroupCurators).Methods("POST")

	r.HandleFunc("/robotenok/payment", AddPayment).Methods("POST")
	r.HandleFunc("/robotenok/payment", UpdatePayment).Methods("PUT")
	r.HandleFunc("/robotenok/payment", RemovePayment).Methods("DELETE")
	r.HandleFunc("/robotenok/payments", SelectPayments).Methods("POST")

	r.HandleFunc("/robotenok/payment-object", AddPaymentObject).Methods("POST")
	r.HandleFunc("/robotenok/payment-object", UpdatePaymentObject).Methods("PUT")
	r.HandleFunc("/robotenok/payment-object", RemovePaymentObject).Methods("DELETE")
	r.HandleFunc("/robotenok/payment-objects", SelectPaymentsObject).Methods("POST")

	r.HandleFunc("/robotenok/group-student", AddGroupStudent).Methods("POST")
	r.HandleFunc("/robotenok/group-student", UpdateGroupStudent).Methods("PUT")
	r.HandleFunc("/robotenok/group-student", RemoveGroupStudent).Methods("DELETE")
	r.HandleFunc("/robotenok/group-students", SelectGroupStudents).Methods("POST")

	r.HandleFunc("/robotenok/course", AddCourse).Methods("POST")
	r.HandleFunc("/robotenok/course", UpdateCourse).Methods("PUT")
	r.HandleFunc("/robotenok/course", RemoveCourse).Methods("DELETE")
	r.HandleFunc("/robotenok/courses", SelectCourses).Methods("POST")

	r.HandleFunc("/robotenok/cost", AddCost).Methods("POST")
	r.HandleFunc("/robotenok/cost", UpdateCost).Methods("PUT")
	r.HandleFunc("/robotenok/cost", RemoveCost).Methods("DELETE")
	r.HandleFunc("/robotenok/costs", SelectCosts).Methods("POST")

	r.HandleFunc("/robotenok/cost-receipt", AddCostsReceipt).Methods("POST")
	r.HandleFunc("/robotenok/cost-receipt", UpdateCostsReceipt).Methods("PUT")
	r.HandleFunc("/robotenok/cost-receipt", RemoveCostsReceipt).Methods("DELETE")
	r.HandleFunc("/robotenok/cost-receipts", SelectCostsReceipts).Methods("POST")

	r.HandleFunc("/robotenok/class", AddClass).Methods("POST")
	r.HandleFunc("/robotenok/class", UpdateClass).Methods("PUT")
	r.HandleFunc("/robotenok/class", RemoveClass).Methods("DELETE")
	r.HandleFunc("/robotenok/classes", SelectClasses).Methods("POST")
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
