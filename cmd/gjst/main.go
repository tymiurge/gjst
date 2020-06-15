package main

import (
	//"encoding/json"
	"fmt"
	"gjst"
	"os"
)

func main()  {
	_, err := os.Open("data/data1.json")
	if err != nil {
		fmt.Println("error")
		os.Exit(1)
	}
	gjst.Parse()
}
