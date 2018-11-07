package entity

// Ref is a reference to another entity.
type Ref interface {
	// GetEntityRefId returns the referenced entity ID.
	GetEntityRefId() string
	// GetEntityRefTypeName returns the referenced entity type name.
	GetEntityRefTypeName() string
}

// entityRef implements Ref
type entityRef struct {
	refID       string
	refTypeName string
}

// NewEntityRef builds a new EntityRef with an entity handle.
func NewEntityRef(ent Entity) Ref {
	return NewEntityRefWithID(ent.GetEntityID(), ent.GetEntityTypeName())
}

// NewEntityRefWithID builds a new EntityRef with an id and type name.
func NewEntityRefWithID(refID, refTypeName string) Ref {
	return &entityRef{refID: refID, refTypeName: refTypeName}
}

// GetEntityRefId returns the referenced entity ID.
func (e *entityRef) GetEntityRefId() string {
	return e.refID
}

// GetEntityRefTypeName returns the referenced entity type name.
func (e *entityRef) GetEntityRefTypeName() string {
	return e.refTypeName
}

// _ is a type assertion
var _ Ref = ((*entityRef)(nil))
