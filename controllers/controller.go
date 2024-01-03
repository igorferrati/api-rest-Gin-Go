package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/igorferrati/api-rest-Gin-Go/models"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"Api diz:": "Eai " + nome + ", tudo bem?",
	})
}
