package swagger

import (
	"github.com/templatedop/api/module"
)

func Module() *module.Module {
	m := module.New("swagger")

	m.Provide(
		buildDocs,
		fiberWrapper,
	)

	return m
}
