package database

import "time"

type Row struct {
	UID   string
	Title string
	Date  time.Time
}
