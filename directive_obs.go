package entitygraph

import (
	"github.com/aperturerobotics/controllerbus/directive"
)

// observeEntityGraph implements ObserveEntityGraph
type observeEntityGraph struct{}

// ObserveEntityGraphDirective is a marker function.
func (g *observeEntityGraph) ObserveEntityGraphDirective() {
	// noop
}

// NewObserveEntityGraph constructs a new ObserveEntityGraph directive.
func NewObserveEntityGraph() ObserveEntityGraph {
	return &observeEntityGraph{}
}

// Validate validates the directive.
// This is a cursory validation to see if the values "look correct."
func (g *observeEntityGraph) Validate() error {
	return nil
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

// _ is a type assertion
var _ ObserveEntityGraph = ((*observeEntityGraph)(nil))
