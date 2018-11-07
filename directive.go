package entitygraph

import (
	"github.com/aperturerobotics/controllerbus/directive"
)

// CollectEntityGraph is a directive to collect entity graph entities.
type CollectEntityGraph interface {
	// Directive indicates CollectEntityGraph is a directive.
	directive.Directive

	// CollectEntityGraphDirective is a marker function.
	// noop
	CollectEntityGraphDirective()
}

// ObserveEntityGraph is a directive to observe the entity graph.
type ObserveEntityGraph interface {
	// Directive indicates CollectEntityGraph is a directive.
	directive.Directive

	// ObserveEntityGraphDirective is a marker function.
	// noop
	ObserveEntityGraphDirective()
}
