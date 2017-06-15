package parser

import (
	"testing"

	"github.com/piokaczm/godzira/commands/task"
	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	queue := &task.Queue{}

	t.Run("with simple config", func(*testing.T) {
		// simple config consists of 3 basic tasks (test, build, deploy) per host, so we're asserting queue length
		t.Run("when staging", func(*testing.T) {
			err := Read(queue, "fixtures/simple_config.yml", "staging")
			assert.NoError(t, err)

			assert.Equal(t, 3, queue.Len(), "adds 3 basic tasks")
		})

		t.Run("when production", func(*testing.T) {
			err := Read(queue, "fixtures/simple_config.yml", "production")
			assert.NoError(t, err)

			assert.Equal(t, 6, queue.Len(), "adds 6 basic tasks") // 3 tasks * 2 hosts
		})
	})
}
