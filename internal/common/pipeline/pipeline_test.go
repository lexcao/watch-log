package pipeline

import (
	"testing"

	"github.com/lexcao/watch-log/pkg/component"
	"github.com/lexcao/watch-log/pkg/model"
	"github.com/stretchr/testify/assert"
)

type mockPipeline struct {
	called string
	order  int
}

func (p mockPipeline) Pipe(entry *model.Entry) {
	entry.Origin = p.called
}

func (p mockPipeline) Order() int {
	return p.order
}

type replacedPipeline struct {
	order int
}

func (r replacedPipeline) Pipe(*model.Entry) {
}

func (r replacedPipeline) Order() int {
	return r.order
}

func TestProcessPipeline_Pipe(t *testing.T) {
	// Given
	called := "Called Mock"
	mock := mockPipeline{
		called: called,
		order:  1,
	}
	pipeline := NewProcessPipeline()
	pipeline.AddPipe(mock)
	entry := new(model.Entry)

	// When
	pipeline.Pipe(entry)

	// Then
	assert.Equal(t, called, entry.Origin)
}

func TestProcessPipeline_AddPipe(t *testing.T) {
	mock := mockPipeline{}

	t.Run("Case 1 added when absent", func(t *testing.T) {
		// Given
		pipeline := ProcessPipeline{pipes: map[int]component.Pipeline{}}
		assert.Equal(t, 0, len(pipeline.pipes))

		// When
		pipeline.AddPipe(mock)

		// Then
		assert.Equal(t, 1, len(pipeline.pipes))
	})

	t.Run("Case 2 replaced when existed with same order", func(t *testing.T) {
		// Given
		pipeline := ProcessPipeline{pipes: map[int]component.Pipeline{}}
		order := 100
		mock.order = order
		pipeline.AddPipe(mock)
		replaced := replacedPipeline{}
		replaced.order = order
		assert.Equal(t, 1, len(pipeline.pipes))

		// When
		pipeline.AddPipe(replaced)

		// Then
		assert.Equal(t, 1, len(pipeline.pipes))
		assert.IsType(t, replacedPipeline{}, pipeline.pipes[order])
	})
}
