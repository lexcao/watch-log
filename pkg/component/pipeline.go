package component

import "github.com/lexcao/watch-log/pkg/model"

type Pipeline interface {
	Pipe(entry *model.Entry)
	Order() int
}
