package matcher

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOmitMatcher(t *testing.T) {
	type args struct {
		field string
	}
	tests := []struct {
		name  string
		args  args
		key   string
		value string
		want  Matcher
	}{
		{
			"none omit",
			args{},
			"TestKey",
			"TestValue",
			func(key string, value interface{}) (interface{}, bool) {
				return value, true
			},
		},
		{
			"omit key",
			args{"TestKey"},
			"TestKey",
			"TestValue",
			func(key string, value interface{}) (interface{}, bool) {
				return nil, false
			},
		},
		{
			"not a omit key",
			args{"TestKeyNot"},
			"TestKey",
			"TestValue",
			func(key string, value interface{}) (interface{}, bool) {
				return value, true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOK := OmitMatcher(tt.args.field)(tt.key, tt.value)
			want, wantOK := tt.want(tt.key, tt.value)
			assert.Equal(t, gotOK, wantOK)
			assert.Equal(t, got, want)
		})
	}
}

func TestPickMatcher(t *testing.T) {
	type args struct {
		field string
	}
	tests := []struct {
		name  string
		args  args
		key   string
		value string
		want  Matcher
	}{
		{
			"none pick",
			args{},
			"TestKey",
			"TestValue",
			func(key string, value interface{}) (interface{}, bool) {
				return value, true
			},
		},
		{
			"pick key",
			args{"TestKey"},
			"TestKey",
			"TestValue",
			func(key string, value interface{}) (interface{}, bool) {
				return value, true
			},
		},
		{
			"not a pick key",
			args{"TestKey"},
			"TestKeyNot",
			"TestValue",
			func(key string, value interface{}) (interface{}, bool) {
				return nil, false
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotOK := PickMatcher(tt.args.field)(tt.key, tt.value)
			want, wantOK := tt.want(tt.key, tt.value)
			assert.Equal(t, gotOK, wantOK)
			assert.Equal(t, got, want)
		})
	}
}
