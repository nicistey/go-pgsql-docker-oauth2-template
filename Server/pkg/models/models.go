package models

import "time"

type Event struct {
	IDev        int 
	IDus        int 
	Event_name  string 
	Event_time  time.Time
	Description string
	Location    string
	Is_public   bool
}

type User struct {
	IDus     int
	GoogleID string 
	Name     string
	Email    string
}