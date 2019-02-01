package entitygraph_logger

import (
	"github.com/aperturerobotics/controllerbus/bus"
	"github.com/aperturerobotics/controllerbus/directive"
	"github.com/aperturerobotics/entitygraph"
	"github.com/aperturerobotics/entitygraph/entity"
	"github.com/sirupsen/logrus"
)

// AttachBasicLogger attaches a basic logging directive.
func AttachBasicLogger(
	b bus.Bus,
	le *logrus.Entry,
) (directive.Reference, error) {
	_, diRef, err := b.AddDirective(
		entitygraph.NewObserveEntityGraph(),
		bus.NewCallbackHandler(func(val directive.AttachedValue) {
			ent := val.GetValue().(entity.Entity)
			le.Infof("value added: %s: %s", ent.GetEntityTypeName(), ent.GetEntityID())
		}, func(val directive.AttachedValue) {
			ent := val.GetValue().(entity.Entity)
			le.Infof("value removed: %s: %s", ent.GetEntityTypeName(), ent.GetEntityID())
		}, nil),
	)
	return diRef, err
}
