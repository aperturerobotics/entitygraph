package entitygraph_controller

import (
	"github.com/aperturerobotics/entitygraph/entity"
	"github.com/aperturerobotics/entitygraph/store"
)

// storeHandler handles collect events.
type storeHandler struct {
	// c is the controller
	c *Controller
}

// newStoreHandler constructs a new storeHandler.
func newStoreHandler(c *Controller) *storeHandler {
	return &storeHandler{c: c}
}

// HandleEntityAdded handles a new entity being added to the store.
func (h *storeHandler) HandleEntityAdded(ent entity.Entity) {
	// TODO: emit to directive
}

// HandleEntityRemoved handles a entity being removed from the store.
func (h *storeHandler) HandleEntityRemoved(ent entity.Entity) {
	// TODO: emit to directive
}

// _ is a type assertion
var _ store.Handler = ((*storeHandler)(nil))
