package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	testJSON := `
    {
      "a": "thing",
      "b": "another thing"
    }
  `

	var testParsed map[string]string

	json.Unmarshal([]byte(testJSON), &testParsed)
	fmt.Println("Parsed [a]: ", testParsed["a"])
}
