package models

import "time"

type Bank struct {
	ID      int       `json:"id"`
	ACCOUNT string    `json:"account"`
	PID     string    `json:"pid"`
	CLIENT  int       `json:"client_type"`
	OPER    int       `json:"operation_type"`
	CRE     time.Time `json:"created_time"`
	IP      string    `json:"ip"`
}
