package parser

import (
	"testing"

	"github.com/piokaczm/godzira/commands/task"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	queue := &task.Queue{}

	t.Run("with simple config", func(*testing.T) {
		// simple config consists of 2 basic tasks (test, deploy) per host, so we're asserting queue length
		t.Run("when staging", func(*testing.T) {
			errs := Read(queue, "fixtures/simple_config.yml", "staging")
			assert.Len(t, errs, 0)

			assert.Equal(t, 2, queue.Len(), "adds 2 basic tasks")
		})

		t.Run("when production", func(*testing.T) {
			errs := Read(queue, "fixtures/simple_config.yml", "production")
			assert.Len(t, errs, 0)

			assert.Equal(t, 3, queue.Len(), "adds 6 basic tasks") // test + 2*deploy
		})
	})

	t.Run("with pretasks config", func(*testing.T) {
		// simple config consists of 4 basic tasks (test, copy, echo, deploy) per host, so we're asserting queue length
		t.Run("when staging", func(*testing.T) {
			errs := Read(queue, "fixtures/pretasks_config.yml", "staging")
			assert.Len(t, errs, 0)

			assert.Equal(t, 4, queue.Len(), "adds 4 basic tasks")
		})

		t.Run("when production", func(*testing.T) {
			errs := Read(queue, "fixtures/pretasks_config.yml", "production")
			assert.Len(t, errs, 0)

			assert.Equal(t, 6, queue.Len(), "adds 6 basic tasks") // test + 2*copy + echo + 2*deploy
		})
	})

	t.Run("with posttasks config", func(*testing.T) {
		// simple config consists of 4 basic tasks (test, deploy, restart, echo) per host, so we're asserting queue length
		t.Run("when staging", func(*testing.T) {
			errs := Read(queue, "fixtures/posttasks_config.yml", "staging")
			assert.Len(t, errs, 0)

			assert.Equal(t, 4, queue.Len(), "adds 4 tasks")
		})

		t.Run("when production", func(*testing.T) {
			errs := Read(queue, "fixtures/posttasks_config.yml", "production")
			assert.Len(t, errs, 0)

			assert.Equal(t, 6, queue.Len(), "adds 6 tasks") // test + 2*deploy + 2*restart + echo
		})
	})

	t.Run("with posttasks and pretasks config", func(*testing.T) {
		// simple config consists of 5 basic tasks (test, copy, echo, deploy, restart, echo) per host, so we're asserting queue length
		t.Run("when staging", func(*testing.T) {
			errs := Read(queue, "fixtures/post_and_pretasks_config.yml", "staging")
			assert.Len(t, errs, 0)

			assert.Equal(t, 6, queue.Len(), "adds 6 tasks")
		})

		t.Run("when production", func(*testing.T) {
			errs := Read(queue, "fixtures/post_and_pretasks_config.yml", "production")
			assert.Len(t, errs, 0)

			assert.Equal(t, 9, queue.Len(), "adds 9 tasks") // test + 2*copy + echo + 2*deploy + 2*restart + echo
		})
	})
}
