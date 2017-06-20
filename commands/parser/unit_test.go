package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildCopyCommand(t *testing.T) {
	unit := &unit{
		Name:        "copy",
		Path:        "/test/.env",
		Destination: "/remote_test/.env",
	}
	addr := "app@test.com"

	t.Run("with rsync", func(*testing.T) {
		interpreted, err := unit.buildCopyCommand(addr, rsync)
		assert.NoError(t, err)

		assert.Equal(t, "rsync -chavzP /test/.env app@test.com:/remote_test/.env", interpreted.command, "creates proper command")
		assert.Equal(t, "copy", interpreted.name, "creates proper name")
	})

	t.Run("with scp", func(*testing.T) {
		interpreted, err := unit.buildCopyCommand(addr, scp)
		assert.NoError(t, err)

		assert.Equal(t, "scp /test/.env app@test.com:/remote_test/.env", interpreted.command, "creates proper command")
		assert.Equal(t, "copy", interpreted.name, "creates proper name")
	})

	t.Run("with not supported strategy", func(*testing.T) {
		_, err := unit.buildCopyCommand(addr, "not_supported")
		assert.Error(t, err, "raises error")
	})
}

func TestBuildLocalCommand(t *testing.T) {
	unit := &unit{
		Name:    "echo",
		Command: "echo dupa",
	}

	interpreted := unit.buildLocalCommand()
	assert.Equal(t, "echo", interpreted.name, "creates proper name")
	assert.Equal(t, "echo dupa", interpreted.command, "creates proper comand")
}

func TestBuildRemoteCommand(t *testing.T) {
	unit := &unit{
		Name:    "echo",
		Command: "echo dupa",
	}
	addr := "app@test.com"

	interpreted := unit.buildRemoteCommand(addr)
	assert.Equal(t, "echo", interpreted.name, "creates proper name")
	assert.Equal(t, "ssh app@test.com echo dupa", interpreted.command, "creates proper comand")
}
