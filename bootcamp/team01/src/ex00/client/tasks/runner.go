package tasks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"replication/models"
	"replication/utils"
)

// Accepts tasks from frontend
// Call clients and makse sure replication is successfull
// * send results over to client

type Runner interface {
	Start()
}

type runner struct {
	taskChan chan Task
}

func NewRunner(taskChan chan Task) Runner {
	return runner{
		taskChan: taskChan,
	}
}

func (r runner) Start() {
	r.ConsumeTasks(TaskWorker)
}

func (r runner) ConsumeTasks(worker func(task Task)) {
	for task := range r.taskChan {
		go worker(task)
	}
}

func TaskWorker(task Task) {
	payload, err := json.Marshal(models.StorageItem{
		Key:   task.ID,
		Value: task.Payload.(string),
	})

	if err != nil {
		fmt.Println(utils.WithPrefix("task: %v: ", err).Error())
		return
	}

	resp, err := http.Post("http://localhost:8080/api/storage/", "application/json", bytes.NewReader(payload))

	if err != nil {
		fmt.Println(utils.WithPrefix("task: %v: ", err).Error())
		return
	}

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(utils.WithPrefix("task: %v: ", err).Error())
		}
		fmt.Printf("RESP: %s\n", string(bodyBytes))
		return
	}

	if resp.StatusCode == http.StatusBadRequest {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(utils.WithPrefix("task: %v: ", err).Error())
		}
		fmt.Printf("BAD REQUEST: %s\n", string(bodyBytes))
		return
	}
}
