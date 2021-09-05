package converter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringConverter(t *testing.T) {
	// Given
	converter := StringConverter()

	// When
	expect, ok := converter("test", "Should be string")

	// Then
	assert.True(t, ok)
	assert.Equal(t, "Should be string", expect)
	assert.IsType(t, "string", expect)
}
