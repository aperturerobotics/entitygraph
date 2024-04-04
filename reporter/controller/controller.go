package reporter_controller

import (
	"context"
	"sync"

	"github.com/aperturerobotics/entitygraph"
	"github.com/aperturerobotics/entitygraph/entity"
	"github.com/aperturerobotics/entitygraph/reporter"
	"github.com/aperturerobotics/entitygraph/store"

	"github.com/aperturerobotics/controllerbus/bus"
	"github.com/aperturerobotics/controllerbus/controller"
	"github.com/aperturerobotics/controllerbus/directive"

	"github.com/sirupsen/logrus"
)

// Controller manages a Entity Graph reporter.
// It handles CollectEntityGraph directives.
type Controller struct {
	// le is the root logger
	le *logrus.Entry
	// store is the entity store
	store *store.Store
	// bus is the controller bus
	bus bus.Bus
	// ctor is the constructor
	ctor reporter.Constructor
	// reporterCh contains the controlled reporter
	reporterCh chan reporter.Reporter
	// controllerInfo contains the controller info
	controllerInfo *controller.Info

	// mtx guards the collectors map
	mtx sync.Mutex
	// handlers should be called with values from store
	handlers []store.Handler
	// values are the known values
	values map[store.EntityMapKey]entity.Entity
}

// NewController constructs a new entitygraph reporter controller.
func NewController(
	le *logrus.Entry,
	bus bus.Bus,
	info *controller.Info,
	ctor reporter.Constructor,
) *Controller {
	c := &Controller{
		le:             le,
		bus:            bus,
		ctor:           ctor,
		controllerInfo: info,
		values:         make(map[store.EntityMapKey]entity.Entity),
		reporterCh:     make(chan reporter.Reporter, 1),
	}
	c.store = store.NewStore(newStoreHandler(c))
	return c
}

// Execute executes the given controller.
// Returning nil ends execution.
// Returning an error triggers a retry with backoff.
func (c *Controller) Execute(ctx context.Context) error {
	// Construct the reporter.
	r, err := c.ctor(c.le, c.bus, c.store)
	if err != nil {
		return err
	}

	c.reporterCh <- r

	// Wait for the controller to quit.
	return r.Execute(ctx)
}

// HandleDirective asks if the handler can resolve the directive.
// If it can, it returns a resolver. If not, returns nil.
// Any unexpected errors are returned for logging.
// It is safe to add a reference to the directive during this call.
func (c *Controller) HandleDirective(
	ctx context.Context,
	di directive.Instance,
) ([]directive.Resolver, error) {
	dir := di.GetDirective()
	switch d := dir.(type) {
	case entitygraph.CollectEntityGraph:
		return directive.R(c.handleCollectEntityGraph(ctx, di, d))
	}

	return nil, nil
}

// handleCollectEntityGraph handles a CollectEntityGraph directive.
func (c *Controller) handleCollectEntityGraph(
	ctx context.Context,
	di directive.Instance,
	d entitygraph.CollectEntityGraph,
) (directive.Resolver, error) {
	return newCollectResolver(ctx, c), nil
}

// GetControllerInfo returns information about the controller.
func (c *Controller) GetControllerInfo() *controller.Info {
	return c.controllerInfo
}

// GetReporter returns the controlled reporter.
// This may wait for the reporter to be ready.
func (c *Controller) GetReporter(ctx context.Context) (reporter.Reporter, error) {
	var r reporter.Reporter
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case r = <-c.reporterCh:
		c.reporterCh <- r
	}

	return r, nil
}

// Close releases any resources used by the controller.
// Error indicates any issue encountered releasing.
func (c *Controller) Close() error {
	return nil
}

// _ is a type assertion
var _ reporter.Controller = ((*Controller)(nil))
