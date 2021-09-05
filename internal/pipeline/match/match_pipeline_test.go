package match

import (
	"testing"

	"github.com/lexcao/watch-log/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestPipeline_Pipe(t *testing.T) {
	given := make(model.Object)
	given["a"] = 1
	given["b"] = 2
	given["c"] = 3

	tests := []struct {
		name   string
		given  model.Object
		when   model.Object
		expect bool
	}{
		{
			"none match",
			given,
			model.Object{},
			true,
		},
		{
			"match b=2",
			given,
			model.Object{
				"b": 2,
			},
			true,
		},
		{
			"not match d=4",
			given,
			model.Object{
				"d": 4,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			entry := &model.Entry{
				PipelinedObject: given,
			}

			// When
			p := Pipeline{Match: tt.when}
			p.Pipe(entry)

			// Then
			if tt.expect {
				assert.Equal(t, tt.given, entry.PipelinedObject)
			} else {
				assert.Equal(t, model.Object{}, entry.PipelinedObject)
			}
		})
	}
}
