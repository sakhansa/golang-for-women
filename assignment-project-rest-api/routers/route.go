package routers

import (
	"assignment-project-rest-api/controllers"
	"assignment-project-rest-api/database"

	"github.com/gin-gonic/gin"
)

func init() {
	database.StartDB()
}

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/student", controllers.CreateStudent)
	router.GET("/students", controllers.GetAllStudent)
	router.PUT("/student/:studentID", controllers.UpdateStudent)
	router.DELETE("/student/:studentID", controllers.DeleteStudent)
	return router
}
