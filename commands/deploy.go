package commands

import (
	"github.com/codegangsta/cli"
	"ioutil"
)

func Deploy(c *cli.Context) {
}

func parseConfig() Configuration {
	result := Configuration{}

	data, err := ioutil.ReadFile("/config/deploy.yml")
	checkErr(err)

	err = yaml.Unmarshall([]byte(data), &result)
	checkErr(err)

	return result
}

type Configuration struct {
	Goos         string
	Goarch       string
	Environments map[string]map[string]string
	Slack        map[string]string
}
