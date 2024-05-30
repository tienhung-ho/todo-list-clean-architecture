package common

import "time"

type SqlModel struct {
	Id         int        `json:"id" gorm:"column:id;"`
	Created_at *time.Time `json:"created_at" gorm:"column:created_at;"`
	Updated_at *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}
