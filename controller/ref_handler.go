package entitygraph_controller

import (
	"github.com/aperturerobotics/controllerbus/directive"
)

// referenceHandler handles collect events.
type referenceHandler struct {
	// c is the controller
	c *Controller
}

// newReferenceHandler constructs a new referenceHandler.
func newReferenceHandler(c *Controller) *referenceHandler {
	return &referenceHandler{c: c}
}

// HandleValueAdded is called when a value is added to the directive.
func (r *referenceHandler) HandleValueAdded(
	inst directive.Instance,
	val directive.Value,
) {
}

// HandleValueRemoved is called when a value is removed from the directive.
func (r *referenceHandler) HandleValueRemoved(
	inst directive.Instance,
	val directive.Value,
) {
}

// HandleInstanceDisposed is called when a directive instance is disposed.
// This will occur if Close() is called on the directive instance.
func (r *referenceHandler) HandleInstanceDisposed(inst directive.Instance) {
}

var _ directive.ReferenceHandler = ((*referenceHandler)(nil))
