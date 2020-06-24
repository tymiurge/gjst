package gjst

import (
	"fmt"
	"os"
)

type scope struct {
	static map[string]string
}

func newScope() *scope {
	return &scope{
		static: make(map[string]string),
	}
}

func (s *scope) append(key, value string) {
	if _, exists := s.static[key]; exists {
		fmt.Printf("duplicated key %s in tempalte variables scope\n", key)
		os.Exit(1)
	}
	s.static[key] = value
}
