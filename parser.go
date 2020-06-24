package gjst

import (
	"encoding/json"
	"fmt"
	"os"
)

func Parse(b []byte) {
	var template map[string]interface{}
	err := json.Unmarshal([]byte(b), &template)
	if err != nil {
		fmt.Println("error while parsing template file.")
		os.Exit(1)
	}
	if _, exists := template["data"]; !exists {
		fmt.Println("no data key discovered in template.")
		os.Exit(1)
	}
	s := newScope()
	for k := range template {
		if k != "data" {
			s.append(k, fmt.Sprintf("%v", template[k]))
		}

	}
	fmt.Println("ok")
}