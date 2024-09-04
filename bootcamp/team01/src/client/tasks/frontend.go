package tasks

import (
	"bufio"
	"fmt"
	"os"
	"replication/utils"
	"strings"
)

// What front does:
// Parse user inputs,
// Validate them,
// Send valid tasks over,

var (
	QUOTE                   = `'`
	DOUBLE_QUOTE            = `"`
	BACKTICK                = "`"
	ALLOWED_QUOTES []string = []string{QUOTE, DOUBLE_QUOTE, BACKTICK}
)

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
	ID      string
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
		// Need at least 2 params for a task: VERB KEY
		if len(task) < 2 {
			printTaskRules()
			continue
		}

		verb = task[0]
		id := task[1]

		if utils.IsUUIDValid(id) != nil {
			fmt.Println("incorrect UUID for KEY.")
			continue
		}

		if len(task) >= 3 && verb == SET {
			// only set tasks have 3 or more args
			value := parseMeaningfulString(task[2:])
			if len(value) == 0 {
				fmt.Println("Can't set empty value\nUse DELETE if you want to clear/delete the record")
				continue
			}
			f.taskChan <- Task{
				ID:      id,
				Verb:    verb,
				Payload: value,
			}
		} else if len(task) == 2 && (verb == GET || verb == DELETE) {
			// craete GET/DELETE task
			f.taskChan <- Task{
				ID:   id,
				Verb: verb,
			}
		} else {
			printTaskRules()
		}
	}
}

func printTaskRules() {
	fmt.Println("Wrong input plase follow the rules:")
	fmt.Println("syntax: [SET/GET/DELETE] [UUID] (VALIE)")
	fmt.Println("e.g: SET 9e90125c-180e-428d-8450-aa5bcaeafbb1 somevalue")
}

func checkQuotes(str string) (string, bool) {
	for _, quote := range ALLOWED_QUOTES {
		if strings.HasPrefix(str, quote) && strings.HasSuffix(str, quote) {
			return quote, true
		}
	}
	return "", false
}

func parseMeaningfulString(strs []string) string {
	joined := strings.Join(strs, " ")
	raw := strings.Trim(joined, " ")
	if quote, isQuoted := checkQuotes(raw); isQuoted {
		return strings.Trim(raw, quote)
	}

	if len(strs) > 0 {
		return strs[0]
	}
	return ""
}
