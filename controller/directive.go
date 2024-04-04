package entitygraph_controller

import (
	"github.com/aperturerobotics/controllerbus/directive"
	"github.com/aperturerobotics/entitygraph"
)

// CollectEntityGraphDirective implements CollectEntityGraph.
type CollectEntityGraphDirective struct{}

// NewCollectEntityGraphDirective constructs a new CollectEntityGraph directive.
func NewCollectEntityGraphDirective() *CollectEntityGraphDirective {
	return &CollectEntityGraphDirective{}
}

// CollectEntityGraphDirective is a marker function.
func (d *CollectEntityGraphDirective) CollectEntityGraphDirective() {
	// noop
}

// Validate validates the directive.
// This is a cursory validation to see if the values "look correct."
func (d *CollectEntityGraphDirective) Validate() error {
	return nil
}

// GetValueOptions returns options relating to value handling.
func (d *CollectEntityGraphDirective) GetValueOptions() directive.ValueOptions {
	return directive.ValueOptions{}
}

// IsEquivalent checks if the other directive is equivalent. If two
// directives are equivalent, and the new directive does not superceed the
// old, then the new directive will be merged (de-duplicated) into the old.
func (d *CollectEntityGraphDirective) IsEquivalent(other directive.Directive) bool {
	_, ok := other.(entitygraph.CollectEntityGraph)
	return ok
}

// Superceeds checks if the directive overrides another.
// The other directive will be canceled if superceded.
func (d *CollectEntityGraphDirective) Superceeds(other directive.Directive) bool {
	return false
}

// GetName returns the directive's type name.
// This is not necessarily unique, and is primarily intended for display.
func (d *CollectEntityGraphDirective) GetName() string {
	return "CollectEntityGraph"
}

// GetDebugVals returns the directive arguments as k/v pairs.
// This is not necessarily unique, and is primarily intended for display.
func (d *CollectEntityGraphDirective) GetDebugVals() directive.DebugValues {
	return nil
}

// _ is a type constraint
var _ entitygraph.CollectEntityGraph = ((*CollectEntityGraphDirective)(nil))
