package entitygraph

import (
	"github.com/aperturerobotics/controllerbus/directive"
)

// observeEntityGraph implements ObserveEntityGraph
type observeEntityGraph struct{}

// NewObserveEntityGraph constructs a new ObserveEntityGraph directive.
func NewObserveEntityGraph() ObserveEntityGraph {
	return &observeEntityGraph{}
}

// Validate validates the directive.
// This is a cursory validation to see if the values "look correct."
func (g *observeEntityGraph) Validate() error {
	return nil
}

// ObserveEntityGraphDirective is a marker function.
func (g *observeEntityGraph) ObserveEntityGraphDirective() {
	// noop
}

// GetValueOptions returns options relating to value handling.
func (g *observeEntityGraph) GetValueOptions() directive.ValueOptions {
	return directive.ValueOptions{}
}

// IsEquivalent checks if the other directive is equivalent. If two
// directives are equivalent, and the new directive does not superceed the
// old, then the new directive will be merged (de-duplicated) into the old.
func (g *observeEntityGraph) IsEquivalent(other directive.Directive) bool {
	_, ok := other.(ObserveEntityGraph)
	return ok
}

// Superceeds checks if the directive overrides another.
// The other directive will be canceled if superceded.
func (g *observeEntityGraph) Superceeds(other directive.Directive) bool {
	return false
}

// GetName returns the directive's type name.
// This is not necessarily unique, and is primarily intended for display.
func (d *observeEntityGraph) GetName() string {
	return "ObserveEntityGraph"
}

// GetDebugVals returns the directive arguments as k/v pairs.
// This is not necessarily unique, and is primarily intended for display.
func (d *observeEntityGraph) GetDebugVals() directive.DebugValues {
	return nil
}

// _ is a type assertion
var _ ObserveEntityGraph = ((*observeEntityGraph)(nil))
