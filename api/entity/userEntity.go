// user entity
package entity

import (
	"time"
)

type User struct {
	ID         int       `json:"id" sql:"AUTO_INCREMENT"`
	Name       string    `json:"name" sql:"size:255"`
	Alias      string    `josn:"alias" sql:"size:255"`
	Password   string    `json:"password" sql:"size:100"`
	CreateDate time.Time `json:"createdate"`
	Desc       string    `json:"desc"`
}
