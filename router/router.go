package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// função que inicializa o router do gin
	router := gin.Default()
	// definindo a rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// rodando a api
	router.Run(":8080")
}
