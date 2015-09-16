package entity

import (
	"time"
)

type ImagePath struct {
	ID         int       `json:"id"`
	Alias      string    `json:"alias"`
	ImagePath  string    `json:"imagepath sql:"size:255"`
	CreateDate time.Time `json:"createdate"`
}
