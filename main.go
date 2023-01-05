package main

import (
	"test-gethired/activity"
	"test-gethired/db"
	"test-gethired/todo"

	"github.com/gin-gonic/gin"
)

func main() {
	db := db.InitDB()

	activityRepository := activity.NewRepository(db)
	activityService := activity.NewService(activityRepository)
	activityHandler := activity.NewHandler(activityService)

	todoRepository := todo.NewRepository(db)
	todoService := todo.NewService(todoRepository)
	todoHandler := todo.NewHandler(todoService)

	router := gin.Default()
	activityAPI := router.Group("/todolist.api/activity-groups")
	todoAPI := router.Group("/todolist.api/todo-item")

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
