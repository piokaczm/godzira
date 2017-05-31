package task

type Queue struct {
	preTasks    []*Task
	deployTasks []*Task
	postTasks   []*Task
}

func (q *Queue) Exec() {
	if queueIsNotEmpty(q.preTasks) {
		q.print("Running pretasks...\n\n")
		for _, task := range q.preTasks {
			task.Exec()
		}
	}

	q.print("Deploying...\n\n")
	for _, task := range q.deployTasks {
		task.Exec()
	}

	if queueIsNotEmpty(q.postTasks) {
		q.print("Running posttasks...\n\n")
		for _, task := range q.deployTask {
			task.Exec()
		}
	}
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
