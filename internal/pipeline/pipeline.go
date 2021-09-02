package pipeline

import (
	"sort"

	"github.com/lexcao/watch-log/internal/pipeline/field"
	_ "github.com/lexcao/watch-log/internal/pipeline/field"
	"github.com/lexcao/watch-log/internal/pipeline/match"
	_ "github.com/lexcao/watch-log/internal/pipeline/match"
	"github.com/lexcao/watch-log/pkg/component"
	"github.com/lexcao/watch-log/pkg/model"
)

var sortedByOrder []int

type ProcessPipeline struct {
	pipes map[int]component.Pipeline
}

func NewProcessPipeline() ProcessPipeline {
	var process = ProcessPipeline{
		pipes: make(map[int]component.Pipeline),
	}

	process.AddPipe(field.StringPipeline())
	process.AddPipe(field.OmitPipeline(""))
	process.AddPipe(field.PickPipeline(""))
	process.AddPipe(match.Pipeline{Match: make(model.Object)})

	for k := range process.pipes {
		sortedByOrder = append(sortedByOrder, k)
	}
	sort.Ints(sortedByOrder)

	return process
}

func (p ProcessPipeline) Pipe(entry *model.Entry) {
	entry.PipelinedObject = entry.ParsedObject
	if entry.Err != nil {
		return
	}

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
