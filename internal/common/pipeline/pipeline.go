package pipeline

import (
	"sort"

	"github.com/lexcao/watch-log/internal/common/pipeline/field"
	"github.com/lexcao/watch-log/pkg/component"
	"github.com/lexcao/watch-log/pkg/model"
)

type ProcessPipeline struct {
	pipes map[int]component.Pipeline
}

func NewProcessPipeline() ProcessPipeline {
	var process = ProcessPipeline{
		pipes: make(map[int]component.Pipeline),
	}

	process.AddPipe(field.StringPipeline())

	return process
}

func (p ProcessPipeline) Pipe(entry *model.Entry) {
	entry.PipelinedObject = entry.ParsedObject
	if entry.Err != nil {
		return
	}

	var sortedByOrder []int
	for k := range p.pipes {
		sortedByOrder = append(sortedByOrder, k)
	}
	sort.Ints(sortedByOrder)

	for _, order := range sortedByOrder {
		p.pipes[order].Pipe(entry)
	}
}

func (p ProcessPipeline) Order() int {
	return 0
}

func (p ProcessPipeline) AddPipe(pipeline component.Pipeline) {
	p.pipes[pipeline.Order()] = pipeline
}
