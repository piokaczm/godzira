package commands

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"testing"
)

func runCommand(name string, args []string, start_msg string, finish_msg string) {
	fmt.Println(start_msg)

	err := exec.Command(name, args...).Run()

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	} else {
		fmt.Println(finish_msg)
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
