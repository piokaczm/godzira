package main

import (
	"gopkg.in/yaml.v2"
)

type Config struct {
	Mode         string                       `yaml:"mode"`
	Name         string                       `yaml:"binary_name"`
	Strategy     string                       `yaml:"strategy"`
	Goos         string                       `yaml:"goos"`
	Goarch       string                       `yaml:"goarch"`
	Environments map[string]map[string]string `yaml:"environments"`
	Slack        map[string]string            `yaml:"slack"`
	Test         bool                         `yaml:"test"`
	Vendor       bool                         `yaml:"vendor"`
}

type Environment struct {
	Hosts []string
	User  string
	Name  string
}

func (c *Config) load() (*Slack, []*Task) {
	// load yml
	data, err := ioutil.ReadFile("./config/deploy.yml") // maybe use abs here
	checkErr(err)
	err = yaml.Unmarshal([]byte(data), c)
	checkErr(err)

	// prepare slack config
	slack := Slack{}

	// parse envs
	// prepare tasks
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
