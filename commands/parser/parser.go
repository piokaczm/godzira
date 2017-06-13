package parser

func Read() {
	// for config parts
	// - read a part of config
	// - check for special labels (local, remote, copy)
	// - interpret a single task
	// - create new task basing on interpretation
	// - append to a global queue
}

func parse() {
	// actually parse the config
	// return 3 data structures with raw data, one for each step of deployment
}

type configReader struct {
	errors []error
}

func (cr *configReader) read() {
	// read a raw task and govern its' interpretation process
	// in case of an error just append it to errors which will be
	// printed out using Fail method.
}

func (cr *configReader) Fail() {
	// iterate over errors, print them and gracefully terminate the execution
}
