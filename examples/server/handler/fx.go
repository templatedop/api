package handler

import (
	"github.com/templatedop/api/module"
)

var Handlermodule = module.New("handler").Provide(

	NewUserHandler,
	NewService,
)
