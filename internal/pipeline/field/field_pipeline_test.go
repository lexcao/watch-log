package field

import (
	"testing"

	"github.com/lexcao/watch-log/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestPipeline_Pipe(t *testing.T) {
	// Given
	mockMatcher := func(key string, value interface{}) (interface{}, bool) {
		return "Called", true
	}
	fieldPipeline := Pipeline{Predicate: mockMatcher}
	entry := new(model.Entry)
	entry.PipelinedObject = make(model.Object)
	entry.PipelinedObject["mock"] = "Should returns Called"

	// When
	fieldPipeline.Pipe(entry)

	// Then
	assert.Equal(t, "Called", entry.PipelinedObject["mock"])
}
