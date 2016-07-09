package commands

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
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

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func Expect(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func DeepExpect(t *testing.T, a interface{}, b interface{}) {
	eq := reflect.DeepEqual(a, b)
	if eq != true {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}
