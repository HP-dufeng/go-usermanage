package dtos

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/dufeng/usermanager/models"
)

type (

	// TaskDto data transfer object for Task
	TaskDto struct {
		ID          bson.ObjectId `json:"id"`
		CreatedBy   string        `json:"createdby"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Due         time.Time     `json:"due,omitempty"`
		Status      string        `json:"status,omitempty"`
		Tags        []string      `json:"tags,omitempty"`
	}
)

type (
	// CreateTaskInput POST /tasks
	CreateTaskInput struct {
		CreatedBy   string   `json:"createdby"`
		Name        string   `json:"name"`
		Description string   `json:"description"`
		Tags        []string `json:"tags,omitempty"`
	}

	// CreateTaskOutput POST /tasks
	CreateTaskOutput struct {
		TaskDto
	}

	// GetTasksOutput GET /tasks
	GetTasksOutput struct {
		Data []TaskDto `json:"data"`
	}

	// GetTasksByUserOutput Get /tasks/users/{id}
	GetTasksByUserOutput struct {
		Data []TaskDto `json:"data"`
	}

	// UpdateTaskInput PUT /tasks/{id}
	UpdateTaskInput struct {
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Due         time.Time `json:"due,omitempty"`
		Status      string    `json:"status,omitempty"`
		Tags        []string  `json:"tags,omitempty"`
	}
)

// MapToTaskEntity convert CreateTaskInput to Task entity
func (m *CreateTaskInput) MapToTaskEntity() models.Task {
	return models.Task{
		CreatedBy:   m.CreatedBy,
		Name:        m.Name,
		Description: m.Description,
		Tags:        m.Tags,
	}
}

// MapToTaskDto convert Task entity to TaskDto
func MapToTaskDto(m *models.Task) TaskDto {
	return TaskDto{
		ID:          m.Id,
		CreatedBy:   m.CreatedBy,
		Name:        m.Name,
		Description: m.Description,
		CreatedOn:   m.CreatedOn,
		Due:         m.Due,
		Status:      m.Status,
		Tags:        m.Tags,
	}
}
