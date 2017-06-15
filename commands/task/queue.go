package task

import (
	"fmt"
	"os"
)

// Queue is a struct holding tasks for execution during deployment process. It is used
// as a "governor" since it's a structure which runs all the commands.
type Queue struct {
	preTasks    []*Task // tasks executed before actual deployment
	deployTasks []*Task // tasks of the deployment itself
	postTasks   []*Task // tasks executed after actual deployment
}

// Exec executes pre-tasks, deployment tasks and post-tasks queues
func (q *Queue) Exec() {
	q.iterateAndExecute(q.preTasks, "Running pre-tasks...\n\n")
	q.iterateAndExecute(q.deployTasks, "Deploying...\n\n")
	q.iterateAndExecute(q.postTasks, "Running post-tasks...\n\n")
}

func (q *Queue) iterateAndExecute(queue []*Task, msg string) {
	if queueIsNotEmpty(queue) {
		q.print(msg)
		for _, task := range queue {
			err := task.exec()

			if err != nil {
				task.fail()
				os.Exit(2)
			}
		}
	}
}

func (q *Queue) print(msg string) {
	fmt.Printf("%s", yellow(msg))
}

func queueIsNotEmpty(queue []*Task) bool {
	return len(queue) > 0
}

// Append passes task and pushes it to a proper queue basing on its' type
func (q *Queue) Append(task *Task) {
	switch task.taskType {
	case preTask:
		q.preTasks = append(q.preTasks, task)
	case deployTask:
		q.deployTasks = append(q.deployTasks, task)
	case postTask:
		q.postTasks = append(q.postTasks, task)
	}
}
