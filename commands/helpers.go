package commands

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

const (
	deployed = `
                   ,:',:',:'
              __||_||_||_||__
         ____["""""""""""""""]____
         \ " '''''''''''''''''''' \
  ~^~^~^~^~^^~^~^~^~^~^~^~^~~^~^~^^~~^~^
     _            _                      _ 
  __| | ___ _ __ | | ___  _   _  ___  __| |
 / _  |/ _ \ '_ \| |/ _ \| | | |/ _ \/ _  |
| (_| |  __/ |_) | | (_) | |_| |  __/ (_| |
 \__,_|\___| .__/|_|\___/ \__, |\___|\__,_|
           |_|            |___/            
  `
)

func getDir() string {
	wd, _ := os.Getwd()
	dir := strings.Split(wd, "/")
	return fmt.Sprint(dir[len(dir)-1])
}

// maybe make it config func -> see config_parser.go
func slackEnabled(slack map[string]string) bool {
	if len(slack) == 0 {
		return false
	} else {
		return true
	}
}

func checkErrWithMsg(e error, slackConfig map[string]string) {
	if e != nil {
		if slackEnabled(slackConfig) {
			errorMsg(slackConfig)
		}
		panic(e)
	}
}

func checkCommandError(e error, slackConfig map[string]string, output []byte) {
	if e != nil {
		if slackEnabled(slackConfig) {
			errorMsg(slackConfig)
		}
		boldRed := color.New(color.FgRed, color.Bold)
		boldRed.Println(string(output))
		panic(e)
	}
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func blank(str string) bool {
	if len(str) == 0 {
		return true
	}
	return false
}

func notBlank(str string) bool {
	if len(str) == 0 {
		return false
	}
	return true
}

func deployPrint(server string, str string) {
	msg := strings.Join([]string{server, str}, ": ")
	color.Cyan("\n%s", msg)
}
