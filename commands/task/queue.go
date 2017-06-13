package task

import (
	"fmt"

	"github.com/fatih/color"
)

type Queue struct {
	preTasks    []*Task
	deployTasks []*Task
	postTasks   []*Task
}

// Exec executes pre-tasks, deployment and post-tasks queues
func (q *Queue) Exec() {
	q.iterateAndExecute(q.preTasks, "Running pre-tasks...\n\n")
	q.iterateAndExecute(q.deployTasks, "Deploying...\n\n")
	q.iterateAndExecute(q.postTasks, "Running post-tasks...\n\n")
}

func (q *Queue) iterateAndExecute(queue []*Task, msg string) {
	if queueIsNotEmpty(queue) {
		q.print(msg)
		for _, task := range queue {
			task.Exec()
		}
	}
}

func (q *Queue) print(msg string) {
	fmt.Printf("%s", yellow(msg))
}

func queueIsNotEmpty(queue []*Task) bool {
	return len(queue) > 0
}

// pass task and push it to proper queue basing on type
func (q *Queue) appendTask(task *Task) {
	switch task.taskType {
	case preTask:
		q.preTasks = append(q.preTasks, task)
	case deployTask:
		q.deployTasks = append(q.deployTasks, task)
	case postTask:
		q.postTasks = append(q.postTasks, task)
	}
}
