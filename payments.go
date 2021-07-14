package main

type Payment struct {
	ID        int    `json:"id"`
	Active    string `json:"active"`
	Date      string `json:"date"`
	Time      string `json:"time"`
	StudentID string `json:"student_id"`
	Credit    int    `json:"credit"`
	Type      string `json:"type"`
	UserID    int    `json:"user_id"`
}

