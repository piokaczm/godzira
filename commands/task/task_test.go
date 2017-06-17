package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTask(t *testing.T) {
	t.Run("with valid data", func(t *testing.T) {
		task, err := NewTask("test", "test -p command", PreTask)
		assert.NoError(t, err)

		assert.Equal(t, "test", task.name, "sets valid name")
		assert.Equal(t, PreTask, task.taskType, "sets valid type")
		assert.Equal(t, "test", task.cmd.name, "sets calid command name")
		assert.Equal(t, []string{"-p", "command"}, task.cmd.args, "sets calid command name")
	})

	t.Run("with invalid data", func(t *testing.T) {
		_, err := NewTask("test", "", PreTask)
		assert.Error(t, err, "raises error")
	})
}

func TestTaskExec(t *testing.T) {
	t.Run("successful execution", func(t *testing.T) {
		task, err := NewTask("echo", "echo tests", PreTask)
		assert.NoError(t, err)

		err = task.exec()
		assert.NoError(t, err, "doesn't raise an error")
		assert.Nil(t, task.err, "doesn't set an error value in task struct")
		assert.Equal(t, []byte("tests\n"), task.output, "saves command output")
	})

	t.Run("unsuccessful execution", func(t *testing.T) {
		task, err := NewTask("copy", "cp non existing", PreTask)
		assert.NoError(t, err)

		err = task.exec()
		assert.Error(t, err, "raises an error")
		assert.NotNil(t, task.err, "sets an error value in task struct")
		assert.Regexp(t, "cp.*non.*No such file or directory", string(task.output), "saves command output")
	})
}
