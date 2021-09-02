package converter

import (
	"fmt"

	"github.com/lexcao/watch-log/internal/pipeline/field/matcher"
)

func StringConverter() matcher.Matcher {
	return func(key string, value interface{}) (interface{}, bool) {
		return fmt.Sprint(value), true
	}
}
