package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/igorferrati/api-rest-Gin-Go/controllers"
)

func HndleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.Run()
}
