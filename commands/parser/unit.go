package parser

import "fmt"

type unit struct {
	Name        string `yaml:"name"`
	Command     string `yaml:"command"`
	Path        string `yaml:"path"`
	Destination string `yaml:"destination"`
	Label       string `yaml:"type"`
}

// local tasks doesn't have to be fired for each remote machine, so basing on provided
// labels godzira can decide if task should be appended
type interpretedUnit struct {
	name    string
	command string
}

func (u *unit) buildCopyCommand(addr, strategy string) (*interpretedUnit, error) {
	interpretation := &interpretedUnit{name: u.Name}
	var err error

	switch strategy {
	case rsync:
		interpretation.command = fmt.Sprintf("%s %s %s %s:%s", rsync, rsyncArg, u.Path, addr, u.Destination)
	case scp:
		interpretation.command = fmt.Sprintf("%s %s %s:%s", scp, u.Path, addr, u.Destination)
	default:
		err = unsupportedStrategy(u.Name, strategy)
	}
	return interpretation, err
}

func (u *unit) buildLocalCommand() *interpretedUnit {
	return &interpretedUnit{u.Name, u.Command}
}

func (u *unit) buildRemoteCommand(addr string) *interpretedUnit {
	return &interpretedUnit{u.Name, fmt.Sprintf("ssh %s %s", addr, u.Command)}
}
