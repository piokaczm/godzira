package commands

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"regexp"
	"strings"
)

// func (config *Configuration) slackEnabled(bool) {
// 	if len(config.Slack) == 0 {
// 		return false
// 	} else {
// 		return true
// 	}
// }

func getConfig() Configuration {
	return parseConfig(readConfig())
}

// decode yaml data and set Configuration struct fields using it
// set user@server for choosen environment for further deploy and ssh commands
func parseConfig(data []byte) Configuration {
	result := Configuration{}
	err := yaml.Unmarshal([]byte(data), &result)
	checkErr(err)

	return result
}

// create map of servers to deploy to
// { server_1: cos@cos.net, server_2: cos2@cos2.net }
func getServers(environments map[string]map[string]string, env string) ([]string, error) {
	// maybe store user@host already in the struct? separate user and host are not really used right now
	servers := []string{}
	for key, value := range environments[env] {

		pattern := regexp.MustCompile("^(host)(_\\d+)?$")
		if pattern.MatchString(key) {
			// if key == 'host' or 'host_[digit]'
			digit := regexp.MustCompile("\\d+")
			match := digit.FindStringSubmatch(key)
			multiple_hosts := len(match) != 0

			if multiple_hosts {
				// if more than one host
				host_number := match[0]
				user_number := []string{"user_", host_number}
				user := strings.Join(user_number, "")
				user = environments[env][user]

				servers = append(servers, parseServer(user, value))
			} else {
				// if only one host
				user := environments[env]["user"]
				servers = append(servers, parseServer(user, value))
			}
		}
	}

	// if no proper key found return error
	if len(servers) == 0 {
		return nil, errors.New("no proper host in config file!")
	} else {
		return servers, nil
	}
}

func parseServer(user string, host string) string {
	server := []string{user, "@", host}
	return strings.Join(server, "")
}

// read user configuration from config/deploy.yml
func readConfig() []byte {
	data, err := ioutil.ReadFile("./config/deploy.yml")
	checkErr(err)
	return data
}

type Configuration struct {
	AppName      string                       `yaml:"appname"`
	Goos         string                       `yaml:"goos"`
	Goarch       string                       `yaml:"goarch"`
	Environments map[string]map[string]string `yaml:"environments"`
	Slack        map[string]string            `yaml:"slack"`
	Test         bool                         `yaml:"test"`
	Godep        bool                         `yaml:"godep"`
}
