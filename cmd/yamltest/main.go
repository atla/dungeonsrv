package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

var data = `
- templateID: sword_001
  name: Simple Sword
  description: This is a very basic sword
  properties:
   strength: 3
   attackSpeed: 2
`

// correctly populate the data.
type T struct {
	TemplateID  string `yaml:"templateID"`
	Name        string
	Description string
	Properties  map[string]string `yaml:",flow"`
}

func main() {
	// if we use struct containing yaml encoding for yaml formated string
	t := []T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t after unmarshal:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t after marshal:\n%s\n\n", string(d))
}
