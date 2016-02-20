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

func slackEnabled(slack map[string]string) bool {
	if len(slack) == 0 {
		return false
	} else {
		return true
	}
}

func printErr(e error) {
	fmt.Println(e)
	os.Exit(1)
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
