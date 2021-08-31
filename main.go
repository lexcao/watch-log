package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	omit = ""
	pick = ""

	match = map[string]string{}
)

type Node = map[string]interface{}
type PipeFunc = func(Node) Node

const (
	DEFAULT = "\033[0m"
	RED     = "\033[0;31m"
)

func main() {
	// TODO removed by file
	if !onlyPipe() {
		return
	}

	// Loader
	scanner := bufio.NewScanner(bufio.NewReader(os.Stdin))

	for scanner.Scan() {
		// Parser
		line := scanner.Text()
		jsonMap := make(Node)
		err := json.Unmarshal([]byte(line), &jsonMap)
		if err != nil {
			render(err, RED)
			continue
		}

		// Pipeline
		pipe := func(object Node, pipes ...PipeFunc) Node {
			output := object
			for _, pipe := range pipes {
				output = pipe(output)
				if output == nil {
					return nil
				}
			}
			return output
		}

		pipeField := func(object Node) func(...func(string, interface{}) (interface{}, bool)) Node {
			return func(onField ...func(string, interface{}) (interface{}, bool)) Node {
				result := object
				for _, onFieldApply := range onField {
					newObject := make(Node)
					for key, value := range result {
						if newValue, ok := onFieldApply(key, value); ok {
							newObject[key] = newValue
						}
					}
					result = newObject
				}
				return result
			}
		}

		pipeString := func(key string, value interface{}) (interface{}, bool) {
			return fmt.Sprint(value), true
		}

		pipeOmit := func(key string, value interface{}) (interface{}, bool) {
			if len(omit) > 0 && strings.Contains(omit, key) {
				return nil, false
			} else {
				return value, true
			}
		}

		pipePick := func(key string, value interface{}) (interface{}, bool) {
			if len(pick) > 0 && !strings.Contains(pick, key) {
				return nil, false
			} else {
				return value, true
			}
		}

		pipeMatch := func(object Node) Node {
			if len(match) == 0 {
				return object
			}

			found := false
			for key, value := range match {
				if object[key] != nil && strings.Contains(object[key].(string), value) {
					found = true
					break
				}
			}
			if len(match) > 0 && !found {
				return nil
			}
			return object
		}

		output := pipe(pipeField(jsonMap)(pipeString), pipeMatch)
		if output == nil {
			continue
		}
		output = pipe(pipeField(output)(pipePick, pipeOmit))

		// Render
		var keys []string
		for k := range output {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		render("========================================")
		for _, k := range keys {
			fmt.Printf("[%s]: %s\n", k, output[k])
		}
		render("========================================")
	}
}
func render(value interface{}, color ...string) {
	if len(color) == 0 {
		fmt.Println(value)
	} else {
		fmt.Println(color[0], value, DEFAULT)
	}
}

func onlyPipe() bool {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		fmt.Println("The command is intended to work with pipes.")
		fmt.Println("Usage: output.log | wlog")
		return false
	}

	return true
}
