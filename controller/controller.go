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

	"github.com/blang/semver"
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
	observers map[*entityObserver]func(val entity.Entity)
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
		observers: make(map[*entityObserver]func(val entity.Entity)),
	}
	c.store = store.NewStore(newStoreHandler(c))
	return c
}

// GetControllerInfo returns information about the controller.
func (c *Controller) GetControllerInfo() controller.Info {
	return controller.Info{
		Id:      ControllerID,
		Version: Version.String(),
	}
}

// HandleDirective asks if the handler can resolve the directive.
// If it can, it returns a resolver. If not, returns nil.
// Any exceptional errors are returned for logging.
// It is safe to add a reference to the directive during this call.
// The context passed is canceled when the directive instance expires.
func (c *Controller) HandleDirective(
	ctx context.Context,
	inst directive.Instance,
) (directive.Resolver, error) {
	dir := inst.GetDirective()
	if d, ok := dir.(entitygraph.CollectEntityGraph); ok {
		return c.resolveCollectEntityGraph(ctx, inst, d)
	}

	return nil, nil
}

// resolveCollectEntityGraph resolves a CollectEntityGraph directive
func (c *Controller) resolveCollectEntityGraph(
	ctx context.Context,
	inst directive.Instance,
	d entitygraph.CollectEntityGraph,
) (directive.Resolver, error) {
	return newEntityObserver(c), nil
}

// registerObserver registers an observer and returns the initial set
func (c *Controller) registerObserver(obs *entityObserver, cb func(val entity.Entity)) []entity.Entity {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.observers[obs] = cb
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
		newReferenceHandler(c, store),
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
