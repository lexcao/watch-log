package app

import (
	"github.com/lexcao/watch-log/pkg/component"
)

type Option = func(*App)

func Loader(l component.Loader) Option {
	return func(app *App) {
		app.loader = l
	}
}

func Parser(p component.Parser) Option {
	return func(app *App) {
		app.parser = p
	}
}

func Renderer(r component.Renderer) Option {
	return func(app *App) {
		app.render = r
	}
}

func Pipeline(p component.Pipeline) Option {
	return func(app *App) {
		app.pipeline = p
	}
}

func Controller(c component.Controller) Option {
	return func(app *App) {
		app.controller = c
	}
}
