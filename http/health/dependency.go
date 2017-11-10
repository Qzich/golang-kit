package health

import (
	"time"
)

//
// DependencyResolver is the interface for application dependencies resolving.
//
type DependencyResolver interface {
	Check() (float64, error)
	GetName() string
}

//
// Dependency is an external dependency object.
//
type Dependency struct {
	name    string
	latency float32
	checker DependencyChecker
}

//
// Check checks is dependency alive and returns its latency.
//
func (d *Dependency) Check() (float64, error) {
	start := time.Now()
	err := d.checker()
	latency := time.Now().Sub(start).Seconds()

	return latency, err
}

//
// GetName returns dependency name.
//
func (d *Dependency) GetName() string {

	return d.name
}
