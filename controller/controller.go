package controller

import (
	"github.com/gin-gonic/gin"
	"go-todo/models"
)

type Todo = models.Todo

func IndexHandle(ctx *gin.Context) {
	ctx.HTML(200, "index.html", nil)
}

func AddTodo(ctx *gin.Context) {
	var todo Todo
	ctx.BindJSON(&todo)
	if err := models.AddTodo(&todo); err != nil {
		ctx.JSON(200, gin.H{
			"code":    1,
			"message": err,
		})
	} else {
		ctx.JSON(200, gin.H{
			"code":    0,
			"message": "success",
			"data":    &todo,
		})
	}
}

func UpdateTodo(ctx *gin.Context) {
	ID, isExist := ctx.Params.Get("id")
	if !isExist {
		ctx.JSON(200, gin.H{
			"code":    1,
			"message": "id字段未提供~",
		})
		return
	}
	todo, err := models.GetTodo(ID)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    1,
			"message": err,
		})
		return
	}
	ctx.BindJSON(&todo)
	err = models.ModifyTodo(todo)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    1,
			"message": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"data": todo,
	})
}
func DeleteTodo(ctx *gin.Context) {
	ID, isExist := ctx.Params.Get("id")
	if !isExist {
		ctx.JSON(200, gin.H{
			"code":    1,
			"message": "id字段未提供~",
		})
		return
	}
	err := models.DeleteTodo(ID)
	if err != nil {
		ctx.JSON(200, gin.H{
			"code":    1,
			"message": err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 0,
	})
}

func GetAllTodos(ctx *gin.Context) {
	if todos, err := models.GetAllTodos(); err != nil {
		ctx.JSON(200, gin.H{
			"code":    1,
			"message": err,
		})
	} else {
		ctx.JSON(200, todos)
	}
}
