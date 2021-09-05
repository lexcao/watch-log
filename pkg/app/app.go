package app

import (
	"os"

	"github.com/lexcao/watch-log/internal/common/loader/file"
	"github.com/lexcao/watch-log/internal/common/parser/json"
	"github.com/lexcao/watch-log/internal/common/pipeline"
	"github.com/lexcao/watch-log/pkg/component"
	"github.com/lexcao/watch-log/pkg/model"
)

// App for watch log, Flow:
// Loader -> Parser -> [Pipeline] -> Renderer <-> Controller
type App struct {
	controller component.Controller
	loader     component.Loader
	parser     component.Parser
	pipeline   component.Pipeline
	render     component.Renderer
}

func New(opts ...Option) *App {
	app := &App{}

	app.Apply(defaultOptions()...)
	app.Apply(opts...)

	return app
}

func (app *App) Apply(opts ...Option) {
	for _, option := range opts {
		option(app)
	}
}

func (app *App) Run() {
	line := app.loader.Load(os.Stdin)

	for line.HasNext() {
		origin := line.Next()
		entry := model.NewEntry(origin)

		app.parser.Parse(entry)

		app.pipeline.Pipe(entry)

		app.render.Render(entry)
	}
}

func defaultOptions() []Option {
	return []Option{
		Loader(file.LiveTailLoader{}),
		Parser(json.Parser{}),
		Pipeline(pipeline.NewProcessPipeline()),
	}
}
