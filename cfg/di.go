package cfg

import (
	"github.com/sarulabs/di"

	"github.com/ameteiko/golang-kit/errors"
)

//
// DIContainer is a dependency resolver object.
//
type DIContainer struct {
	ctx     di.Context
	builder *di.Builder

	ignoredDeps map[string]bool
}

//
// Custom registration helpers.
//
type dependencyRegistrar func(ctx di.Context) (interface{}, error)
type dependencyDisposer func(obj interface{})

//
// NewDIContainer returns an instance of the DI DIContainer.
//
func NewDIContainer() *DIContainer {
	c := DIContainer{}
	c.builder, _ = di.NewBuilder(di.App, di.Request)
	c.ignoredDeps = make(map[string]bool)

	return &c
}

//
// Build builds application dependencies.
//
func (c *DIContainer) Build() error {
	c.ctx = c.builder.Build()

	return nil
}

//
// Get returns the dependency by its name.
//
func (c *DIContainer) Get(name string) interface{} {

	return c.ctx.Get(name)
}

//
// SetDependency sets the dependency by its name.
//
func (c *DIContainer) SetDependency(name string, obj interface{}) error {
	c.ignoredDeps[name] = true

	return c.builder.Set(name, obj)
}

//
// IsDependencySet returns true if dependency was set already.
//
func (c *DIContainer) IsDependencySet(name string) bool {

	return c.ignoredDeps[name]
}

//
// RegisterDependency registers a dependency in DI.
//
func (c *DIContainer) RegisterDependency(
	depName string,
	registrar dependencyRegistrar,
	disposer dependencyDisposer,
) error {

	err := c.builder.AddDefinition(
		di.Definition{
			Name:  depName,
			Build: registrar,
			Close: disposer,
		})
	if nil != err {
		return errors.WithMessage(err, `kit@cfg.DIContainer:RegisterDependency [error on dependency (%s) registration]`, depName)
	}

	return nil
}
