package parser

import (
	"testing"

	"github.com/piokaczm/godzira/commands/task"
	// "github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	queue := &task.Queue{}
	Read(queue, "fixtures/simple_config.yml")
}
