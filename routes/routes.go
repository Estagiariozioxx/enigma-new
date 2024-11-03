package routes

import (
    "enigma-new/controllers"
    "enigma-new/middlewares"
    "github.com/gin-gonic/gin"
    "github.com/swaggo/gin-swagger"  
    files "github.com/swaggo/files" 
)

func SetupRouter() *gin.Engine {
   
    router := gin.Default()

       router.GET("teste/swagger/doc.json", func(c *gin.Context) { //solução para o swagger localizar o json
        c.File("./docs/swagger.json")
    })

    router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler, ginSwagger.URL("http://localhost:8080/teste/swagger/doc.json"))) 


	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "testando o enigma",
		})
	})

    api := router.Group("/api")
    api.POST("/login", controllers.Loginn)  
	api.POST("/users", controllers.CreateUser)
    api.Use(middlewares.AuthMiddleware())
    {
        api.GET("/current-key", controllers.GetCurrentKey)
        api.GET("/keys", controllers.ListKeys)

        api.GET("/users", controllers.ListUsers) 
        api.PUT("/users/:id", controllers.UpdateUser) 
        api.DELETE("/users/:id", controllers.DeleteUser) 

        api.POST("/decrypt", controllers.DecryptDocument) 
    }

    return router
}
