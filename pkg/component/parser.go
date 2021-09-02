package component

import "github.com/lexcao/watch-log/pkg/model"

type Parser interface {
	Parse(entry *model.Entry)
}
