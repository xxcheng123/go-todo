package routers

import (
	"github.com/gin-gonic/gin"
	"go-todo/controller"
)

func SetupRputer() *gin.Engine {
	r := gin.Default()
	//加载模板
	r.LoadHTMLGlob("templates/*")
	//加载静态资源
	r.Static("/static", "static")
	r.GET("/", controller.IndexHandle)

	v1Group := r.Group("/v1")
	{
		//提交新的
		v1Group.POST("/todo", controller.AddTodo)
		//获取全部
		v1Group.GET("/todo", controller.GetAllTodos)
		//更新
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
