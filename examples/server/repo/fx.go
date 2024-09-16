package repository

import (
	"github.com/templatedop/api/module"
)

var Repomodule = module.New("repository").Provide(
	NewUserRepository,
)
