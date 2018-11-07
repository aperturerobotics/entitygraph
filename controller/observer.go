package entitygraph_controller

import (
	"context"
	"sync"

	"github.com/aperturerobotics/entitygraph/entity"

	"github.com/aperturerobotics/controllerbus/directive"
)

// entityObserver is an entity observer.
type entityObserver struct {
	// c is the controller
	c *Controller
}

// newEntityObserver constructs a new entityObserver
func newEntityObserver(c *Controller) *entityObserver {
	return &entityObserver{c: c}
}

// Resolve resolves the values, emitting them to the handler.
// The resolver may be canceled and restarted multiple times.
// Any fatal error resolving the value is returned.
// The resolver will not be retried after returning an error.
// Values will be maintained from the previous call.
func (e *entityObserver) Resolve(ctx context.Context, handler directive.ResolverHandler) error {
	// Register the observer with the system and get the initial set
	var valueIDS []uint32
	var valueMtx sync.Mutex
	var disposed bool
	addVal := func(ent entity.Entity) {
		if disposed {
			return
		}

		id, ok := handler.AddValue(ent)
		if ok {
			valueIDS = append(valueIDS, id)
		}
	}

	valueMtx.Lock()
	initialSet := e.c.registerObserver(e, func(ent entity.Entity) {
		valueMtx.Lock()
		defer valueMtx.Unlock()
		addVal(ent)
	})
	for _, is := range initialSet {
		addVal(is)
	}
	valueMtx.Unlock()

	defer func() {
		valueMtx.Lock()
		disposed = true
		vids := valueIDS
		valueIDS = nil
		valueMtx.Unlock()

		e.c.clearObserver(e)
		for _, valID := range vids {
			_, _ = handler.RemoveValue(valID)
		}
	}()

	<-ctx.Err()
	return nil
}

// _ is a type assertion
var _ directive.Resolver = ((*entityObserver)(nil))
