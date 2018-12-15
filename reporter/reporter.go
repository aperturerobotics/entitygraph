package reporter

import (
	"context"

	"github.com/aperturerobotics/controllerbus/bus"
	"github.com/aperturerobotics/controllerbus/controller"
	"github.com/aperturerobotics/controllerbus/directive"
	"github.com/aperturerobotics/entitygraph/store"
	"github.com/sirupsen/logrus"
)

// Constructor constructs a reporter with common parameters.
type Constructor func(
	le *logrus.Entry,
	bus bus.Bus,
	store *store.Store,
) (Reporter, error)

// Reporter creates and handles directives, exposing entities to the graph.
type Reporter interface {
	// Handler handles directives.
	directive.Handler

	// Execute executes the reporter's inner controller.
	// Entities should be added or removed from the store.
	Execute(ctx context.Context) error
}

// Controller is a reporter controller.
type Controller interface {
	// Controller is the controllerbus controller interface.
	controller.Controller

	// GetReporter returns the controlled reporter.
	// This may wait for the reporter to be ready.
	GetReporter(ctx context.Context) (Reporter, error)
}
