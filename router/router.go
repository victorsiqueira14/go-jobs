package router

import "github.com/gin-gonic/gin"

// fuction to initialize the router
func Initialize() {
	router := gin.Default()

	InitializeRoutes(router)

	router.Run(":8080")
}
