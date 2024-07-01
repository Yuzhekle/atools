package main

import (
	"encoding/json"
	"strings"
)

func main() {
	str := "dbr: not found"
	if !strings.Contains(str, "not found") {
		println("not found")
	} else {
		println("found")
	}

	str1 := []struct {
		Str string `json:"str"`
	}{
		{"dbr: not found"},
		{"new line"},
	}
	data, err := json.Marshal(str1)
	if err != nil {
		panic(err)
	}
	println(string(data))
}
