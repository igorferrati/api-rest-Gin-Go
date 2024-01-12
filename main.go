package main

import (
	"github.com/igorferrati/api-rest-Gin-Go/database"
	"github.com/igorferrati/api-rest-Gin-Go/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HndleRequests()
}
