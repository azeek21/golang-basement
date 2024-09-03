package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// What front does:
// Parse user inputs,
// Validate them,
// Send valid tasks over,

type EVerb = string

const (
	SET EVerb = "SET"
	// SET UUID VALIE
	GET = "GET"
	// GET UUID
	DELETE = "DELETE"
	// DELETE UUID
)

type Task struct {
	Verb    EVerb
	Payload interface{}
}

type Frontend struct {
	taskChan     chan Task
	rawTasksChan chan string
}

func NewFrontend(tasks chan Task) Frontend {
	rawTasks := make(chan string, 1)
	return Frontend{
		taskChan:     tasks,
		rawTasksChan: rawTasks,
	}
}

func (f Frontend) Start() {
	go f.parseAndStartTask()
	f.infiniteListenAndSendOver()
}

func (f Frontend) infiniteListenAndSendOver() {
	ln := ""
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ln = scanner.Text()
		f.rawTasksChan <- ln
	}
}

func (f Frontend) parseAndStartTask() {
	for ln := range f.rawTasksChan {
		var verb EVerb
		task := strings.Split(ln, " ")
		verb = task[0]
		if len(task) == 3 && verb == SET {
			// create set task
		} else if len(task) == 2 && (verb == GET || verb == DELETE) {
			// craete GET/DELETE task
		} else {
			fmt.Println("Wrong input plase follow the rules:")
			fmt.Println("syntax: [SET/GET/DELETE] [UUID] (VALIE)")
			fmt.Println("e.g: SET 9e90125c-180e-428d-8450-aa5bcaeafbb1 somevalue")
		}

		v := ""
		id := ""
		val := ""
		rd, err := fmt.Sscanf(ln, "%s %s %s", &v, &id, &val)
		if err != nil {
			fmt.Printf("ERR SCANF: %v: %v\n", rd, err.Error())
		}
		fmt.Println("SCANF: ", rd, v, id, val)
	}
}
