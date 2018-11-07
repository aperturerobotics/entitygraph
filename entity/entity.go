package entity

// Entity is a string-ID identified node in the entity graph.
type Entity interface {
	// GetEntityID returns the entity identifier.
	GetEntityID() string
	// GetEntityTypeName returns the entity type name.
	GetEntityTypeName() string
}
