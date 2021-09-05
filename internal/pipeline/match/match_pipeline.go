package match

import (
	"reflect"

	"github.com/lexcao/watch-log/pkg/model"
)

type Pipeline struct {
	Match model.Object
}

func (p Pipeline) Pipe(entry *model.Entry) {
	if len(p.Match) == 0 {
		return
	}

	object := entry.PipelinedObject
	matched := true
	for key, value := range p.Match {
		toMatch, exist := object[key]
		if !exist || !reflect.DeepEqual(toMatch, value) {
			matched = false
			break
		}
	}
	if len(p.Match) > 0 && !matched {
		entry.PipelinedObject = make(model.Object)
	}
}

func (p Pipeline) Order() int {
	return 100
}
