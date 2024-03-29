package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/igorferrati/api-rest-Gin-Go/database"
	"github.com/igorferrati/api-rest-Gin-Go/models"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// ExibeTodosAlunos exibe todos alunos
//
//	@Summary      Exibir todos alunos
//	@Description  Exibir todos alunos
//	@Tags         get alunos
//	@Accept       json
//	@Produce      json
//	@Success      200  {object}   models.Aluno
//	@Failure      400  {object}  httputil.HTTPError
//	@Router       /alunos [get]
func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"Api diz:": "Eai " + nome + ", tudo bem?",
	})
}

// CriaNovoAluno godoc
//
//	@Summary      Cria novo aluno
//	@Description  Rota para cirar um novo aluno
//	@Tags         criar aluno
//	@Accept       json
//	@Produce      json
//	@Param        aluno    body     models.Aluno true "Aluno"
//	@Success      200  {object}   models.Aluno
//	@Failure      400  {object}  httputil.HTTPError
//	@Router       /alunos [post]
func CriaNovoAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	//busco no banco
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.Delete(&aluno, id)

	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso!"})
}

func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	//pesquisa aluno
	database.DB.First(&aluno, id)

	//empacota corpo da requisição
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//edita modelo aluno com souldbindJSON aluno (atualiza campos)
	database.DB.Model(&aluno).UpdateColumns(aluno)
	//mensagem 200, aluno
	c.JSON(http.StatusOK, aluno)
}

func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")

	//busca no db o aluno com CPF = cpf e devolve o primeiro no nosso endereço de aluno (var)
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}
