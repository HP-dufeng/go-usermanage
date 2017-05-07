package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dufeng/usermanager/common"
	"github.com/dufeng/usermanager/controllers/dtos"
	"github.com/dufeng/usermanager/data"
	"github.com/dufeng/usermanager/models"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// CreateTask POST /tasks
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var inputData dtos.CreateTaskInput
	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Task data", 500)
		return
	}

	taskEntity := inputData.MapToTaskEntity()
	context := NewContext()
	defer context.Close()

	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}

	// Insert a task
	repo.Create(&taskEntity)

	outputData := dtos.CreateTaskOutput{dtos.MapToTaskDto(&taskEntity)}
	j, err := json.Marshal(outputData)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

// GetTasks GET /tasks
func GetTasks(w http.ResponseWriter, r *http.Request) {
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}

	var taskDtos []dtos.TaskDto
	for _, item := range repo.GetAll() {
		taskDtos = append(taskDtos, dtos.MapToTaskDto(&item))
	}
	outputData := dtos.GetTasksOutput{Data: taskDtos}
	j, err := json.Marshal(outputData)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// GetTaskById GET /tasks/{id}
func GetTaskById(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]

	context := NewContext()
	defer context.Close()

	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}

	task, err := repo.GetById(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			common.DisplayAppError(w, err, "An unexpected error has occured", 500)
		}

		return
	}

	j, err := json.Marshal(task)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// GetTasksByUser GET /tasks/users/{id}
func GetTasksByUser(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	userId := vars["id"]

	context := NewContext()
	defer context.Close()
	col := context.DbCollection("tasks")
	repo := &data.TaskRepository{C: col}

	var taskDtos []dtos.TaskDto
	for _, item := range repo.GetByUser(userId) {
		taskDtos = append(taskDtos, dtos.MapToTaskDto(&item))
	}
	outputData := dtos.GetTasksByUserOutput{Data: taskDtos}
	j, err := json.Marshal(outputData)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// UpdateTask PUT /tasks/{id}
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := bson.ObjectIdHex(vars["id"])
	var inputData dtos.UpdateTaskInput
	err := json.NewDecoder(r.Body).Decode(&inputData)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid Task data", 500)
		return
	}

	taskEntity := models.Task{
		Id:          id,
		Name:        inputData.Name,
		Description: inputData.Description,
		Due:         inputData.Due,
		Status:      inputData.Status,
		Tags:        inputData.Tags,
	}
	context := NewContext()
	defer context.Close()
	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}

	if err := repo.Update(&taskEntity); err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	context := NewContext()
	defer context.Close()

	c := context.DbCollection("tasks")
	repo := &data.TaskRepository{c}

	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(w, err, "An unexpected error has occurred", 500)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
