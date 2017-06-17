package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue()
	assert.NotNil(t, queue.preTasks, "initializes pretasks array")
	assert.NotNil(t, queue.deployTasks, "initializes deploy tasks array")
	assert.NotNil(t, queue.postTasks, "initializes posttasks array")
}

func TestAppend(t *testing.T) {
	t.Run("with pretask", func(t *testing.T) {
		queue := NewQueue()
		task, err := NewTask("test", "test", PreTask)
		assert.NoError(t, err)

		err = queue.Append(task)
		assert.NoError(t, err)
		assert.Equal(t, 1, queue.Len(), "updates length")
		assert.Len(t, queue.preTasks, 1, "append to proper queue")
		assert.Equal(t, task, queue.preTasks[0], "appends the task itself")
	})

	t.Run("with posttask", func(t *testing.T) {
		queue := NewQueue()
		task, err := NewTask("test", "test", PostTask)
		assert.NoError(t, err)

		err = queue.Append(task)
		assert.NoError(t, err)
		assert.Equal(t, 1, queue.Len(), "updates length")
		assert.Len(t, queue.postTasks, 1, "append to proper queue")
		assert.Equal(t, task, queue.postTasks[0], "appends the task itself")
	})

	t.Run("with deploy task", func(t *testing.T) {
		queue := NewQueue()
		task, err := NewTask("test", "test", DeployTask)
		assert.NoError(t, err)

		err = queue.Append(task)
		assert.NoError(t, err)
		assert.Equal(t, 1, queue.Len(), "updates length")
		assert.Len(t, queue.deployTasks, 1, "append to proper queue")
		assert.Equal(t, task, queue.deployTasks[0], "appends the task itself")
	})

	t.Run("with not supported type", func(t *testing.T) {
		queue := NewQueue()
		task, err := NewTask("test", "test", 10)
		assert.NoError(t, err)

		err = queue.Append(task)
		assert.Error(t, err, "raises an error")
		assert.Equal(t, 0, queue.Len(), "does not increment length")
	})
}

func TestQueueExec(t *testing.T) {
	queue := NewQueue()

	t.Run("successful execution", func(t *testing.T) {
		task, err := NewTask("echo", "echo", DeployTask)
		assert.NoError(t, err)
		err = queue.Append(task)
		assert.NoError(t, err)

		assert.Equal(t, 1, queue.Len(), "before execution there is a task on queue")

		err = queue.Exec()
		assert.NoError(t, err)
		assert.Zero(t, queue.Len(), "after execution there is no task on queue")
	})

	t.Run("unsuccessful execution", func(t *testing.T) {
		task, err := NewTask("failing", "cp non existing", PreTask)
		assert.NoError(t, err)
		err = queue.Append(task)
		assert.NoError(t, err)

		task, err = NewTask("echo", "echo", PreTask)
		assert.NoError(t, err)
		err = queue.Append(task)
		assert.NoError(t, err)

		assert.Equal(t, 2, queue.Len(), "before execution there are 2 tasks on queue")

		err = queue.Exec()
		assert.Error(t, err)
		assert.Equal(t, 1, queue.Len(), "after fail the remaining task was not executed")
	})
}

func TestQueueIsNotEmpty(t *testing.T) {
	queue := NewQueue()
	task, err := NewTask("test", "test", DeployTask)
	assert.NoError(t, err)
	err = queue.Append(task)
	assert.NoError(t, err)

	assert.False(t, queueIsNotEmpty(queue.postTasks), "returns false for empty queue")
	assert.True(t, queueIsNotEmpty(queue.deployTasks), "returns true for non empty queue")
}
