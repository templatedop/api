package port
import (
	"github.com/templatedop/api/module"
)
func Module() *module.Module {
	m := module.New("error")
	m.Provide()
	return m
}
