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
	h.c.mtx.Lock()
	defer h.c.mtx.Unlock()

	key := store.NewEntityMapKey(ent)
	h.c.entities[key] = ent
	for _, obs := range h.c.observers {
		if obs.addedCb != nil {
			obs.addedCb(ent)
		}
	}
}

// HandleEntityRemoved handles a entity being removed from the store.
func (h *storeHandler) HandleEntityRemoved(ent entity.Entity) {
	h.c.mtx.Lock()
	defer h.c.mtx.Unlock()

	key := store.NewEntityMapKey(ent)
	delete(h.c.entities, key)
	for _, obs := range h.c.observers {
		if obs.removedCb != nil {
			obs.removedCb(ent)
		}
	}
}

// _ is a type assertion
var _ store.Handler = ((*storeHandler)(nil))
