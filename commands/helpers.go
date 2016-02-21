package commands

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"
)

func getDir() string {
	wd, _ := os.Getwd()
	dir := strings.Split(wd, "/")
	return fmt.Sprint(dir[len(dir)-1])
}

// make it config function
func slackEnabled(slack map[string]string) bool {
	if len(slack) == 0 {
		return false
	} else {
		return true
	}
}

func checkErrWithMsg(e error, slackConfig map[string]string) {
	if e != nil {
		errorMsg(slackConfig)
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
