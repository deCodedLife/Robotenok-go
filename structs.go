package main

type ResponceError struct {
	Status      int32
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

type Cost struct {
	ID      int    `json:"id"`
	Active  int    `json:"active"`
	Product string `json:"product"`
	Cost    int    `json:"cost"`
	Date    string `json:"date"`
	Time    string `json:"time"`
}

type ActiveClasses struct {
	ID          int    `json:"id"`
	Active      int    `json:"active"`
	UserID      int    `json:"user_id"`
	GroupID     int    `json:"group_id"`
	Date        string `json:"date"`
	Time        string `json:"time"`
	HostAddress string `json:"host_address"`
}

type ConfirmedDevices struct {
	ID     int32  `json:"id"`
	Hash   string `json:"hash"`
	Active int8   `json:"active"`
}
