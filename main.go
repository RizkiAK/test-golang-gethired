package main

import (
	"test-gethired/activites"
	"test-gethired/db"
	"test-gethired/todos"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.InitDB()

	activityRepository := activites.NewRepository(db)
	activityService := activites.NewService(activityRepository)
	activityHandler := activites.NewHandler(activityService)

	todoRepository := todos.NewRepository(db)
	todoService := todos.NewService(todoRepository)
	todoHandler := todos.NewHandler(todoService)

	router := gin.Default()
	activityAPI := router.Group("/todolist.api.devcode.gethired.id/activity-groups")
	todoAPI := router.Group("/todolist.api.devcode.gethired.id/todo-item")

	activityAPI.GET("/", activityHandler.GetAll)
	activityAPI.GET("/:id", activityHandler.GetOne)
	activityAPI.POST("/", activityHandler.Created)
	activityAPI.PATCH("/:id", activityHandler.Updated)
	activityAPI.DELETE("/:id", activityHandler.Delete)

	todoAPI.GET("/", todoHandler.GetAll)
	todoAPI.GET("/:id", todoHandler.GetOne)
	todoAPI.POST("/", todoHandler.Created)
	todoAPI.PATCH("/:id", todoHandler.Updated)
	todoAPI.DELETE("/:id", todoHandler.Delete)

	router.Run(":3030")
}
