package model

import (
	"errors"
	"todo-list/common"
)

var (
	ErrTitleBlank  = errors.New("cannot leave title blank")
	ErrItemDeleted = errors.New("item is deleted!")
)

type TodoItem struct {
	Title       string      `json:"title" gorm:"column:title;"`
	Description string      `json:"description" gorm:"column:description;"`
	Status      *ItemStatus `json:"status" gorm:"column:status;"`
	common.SqlModel
}

func (TodoItem) TableName() string { return "todo_items" }
