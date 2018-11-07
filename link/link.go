package link

import "github.com/aperturerobotics/entitygraph/entity"

// Link is an edge/relationship between two nodes.
type Link interface {
	// Entity indicates that Link is a subclass of Entity.
	entity.Entity

	// GetEdgeFrom returns the reference to the entity this link starts at.
	GetEdgeFrom() entity.Ref
	// GetEdgeTo returns the reference to the entity this link ends at.
	GetEdgeTo() entity.Ref
}
