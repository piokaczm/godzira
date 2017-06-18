package parser

import (
	"testing"

	"github.com/piokaczm/godzira/commands/task"
	"github.com/stretchr/testify/assert"
)

func TestParsing(t *testing.T) {
	path := "fixtures/simple_config.yml"

	t.Run("with existing environment", func(t *testing.T) {
		env := "staging"
		config, err := New(path, env)
		assert.NoError(t, err)

		// basic data
		assert.Equal(t, config.Strategy, "scp")
		assert.Equal(t, config.Test, true)
		assert.Equal(t, config.Name, "test_app")
		assert.Equal(t, config.BinPath, "test_name")

		// slack data
		assert.Equal(t, config.Slack.Channel, "testchannel")
		assert.Equal(t, config.Slack.Name, "testname")
		assert.Equal(t, config.Slack.Webhook, "testwebhook")
		assert.Equal(t, config.Slack.Emoji, ":test:")

		// environments
		assert.Len(t, config.Environments["staging"], 1, "adds all staging hosts")
		assert.Equal(t, config.Environments["staging"][0].Host, "pizdki.net")
		assert.Equal(t, config.Environments["staging"][0].User, "pizdek")
		assert.Equal(t, config.Environments["staging"][0].Path, "pizdek/app/")

		assert.Len(t, config.Environments["production"], 2, "adds all staging hosts")
		assert.Equal(t, config.Environments["production"][0].Host, "real.net")
		assert.Equal(t, config.Environments["production"][0].User, "app1")
		assert.Equal(t, config.Environments["production"][0].Path, "current/binaries/")
		assert.Equal(t, config.Environments["production"][1].Host, "real2.net")
		assert.Equal(t, config.Environments["production"][1].User, "app2")
		assert.Equal(t, config.Environments["production"][1].Path, "current/binaries/")
	})

	t.Run("with non existing environment", func(t *testing.T) {
		env := "not existing"
		_, err := New(path, env)
		assert.Error(t, err)
	})

	t.Run("with non malformed file", func(t *testing.T) {
		path = "fixtures/malformed_file.yml"
		env := "not existing"
		_, err := New(path, env)
		assert.Error(t, err)
	})
}

func TestRead(t *testing.T) {
	t.Run("with simple config", func(*testing.T) {
		// simple config consists of 2 basic tasks (test, deploy) per host, so we're asserting queue length
		t.Run("when staging", func(*testing.T) {
			config, err := New("fixtures/simple_config.yml", "staging")
			assert.NoError(t, err)
			queue := task.NewQueue()
			errs := Read(config, queue)
			assert.Len(t, errs, 0)

			assert.Equal(t, 2, queue.Len(), "adds 2 basic tasks")
		})

		t.Run("when production", func(*testing.T) {
			config, err := New("fixtures/simple_config.yml", "production")
			assert.NoError(t, err)
			queue := task.NewQueue()
			errs := Read(config, queue)
			assert.Len(t, errs, 0)

			assert.Equal(t, 3, queue.Len(), "adds 6 basic tasks") // test + 2*deploy
		})
	})

	t.Run("with pretasks config", func(*testing.T) {
		// simple config consists of 4 basic tasks (test, copy, echo, deploy) per host, so we're asserting queue length
		t.Run("when staging", func(*testing.T) {
			config, err := New("fixtures/pretasks_config.yml", "staging")
			assert.NoError(t, err)
			queue := task.NewQueue()
			errs := Read(config, queue)
			assert.Len(t, errs, 0)

			assert.Equal(t, 4, queue.Len(), "adds 4 basic tasks")
		})

		t.Run("when production", func(*testing.T) {
			config, err := New("fixtures/pretasks_config.yml", "production")
			assert.NoError(t, err)
			queue := task.NewQueue()
			errs := Read(config, queue)
			assert.Len(t, errs, 0)

			assert.Equal(t, 6, queue.Len(), "adds 6 basic tasks") // test + 2*copy + echo + 2*deploy
		})
	})

	t.Run("with posttasks config", func(*testing.T) {
		// simple config consists of 4 basic tasks (test, deploy, restart, echo) per host, so we're asserting queue length
		t.Run("when staging", func(*testing.T) {
			config, err := New("fixtures/posttasks_config.yml", "staging")
			assert.NoError(t, err)
			queue := task.NewQueue()
			errs := Read(config, queue)
			assert.Len(t, errs, 0)

			assert.Equal(t, 4, queue.Len(), "adds 4 tasks")
		})

		t.Run("when production", func(*testing.T) {
			config, err := New("fixtures/posttasks_config.yml", "production")
			assert.NoError(t, err)
			queue := task.NewQueue()
			errs := Read(config, queue)
			assert.Len(t, errs, 0)

			assert.Equal(t, 6, queue.Len(), "adds 6 tasks") // test + 2*deploy + 2*restart + echo
		})
	})

	t.Run("with posttasks and pretasks config", func(*testing.T) {
		// simple config consists of 5 basic tasks (test, copy, echo, deploy, restart, echo) per host, so we're asserting queue length
		t.Run("when staging", func(*testing.T) {
			config, err := New("fixtures/post_and_pretasks_config.yml", "staging")
			assert.NoError(t, err)
			queue := task.NewQueue()
			errs := Read(config, queue)
			assert.Len(t, errs, 0)

			assert.Equal(t, 6, queue.Len(), "adds 6 tasks")
		})

		t.Run("when production", func(*testing.T) {
			config, err := New("fixtures/post_and_pretasks_config.yml", "production")
			assert.NoError(t, err)
			queue := task.NewQueue()
			errs := Read(config, queue)
			assert.Len(t, errs, 0)

			assert.Equal(t, 9, queue.Len(), "adds 9 tasks") // test + 2*copy + echo + 2*deploy + 2*restart + echo
		})
	})

	t.Run("with malformed config", func(*testing.T) {
		config, err := New("fixtures/malformed_config.yml", "staging")
		assert.NoError(t, err)
		queue := task.NewQueue()
		errs := Read(config, queue)
		assert.Len(t, errs, 4, "appends error for each malformation")
	})
}

func TestInterpretSingleTask(t *testing.T) {
	conf, err := New("fixtures/post_and_pretasks_config.yml", "production") // 2 production hosts
	assert.NoError(t, err)

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

	cr.appendTask("test", "test command", "host", 1)
	assert.Len(t, cr.errors, 0, "no errors occured")
	assert.Equal(t, 1, cr.queue.Len(), "appends the task")
}
