package field

import (
	"github.com/lexcao/watch-log/internal/pipeline/field/converter"
	"github.com/lexcao/watch-log/internal/pipeline/field/matcher"
	"github.com/lexcao/watch-log/pkg/model"
)

type Pipeline struct {
	order     int
	Predicate matcher.Matcher
}

func (p Pipeline) Pipe(entry *model.Entry) {
	object := make(model.Object)
	for key, value := range entry.PipelinedObject {
		if newValue, ok := p.Predicate(key, value); ok {
			object[key] = newValue
		}
	}
	entry.PipelinedObject = object
}

func (p Pipeline) Order() int {
	return p.order
}

func StringPipeline() Pipeline {
	return Pipeline{
		order:     10,
		Predicate: converter.StringConverter(),
	}
}

func OmitPipeline(field string) Pipeline {
	return Pipeline{
		order:     20,
		Predicate: matcher.OmitMatcher(field),
	}
}

func PickPipeline(field string) Pipeline {
	return Pipeline{
		order:     30,
		Predicate: matcher.PickMatcher(field),
	}
}
