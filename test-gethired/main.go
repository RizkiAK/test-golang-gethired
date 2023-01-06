package main

import (
	"test-gethired/activities"
	"test-gethired/db"
	"test-gethired/todos"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.InitDB()

	activityRepository := activities.NewRepository(db)
	activityService := activities.NewService(activityRepository)
	activityHandler := activities.NewHandler(activityService)

	todoRepository := todos.NewRepository(db)
	todoService := todos.NewService(todoRepository)
	todoHandler := todos.NewHandler(todoService)

	router := gin.Default()
	activityAPI := router.Group("/")
	todoAPI := router.Group("/")

	activityAPI.GET("/activity-groups", activityHandler.GetAll)
	activityAPI.GET("/activity-groups/:id", activityHandler.GetOne)
	activityAPI.POST("/activity-groups", activityHandler.Created)
	activityAPI.PATCH("/activity-groups/:id", activityHandler.Updated)
	activityAPI.DELETE("/activity-groups/:id", activityHandler.Delete)

	todoAPI.GET("/todo-items", todoHandler.GetAll)
	todoAPI.GET("/todo-items/:id", todoHandler.GetOne)
	todoAPI.POST("/todo-items", todoHandler.Created)
	todoAPI.PATCH("/todo-items/:id", todoHandler.Updated)
	todoAPI.DELETE("/todo-items/:id", todoHandler.Delete)

	router.Run(":3030")
}
