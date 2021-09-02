package component

import "github.com/lexcao/watch-log/pkg/model"

type Renderer interface {
	Render(entry *model.Entry)
}
