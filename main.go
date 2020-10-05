package main

import (
	"github.com/gin-contrib/cors" 
	"github.com/gin-gonic/gin"
	"manipulasi/controllers"
)

func main() {
	router := gin.Default() 

	router.Use(cors.Default()) 
	router.LoadHTMLGlob("views/*")
	router.GET("/", func(c *gin.Context){
		c.HTML(200, "index.html",nil)
	})
	router.GET("/detail/:id", func(c *gin.Context){
		c.HTML(200, "detail.html",gin.H{
			"data" : controllers.GetOneCompanies,
		})
	})
	router.GET("/getAllCompanies", controllers.GetAllCompanies)
	router.POST("/postOneCompanies", controllers.PostOneCompanies)

	router.GET("/detailin/:id", controllers.GetOneCompanies)
	router.POST("/update", controllers.UpdateOneCompanies)

	router.Run(":8000")
}