package database

import (
	"context"

	"github.com/GoHyperrr/mdk"
)

type Module struct{}

func NewModule() mdk.Module {
	return &Module{}
}

func (m *Module) ID() string {
	return "core.database"
}

func (m *Module) Models() []any {
	return nil
}

func (m *Module) Routes() []mdk.Route {
	return nil
}

func (m *Module) Init(ctx context.Context, rt mdk.Runtime) error {
	return nil
}

func (m *Module) Shutdown(ctx context.Context) error {
	return nil
}

func init() {
	mdk.Register(func() mdk.Module {
		return NewModule()
	})
}
