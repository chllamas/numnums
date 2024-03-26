package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type enumArr []string

type NumNumInput struct {
	target string
	enums  string
}

func main() {
	filename := "test.json"
	fileData, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", filename, err.Error())
		return
	}

	err = json.Unmarshal(fileData, nil)
}
