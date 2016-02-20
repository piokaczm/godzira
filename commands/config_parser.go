package commands

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

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
func getServers(config *Configuration, env string) (map[string]string, error) {
	// maybe store user@host already in the struct? separate user and host are not really used right now
	servers := make(map[string]string)
	i := 1
	for key, value := range config.Environments[env] {

		pattern := regexp.MustCompile("^(host)(_)?(\\d+)?$") // it's stupid why 3 groups, 2 should be enough, _ is mandatory for multiple hosts
		if pattern.MatchString(key) {
			// if key == 'host' or 'host_[digit]'
			digit := regexp.MustCompile("^\\d+$")
			match := pattern.FindStringSubmatch(key)

			no := strconv.Itoa(i)
			var host_number string
			server := []string{"server_", no}   // look down m8
			new_key := strings.Join(server, "") // its not necessery, those keys are not used anywhere, name it via key i guess

			// try to handle it smarter
			if len(match) >= 3 {
				host_number = match[3]
			} else {
				host_number = ""
			}

			if digit.MatchString(host_number) {
				// if more than one host
				user_number := []string{"user_", host_number}
				user := strings.Join(user_number, "")
				user = config.Environments[env][user]

				servers[new_key] = parseServer(user, value)
				i += 1
			} else {
				// if only one host
				user := config.Environments[env]["user"]
				servers[new_key] = parseServer(user, value)
				i += 1
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
