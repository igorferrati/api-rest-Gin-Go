package main

import (
	"github.com/igorferrati/api-rest-Gin-Go/database"
	"github.com/igorferrati/api-rest-Gin-Go/models"
	"github.com/igorferrati/api-rest-Gin-Go/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Igor Ferrati", CPF: "12345678910", RG: "123456789"},
		{Nome: "Carla Cartoman", CPF: "12345678910", RG: "123456789"},
	}
	routes.HndleRequests()
}
