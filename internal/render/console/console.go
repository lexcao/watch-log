package console

import (
	"fmt"
	"sort"

	"github.com/lexcao/watch-log/pkg/model"
)

const (
	DEFAULT = "\033[0m"
	RED     = "\033[0;31m"
)

type Renderer struct {
}

func (c Renderer) Render(entry *model.Entry) {
	output := entry.PipelinedObject

	var keys []string
	for k := range output {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	render("========================================")
	if entry.Err != nil {
		render(entry.Err, RED)
	} else {
		for _, k := range keys {
			fmt.Printf("[%s]: %s\n", k, output[k])
		}
	}
	render("========================================")
}

func render(value interface{}, color ...string) {
	if len(color) == 0 {
		fmt.Println(value)
	} else {
		fmt.Println(color[0], value, DEFAULT)
	}
}
