package entities

import (
	"time"
)

type Entry struct {
	ID    int
	Title string
	Text  string
	Date  time.Time
}
