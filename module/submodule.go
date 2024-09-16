package module

import "github.com/templatedop/api/as"


type SubModule struct {
	provides     []any
	provideHooks []ProvideHook
	assers       []as.Asser
}

func NewSubModule() *SubModule {
	return &SubModule{
		provides: make([]any, 0),
	}
}

func (sm *SubModule) AddProvideHook(hooks ...ProvideHook) *SubModule {
	sm.provideHooks = append(sm.provideHooks, hooks...)
	return sm
}

func (sm *SubModule) AddAsser(assers ...as.Asser) *SubModule {
	sm.assers = append(sm.assers, assers...)
	return sm
}

// Provide add constructor to provide pool
func (sm *SubModule) Provide(provide ...any) *SubModule {
	sm.provides = append(sm.provides, provide...)
	return sm
}
