package main

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const (
	rsync = "rsync"
	scp   = "scp"
)

type Config struct {
	Mode         string                         `yaml:"mode"`
	Name         string                         `yaml:"binary_name"`
	Strategy     string                         `yaml:"strategy"`
	Goos         string                         `yaml:"goos"`
	Goarch       string                         `yaml:"goarch"`
	Environments map[string]map[string][]string `yaml:"environments"`
	Slack        map[string]string              `yaml:"slack"`
	Test         bool                           `yaml:"test"`
	Vendor       bool                           `yaml:"vendor"`
	CurrentEnv   string
}

type Environment struct {
	Hosts []string
	User  string
	Name  string
}

// load yml with config, prepare tasks and slack config based on it
func loadConfig(path string, currentEnv string) (*Config, *Slack, *Environment, []*Task) {
	currentUser := os.Getenv("USER")
	c := Config{CurrentEnv: currentEnv}
	// load yml
	data, err := ioutil.ReadFile(path) // maybe use abs here
	checkErr(err)
	err = yaml.Unmarshal([]byte(data), &c)
	checkErr(err)
	// check strategy
	c.chooseStrategy()
	// prepare slack config
	slack := loadSlack(&c, currentUser)
	// prepare env
	env := loadEnv(&c, currentUser)
	// prepare tasks
	return &c, slack, env, []*Task{}
}

// return pointer to Environment struct with current env variables
func loadEnv(c *Config, currentUser string) *Environment {
	return &Environment{
		Hosts: c.getHosts(),
		User:  currentUser,
		Name:  c.CurrentEnv,
	}
}

// retrieve hosts from config, panic if no hosts found.
func (c *Config) getHosts() []string {
	hosts := c.Environments[c.CurrentEnv]["hosts"]
	if len(hosts) == 0 {
		err := errors.New("No hosts found in config file!")
		panic(err)
	}
	return hosts
}

// if no strategy specified fallback to rsync
func (c *Config) chooseStrategy() {
	if blank(c.Strategy) {
		c.Strategy = rsync
	}
}
