package main

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
