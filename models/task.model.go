package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	Id          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id,omitempty"`
	Name        string    `gorm:"varchar(500);nol null" json:"name,omitempty"`
	Description string    `gorm:"varchar(1000);nol null" json:"description,omitempty"`
	Sprint      int       `gorm:"not null" json:"sprint,omitempty"`
	Assignee    string    `gorm:"not null" json:"assignee,omitempty"`
	StoryPoint  int       `gorm:"not null" json:"storyPoint,omitempty"`
	Status      string    `gorm:"not null" json:"status,omitempty"`
	CreateAt    time.Time `gorm:"not null" json:"createAt,omitempty"`
	UpdateAt    time.Time `gorm:"not null" json:"updateAt,omitempty"`
}

type CreateTaskRequest struct {
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Sprint      int       `json:"sprint" binding:"required"`
	Assignee    string    `json:"assignee" binding:"required"`
	StoryPoint  int       `json:"storyPoint" binding:"required"`
	Status      string    `json:"status" binding:"required"`
	CreateAt    time.Time `json:"createAt,omitempty"`
	UpdateAt    time.Time `json:"updateAt,omitempty"`
}

type UpdateTask struct {
	Name         string    `json:"name,omitempty"`
	Description  string    `json:"description,omitempty"`
	Sprint       int       `json:"sprint,omitempty"`
	Assignee     string    `json:"assignee,omitempty"`
	StoryPoint   int       `json:"storyPoint,omitempty"`
	Status       string    `json:"status,omitempty"`
	CreateAt     time.Time `json:"createAt,omitempty"`
	LastUpdateAt time.Time `json:"updateAt,omitempty"`
}
