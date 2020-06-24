package main

import (
	//"encoding/json"
	"fmt"
	"gjst"
	"io/ioutil"
	"os"
)

func main()  {
	file, err := os.Open("data/simple_case.json")
	if err != nil {
		fmt.Println("error while opening file")
		os.Exit(1)
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("error while reading file")
		os.Exit(1)
	}
	gjst.Parse(bytes)
}
