package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/panhdjf/scrum/controllers"
)

type TaskRouteController struct {
	taskController controllers.TaskController
}

func NewRouteTaskController(taskController controllers.TaskController) TaskRouteController {
	return TaskRouteController{taskController}
}

func (tc *TaskRouteController) TaskRouter(rg *gin.RouterGroup) {

	router := rg.Group("tasks")
	// router.Use(middleware.DeserializeUser())
	router.POST("/", tc.taskController.CreatedTask)
	router.GET("/", tc.taskController.FindTasks)
	router.PUT("/:taskId", tc.taskController.UpdateTask)
	router.GET("/:taskId", tc.taskController.FindTaskById)

	router.DELETE("/:taskId", tc.taskController.DeleteTask)

	router1 := rg.Group("task")
	router1.GET("/:sprint", tc.taskController.ManagerTask)
}
