package models

import (
	"go-todo/dao"
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	ID        uint           `json:"id" gorm:"primarykey"`
	Title     string         `json:"title"`
	Status    bool           `json:"status"`
}

func AddTodo(todo *Todo) (err error) {
	return dao.DB.Create(&todo).Error
}

func ModifyTodo(todo *Todo) (err error) {
	return dao.DB.Save(&todo).Error
}

func DeleteTodo(ID string) (err error) {
	return dao.DB.Delete(&Todo{}, ID).Error
}
func GetAllTodos() (todos []*Todo, err error) {
	err = dao.DB.Find(&todos).Error
	return todos, err
}
func GetTodo(ID string) (todo *Todo, err error) {
	todo = new(Todo)
	err = dao.DB.First(&todo, ID).Error
	return todo, err
}
