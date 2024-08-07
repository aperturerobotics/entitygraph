package entitygraph_controller

import (
	"context"
	"sync"

	"github.com/aperturerobotics/controllerbus/bus"
	"github.com/aperturerobotics/controllerbus/controller"
	"github.com/aperturerobotics/controllerbus/directive"

	"github.com/aperturerobotics/entitygraph"
	"github.com/aperturerobotics/entitygraph/entity"
	"github.com/aperturerobotics/entitygraph/store"

	"github.com/blang/semver/v4"
	"github.com/sirupsen/logrus"
)

// Version is the version of the controller implementation.
var Version = semver.MustParse("0.0.1")

// ControllerID is the ID of the controller.
const ControllerID = "entitygraph/collector/1"

// Controller implements the entity graph aggregation controller.
// The controller issues a CollectEntityGraph directive, aggregates entities,
// and fulfills ObserveEntityGraph directives.
type Controller struct {
	// le is the log entry
	le *logrus.Entry
	// bus is the controller bus
	bus bus.Bus
	// conf is the configuration
	conf *Config
	// store is the store
	store *store.Store

	// mtx guards entities
	mtx sync.Mutex
	// entities is the map of known entities
	// we have to store it here so we can emit when the directive is first made
	entities map[store.EntityMapKey]entity.Entity
	// observer is the map of entity observers
	observers map[*entityObserver]*registeredObserver
}

// NewController constructs a new entity graph controller.
func NewController(
	le *logrus.Entry,
	bus bus.Bus,
	conf *Config,
) *Controller {
	c := &Controller{
		le:        le,
		bus:       bus,
		conf:      conf,
		entities:  make(map[store.EntityMapKey]entity.Entity),
		observers: make(map[*entityObserver]*registeredObserver),
	}
	c.store = store.NewStore(newStoreHandler(c))
	return c
}

// GetControllerInfo returns information about the controller.
func (c *Controller) GetControllerInfo() *controller.Info {
	return controller.NewInfo(
		ControllerID,
		Version,
		"entity-graph observer",
	)
}

// HandleDirective asks if the handler can resolve the directive.
func (c *Controller) HandleDirective(
	ctx context.Context,
	inst directive.Instance,
) ([]directive.Resolver, error) {
	dir := inst.GetDirective()
	if d, ok := dir.(entitygraph.ObserveEntityGraph); ok {
		return directive.R(c.resolveObserveEntityGraph(ctx, inst, d))
	}

	return nil, nil
}

// resolveObserveEntityGraph resolves a ObserveEntityGraph directive
func (c *Controller) resolveObserveEntityGraph(
	ctx context.Context,
	inst directive.Instance,
	d entitygraph.ObserveEntityGraph,
) (directive.Resolver, error) {
	return newEntityObserver(c), nil
}

// registerObserver registers an observer and returns the initial set
func (c *Controller) registerObserver(robs *registeredObserver) []entity.Entity {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.observers[robs.entityObs] = robs
	initialSet := make([]entity.Entity, len(c.entities))
	i := 0
	for _, ent := range c.entities {
		initialSet[i] = ent
		i++
	}
	return initialSet
}

// clearObserver removes an observer
func (c *Controller) clearObserver(obs *entityObserver) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	delete(c.observers, obs)
}

// Execute executes the given controller.
// Returning nil ends execution.
// Returning an error triggers a retry with backoff.
func (c *Controller) Execute(ctx context.Context) error {
	// Register collect entity graph directive
	di, diRef, err := c.bus.AddDirective(
		NewCollectEntityGraphDirective(),
		newReferenceHandler(c, c.store),
	)
	if err != nil {
		return err
	}
	defer diRef.Release()

	_ = di
	c.le.Info("entitygraph aggregation controller running")
	<-ctx.Done()
	return nil
}

// Close releases any resources used by the controller.
// Error indicates any issue encountered releasing.
func (c *Controller) Close() error {
	return nil
}

// _ is a type assertion
var _ controller.Controller = ((*Controller)(nil))
