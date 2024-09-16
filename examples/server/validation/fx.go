package validation

import (
	"github.com/templatedop/api/module"
)
var ValidatorModule = module.New("cvalidator").Provide(
	NewHelloTextValidator,
	NewQueryTextValidator,
)