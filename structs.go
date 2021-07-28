package main

type ResponceError struct {
	Status      int
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
	Status   int       `json:"status"`
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

type ConfirmedDevices struct {
	ID     int32  `json:"id"`
	Hash   string `json:"hash"`
	Active int8   `json:"active"`
}
