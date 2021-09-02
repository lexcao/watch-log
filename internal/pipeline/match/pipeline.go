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
	found := false
	for key, value := range p.Match {
		if object != nil && reflect.DeepEqual(object[key], value) {
			found = true
			break
		}
	}
	if len(p.Match) > 0 && !found {
		return
	}
	entry.PipelinedObject = make(model.Object)
}

func (p Pipeline) Order() int {
	return 100
}
