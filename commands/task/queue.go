package task

import "fmt"

// Queue is a struct holding tasks for execution during deployment process. It is used
// as a "governor" since it's a structure which runs all the commands.
type Queue struct {
	preTasks    []*Task // tasks executed before actual deployment
	deployTasks []*Task // tasks of the deployment itself
	postTasks   []*Task // tasks executed after actual deployment
	length      int
}

func NewQueue() *Queue {
	return &Queue{
		preTasks:    make([]*Task, 0),
		deployTasks: make([]*Task, 0),
		postTasks:   make([]*Task, 0),
	}
}

func (q *Queue) Len() int {
	return q.length
}

// Exec executes pre-tasks, deployment tasks and post-tasks queues
func (q *Queue) Exec() (err error) {
	err = q.iterateAndExecute(q.preTasks, "Running pre-tasks...\n\n")
	if err != nil { // can't collect errors for processing later as it should fail right away
		return err
	}

	err = q.iterateAndExecute(q.deployTasks, "Deploying...\n\n")
	if err != nil {
		return err
	}

	err = q.iterateAndExecute(q.postTasks, "Running post-tasks...\n\n")
	return err
}

func (q *Queue) iterateAndExecute(queue []*Task, msg string) {
	if queueIsNotEmpty(queue) {
		q.print(msg)
		for _, task := range queue {
			err := task.exec()
			q.length--

			if err != nil {
				task.fail()
				return
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
func (q *Queue) Append(task *Task) error {
	switch task.taskType {
	case PreTask:
		q.preTasks = append(q.preTasks, task)
	case DeployTask:
		q.deployTasks = append(q.deployTasks, task)
	case PostTask:
		q.postTasks = append(q.postTasks, task)
	default:
		return fmt.Errorf("provided task does not belong to any valid queue")
	}
	q.length++
	return nil
}
