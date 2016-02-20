package commands

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
