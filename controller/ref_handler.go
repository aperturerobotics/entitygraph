package entitygraph_controller

import (
	"github.com/aperturerobotics/controllerbus/directive"
	"github.com/aperturerobotics/entitygraph/entity"
	"github.com/aperturerobotics/entitygraph/store"
)

// referenceHandler handles collect events.
type referenceHandler struct {
	// c is the controller
	c *Controller
	// store is the store
	store *store.Store
}

// newReferenceHandler constructs a new referenceHandler.
func newReferenceHandler(c *Controller, store *store.Store) *referenceHandler {
	return &referenceHandler{c: c, store: store}
}

// HandleValueAdded is called when a value is added to the directive.
func (r *referenceHandler) HandleValueAdded(
	inst directive.Instance,
	val directive.Value,
) {
	valEnt, valEntOk := val.(entity.Entity)
	if !valEntOk {
		r.c.le.Warn("ignoring non-entity directive value added")
		return
	}

	r.store.AddEntityObj(valEnt)
}

// HandleValueRemoved is called when a value is removed from the directive.
func (r *referenceHandler) HandleValueRemoved(
	inst directive.Instance,
	val directive.Value,
) {
	valEnt, valEntOk := val.(entity.Entity)
	if !valEntOk {
		r.c.le.Warn("ignoring non-entity directive value removed")
		return
	}

	r.store.RemoveEntityObj(valEnt)
}

// HandleInstanceDisposed is called when a directive instance is disposed.
// This will occur if Close() is called on the directive instance.
func (r *referenceHandler) HandleInstanceDisposed(inst directive.Instance) {
	// noop
}

var _ directive.ReferenceHandler = ((*referenceHandler)(nil))
