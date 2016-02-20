package commands

import (
	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func Deploy(c *cli.Context) {
}

func parseConfig(data []byte) Configuration {
	result := Configuration{}
	err := yaml.Unmarshal([]byte(data), &result)
	checkErr(err)

	return result
}

func readConfig() []byte {
	data, err := ioutil.ReadFile("/config/deploy.yml")
	checkErr(err)
	return data
}

type Configuration struct {
	Goos         string
	Goarch       string
	Environments map[string]map[string]string `yaml:"environments"`
	Slack        map[string]string            `yaml:"slack"`
}
