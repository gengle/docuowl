//+build ignore

package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"strings"
)

type Emoji struct {
	Symbol  string   `json:"emoji"`
	Aliases []string `json:"aliases"`
}

func main() {
	data, err := os.ReadFile("../../static/emoji.json")
	if err != nil {
		panic(err)
	}

	var emojis []Emoji
	if err = json.Unmarshal(data, &emojis); err != nil {
		panic(err)
	}

	file := []string{
		"package html",
		"",
		"// Code generated by go:generate. DO NOT EDIT.",
		"",
		"var emojiAliases = map[string]string{",
	}

	for _, e := range emojis {
		for _, a := range e.Aliases {
			file = append(file, fmt.Sprintf(`"%s": "%s",`, a, e.Symbol))
		}
	}

	file = append(file, "}", "")

	formatted, err := format.Source([]byte(strings.Join(file, "\n")))
	if err != nil {
		panic(err)
	}

	if err = os.WriteFile("./emoji_list.go", formatted, os.ModePerm); err != nil {
		panic(err)
	}
}