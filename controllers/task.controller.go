package controllers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/panhdjf/scrum/models"
	"gorm.io/gorm"
)

type TaskController struct {
	DB *gorm.DB
}

func NewTaskController(DB *gorm.DB) TaskController {
	return TaskController{DB}
}

func (tc *TaskController) CreatedTask(ctx *gin.Context) {
	// currentUser := ctx.MustGet("currentUser").(models.Task)
	var payload *models.CreateTaskRequest
	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	now := time.Now()
	newTask := models.Task{
		Name:        payload.Name,
		Description: payload.Description,
		Sprint:      payload.Sprint,
		Assignee:    payload.Assignee,
		StoryPoint:  payload.StoryPoint,
		Status:      payload.Status,
		CreateAt:    now,
		UpdateAt:    now,
	}
	new_task := newTask
	result := tc.DB.Create(&newTask)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "duplicate key") {
			ctx.JSON(http.StatusConflict, gin.H{"status": "fail", "message": "Task with that name already exists"})
			return
		}
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": new_task})
}

func (tc *TaskController) UpdateTask(ctx *gin.Context) {
	taskId := ctx.Param("taskId")
	// currentUser := ctx.MustGet("currentUser").(models.Task)

	var payload *models.UpdateTask
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	var updateTask models.Task
	result := tc.DB.First(&updateTask, "id = ?", taskId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No Task with that name already exists"})
		return
	}
	now := time.Now()
	taskToUpdate := models.Task{
		Name:        payload.Name,
		Description: payload.Description,
		Sprint:      payload.Sprint,
		Assignee:    payload.Assignee,
		StoryPoint:  payload.StoryPoint,
		Status:      payload.Status,
		CreateAt:    updateTask.CreateAt,
		UpdateAt:    now,
	}

	tc.DB.Model(&updateTask).Updates(taskToUpdate)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": updateTask})
}

func (tc *TaskController) FindTaskById(ctx *gin.Context) {
	taskId := ctx.Param("taskId")

	var task models.Task
	result := tc.DB.First(&task, "id = ?", taskId)
	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No Task with that name already exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": task})
}

func (tc *TaskController) FindTasks(ctx *gin.Context) {
	var page = ctx.DefaultQuery("page", "1")
	var limit = ctx.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var tasks []models.Task
	results := tc.DB.Limit(intLimit).Offset(offset).Find(&tasks)
	if results.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "results": len(tasks), "data": tasks})
}

func (tc *TaskController) DeleteTask(ctx *gin.Context) {
	taskId := ctx.Param("taskId")

	result := tc.DB.Delete(&models.Task{}, "id = ?", taskId)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"status": "fail", "message": "No Task with that name already exists"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "delete success"})
}
