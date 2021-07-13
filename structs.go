package main

type ResponceError struct {
	Status int32
	Description string
}

func (r *ResponceError) WrongDataError() {
	r.Status = 400
	r.Description = "Client send a wrong or empty data"
}

func (r *ResponceError) SecurityError() {
	r.Status = 401
	r.Description = "Security error"
}

func (r *ResponceError) UnknownError() {
	r.Status = 500
	r.Description = "Something went wrong"
}

type Database struct {
	Username string
	Password string
	Database string
}

type Error struct {
	Error interface{} `json:"error"`
}

type Response struct {
	Status   int32       `json:"status"`
	Response interface{} `json:"response"`
}

type Request struct {
	UserID int         `json:"user_id"`
	Token  string      `json:"token"`
	Body   interface{} `json:"body"`
}

type AuthData struct {
	Login string `json:"login"`
	Hash  string `json:"hash"`
}

type Payment struct {
	ID        int32  `json:"id"`
	Active    string `json:"active"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	StudentID string `json:"student_id"`
	Credit    int32  `json:"credit"`
	Type      string `json:"type"`
	UserID    int32  `json:"user_id"`
}

type GroupType struct {
	ID     int32  `json:"id"`
	Active int8   `json:"active"`
	Name   string `json:"name"`
}

type Group struct {
	ID        int32  `json:"id"`
	Active    int8   `json:"active"`
	Name      string `json:"name"`
	Time      string `json:"time"`
	Duration  int32  `json:"duration"`
	Weekday   int32  `json:"weekday"`
	GroupType int32  `json:"group_type"`
}

type GroupStudents struct {
	ID        int32 `json:"id"`
	Active    int8  `json:"active"`
	GroupID   int32 `json:"group_id"`
	StudentID int32 `json:"student_id"`
}

type Course struct {
	ID      int32  `json:"id"`
	Active  int8   `json:"active"`
	Name    string `json:"name"`
	Payment int32  `json:"payment"`
	Lessons int8   `json:"lessons"`
	Image   string `json:"image"`
}

type CourseGroups struct {
	ID       int32 `json:"id"`
	CourseID int32 `json:"course_id"`
	GroupID  int32 `json:"group_id"`
}

type Cost struct {
	ID      int32  `json:"id"`
	Active  int8   `json:"active"`
	Product string `json:"product"`
	Cost    int8   `json:"cost"`
	Date    string `json:"date"`
	Time    string `json:"time"`
}

type ConfirmedDevices struct {
	ID     int32  `json:"id"`
	Hash   string `json:"hash"`
	Active int8   `json:"active"`
}
