package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/cli/go-gh/v2"
)

func main() {
	var repos []string
	if file, err := os.ReadFile("./input/repos"); err != nil {
		panic(err)
	} else {
		ls := strings.Split(string(file), "\n")
		for _, l := range ls {
			if l != "" {
				repos = append(repos, l)
			}
		}
	}

	var prefixes []string
	if file, err := os.ReadFile("./input/prefixes"); err != nil {
		panic(err)
	} else {
		ls := strings.Split(string(file), "\n")
		for _, l := range ls {
			if l != "" {
				prefixes = append(prefixes, l)
			}
		}
	}

	type repo struct {
		Name string `json:"name"`
	}
	for _, prefix := range prefixes {
		list, _, err := gh.Exec([]string{"repo", "list", prefix, "--json", "name"}...)
		if err != nil {
			panic(err)
		}
		var rs []repo
		if err := json.Unmarshal(list.Bytes(), &rs); err != nil {
			panic(err)
		}
		for _, r := range rs {
			repos = append(repos, fmt.Sprintf("%s/%s", prefix, r.Name))
		}
	}

	// clear output dir
	if err := os.RemoveAll("./output"); err != nil {
		panic(err)
	}
	// clone repos
	for _, r := range repos {
		fmt.Printf("cloning %s...\n", r)
		_, _, err := gh.Exec([]string{"repo", "clone", r, fmt.Sprintf("./output/%s", r)}...)
		if err != nil {
			panic(err)
		}
	}

	fmt.Println("all repos cloned")
}
