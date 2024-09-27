package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/templatedop/api/modules/server/route"
)

type Handler interface {
	Routes() []route.Route
	Prefix() string
	Middlewares() []fiber.Handler
	Name() string
}

type Base struct {
	prefix string
	name   string
	mws    []fiber.Handler
}

func New(name string) *Base {
	return &Base{name: name}
}

func (b *Base) Name() string {
	return b.name
}

func (b *Base) Prefix() string {
	return b.prefix
}

func (b *Base) Middlewares() []fiber.Handler {
	return b.mws
}

func (b *Base) Routes() []route.Route {
	panic("need to declare routes for controller: " + b.name)
}
func (b *Base) AddPrefix(p string) *Base {
	b.prefix = b.prefix + p
	return b
}

func (b *Base) SetPrefix(p string) *Base {
	b.prefix = p
	return b
}

func (b *Base) AddMiddleware(mw fiber.Handler) *Base {
	b.mws = append(b.mws, mw)
	return b
}
