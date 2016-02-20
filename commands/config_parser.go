package commands

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// decode yaml data and set Configuration struct fields using it
// set user@server for choosen environment for further deploy and ssh commands
func parseConfig(data []byte) Configuration {
	result := Configuration{}
	err := yaml.Unmarshal([]byte(data), &result)
	checkErr(err)

	return result
}

// { server_1: cos@cos.net, server_2: cos2@cos2.net }
func setServers(config *Configuration, env string) map[string]string {
	servers := make(map[string]string)
	i := 1
	for key, value := range config.Environments[env] {

		pattern := regexp.MustCompile("^(host)(_)?(\\d+)?$") // we need to get int as well for user matching
		digit := regexp.MustCompile("^\\d+$")
		server_number := pattern.FindStringSubmatch(key)[2]

		if pattern.MatchString(key) && digit.MatchString(server_number) {
			no := strconv.Itoa(i)

			server := []string{"server_", no}
			new_key := strings.Join(server, "")

			user_number := []string{"user_", server_number}
			user := strings.Join(user_number, "")
			user = config.Environments[key][user]

			servers[new_key] = parseServer(user, value)
			i += 1
		}
	}
	return servers
}

func parseServer(user string, host string) string {
	server := []string{user, "@", host}
	return strings.Join(server, "")
}

// read user configuration from config/deploy.yml
func readConfig() []byte {
	data, err := ioutil.ReadFile("/config/deploy.yml")
	checkErr(err)
	return data
}

type Configuration struct {
	Goos         string                       `yaml:"goos"`
	Goarch       string                       `yaml:"goarch"`
	Environments map[string]map[string]string `yaml:"environments"`
	Slack        map[string]string            `yaml:"slack"`
	Test         bool                         `yaml:"test"`
	Godep        bool                         `yaml:"godep"`
}
