package parser

import (
	"testing"

	"github.com/piokaczm/godzira/commands/task"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {

	t.Run("with simple config", func(*testing.T) {
		// simple config consists of 2 basic tasks (test, deploy) per host, so we're asserting queue length
		t.Run("when staging", func(*testing.T) {
			queue := task.NewQueue()
			errs := Read(queue, "fixtures/simple_config.yml", "staging")
			assert.Len(t, errs, 0)

			assert.Equal(t, 2, queue.Len(), "adds 2 basic tasks")
		})

		t.Run("when production", func(*testing.T) {
			queue := task.NewQueue()
			errs := Read(queue, "fixtures/simple_config.yml", "production")
			assert.Len(t, errs, 0)

			assert.Equal(t, 3, queue.Len(), "adds 6 basic tasks") // test + 2*deploy
		})
	})

	t.Run("with pretasks config", func(*testing.T) {
		// simple config consists of 4 basic tasks (test, copy, echo, deploy) per host, so we're asserting queue length
		t.Run("when staging", func(*testing.T) {
			queue := task.NewQueue()
			errs := Read(queue, "fixtures/pretasks_config.yml", "staging")
			assert.Len(t, errs, 0)

			assert.Equal(t, 4, queue.Len(), "adds 4 basic tasks")
		})

		t.Run("when production", func(*testing.T) {
			queue := task.NewQueue()
			errs := Read(queue, "fixtures/pretasks_config.yml", "production")
			assert.Len(t, errs, 0)

			assert.Equal(t, 6, queue.Len(), "adds 6 basic tasks") // test + 2*copy + echo + 2*deploy
		})
	})

	t.Run("with posttasks config", func(*testing.T) {
		// simple config consists of 4 basic tasks (test, deploy, restart, echo) per host, so we're asserting queue length
		t.Run("when staging", func(*testing.T) {
			queue := task.NewQueue()
			errs := Read(queue, "fixtures/posttasks_config.yml", "staging")
			assert.Len(t, errs, 0)

			assert.Equal(t, 4, queue.Len(), "adds 4 tasks")
		})

		t.Run("when production", func(*testing.T) {
			queue := task.NewQueue()
			errs := Read(queue, "fixtures/posttasks_config.yml", "production")
			assert.Len(t, errs, 0)

			assert.Equal(t, 6, queue.Len(), "adds 6 tasks") // test + 2*deploy + 2*restart + echo
		})
	})

	t.Run("with posttasks and pretasks config", func(*testing.T) {
		// simple config consists of 5 basic tasks (test, copy, echo, deploy, restart, echo) per host, so we're asserting queue length
		t.Run("when staging", func(*testing.T) {
			queue := task.NewQueue()
			errs := Read(queue, "fixtures/post_and_pretasks_config.yml", "staging")
			assert.Len(t, errs, 0)

			assert.Equal(t, 6, queue.Len(), "adds 6 tasks")
		})

		t.Run("when production", func(*testing.T) {
			queue := task.NewQueue()
			errs := Read(queue, "fixtures/post_and_pretasks_config.yml", "production")
			assert.Len(t, errs, 0)

			assert.Equal(t, 9, queue.Len(), "adds 9 tasks") // test + 2*copy + echo + 2*deploy + 2*restart + echo
		})
	})
}

func TestInterpretSingleTask(t *testing.T) {
	conf, err := parse("fixtures/post_and_pretasks_config.yml") // 2 production hosts
	assert.NoError(t, err)
	conf.env = "production"

	t.Run("with copy task", func(*testing.T) {
		unit := &unit{
			Name:    "test",
			Command: "test command",
			Label:   copyLabel,
		}

		tasks, err := conf.interpretSingleTask(unit)
		assert.NoError(t, err)
		assert.Len(t, tasks, 2, "adds copy task for each host")
	})

	t.Run("with local task", func(*testing.T) {
		unit := &unit{
			Name:    "test",
			Command: "test command",
			Label:   localLabel,
		}

		tasks, err := conf.interpretSingleTask(unit)
		assert.NoError(t, err)
		assert.Len(t, tasks, 1, "adds local task once only")
	})

	t.Run("with remote task", func(*testing.T) {
		unit := &unit{
			Name:    "test",
			Command: "test command",
			Label:   remoteLabel,
		}

		tasks, err := conf.interpretSingleTask(unit)
		assert.NoError(t, err)
		assert.Len(t, tasks, 2, "adds remote task each host")
	})

	t.Run("with task with unsupported label", func(*testing.T) {
		unit := &unit{
			Name:    "test",
			Command: "test command",
			Label:   "unsupported label",
		}

		_, err := conf.interpretSingleTask(unit)
		assert.Error(t, err, "raises and error")
	})
}

func TestAppendTask(t *testing.T) {
	cr := configReader{queue: task.NewQueue()}

	cr.appendTask("test", "test command", 1)
	assert.Len(t, cr.errors, 0, "no errors occured")
	assert.Equal(t, 1, cr.queue.Len(), "appends the task")
}
