package app

import (
	"os"

	"github.com/lexcao/watch-log/internal/loader/file"
	"github.com/lexcao/watch-log/internal/parser/json"
	"github.com/lexcao/watch-log/internal/pipeline"
	"github.com/lexcao/watch-log/internal/pipeline/match"
	"github.com/lexcao/watch-log/pkg/component"
	"github.com/lexcao/watch-log/pkg/model"
)

// Component for watch log, Flow:
// Loader -> Parser -> [Pipeline] -> Renderer <-> Controller
type Component struct {
	controller component.Controller
	loader     component.Loader
	parser     component.Parser
	pipeline   component.Pipeline
	render     component.Renderer
}

func Run(render component.Renderer) {
	loader := file.LiveTailLoader{}
	parser := json.Parser{}
	pipe := pipeline.NewProcessPipeline()

	//pipe.AddPipe(field.OmitPipeline("ts"))
	//pipe.AddPipe(field.PickPipeline("language"))
	pipe.AddPipe(match.Pipeline{Match: model.Object{
		"language": "Java",
	}})

	line := loader.Load(os.Stdin)

	for line.HasNext() {
		origin := line.Next()
		entry := model.NewEntry(origin)

		parser.Parse(entry)

		pipe.Pipe(entry)

		render.Render(entry)
	}
}
