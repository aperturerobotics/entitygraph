package store

import "github.com/aperturerobotics/entitygraph/entity"

// EntityMapKey is the tuple key for the entities map
type EntityMapKey struct {
	id       string
	typeName string
}

// NewEntityMapKey constructs a new EntityMapKey from an Entity
func NewEntityMapKey(ent entity.Entity) EntityMapKey {
	return EntityMapKey{
		id:       ent.GetEntityID(),
		typeName: ent.GetEntityTypeName(),
	}
}

// storeEntity is an entity in the store.
type storeEntity struct {
	Ent    entity.Entity
	Key    EntityMapKey
	Values []storeEntityValue
}

// newStoreEntity constructs a new storeEntity from an Entity.
func newStoreEntity(ent entity.Entity) *storeEntity {
	return &storeEntity{
		Ent: ent,
		Key: NewEntityMapKey(ent),
		Values: []storeEntityValue{
			newStoreEntityValue(ent),
		},
	}
}

// PushEntity pushes an entity to the store value.
func (s *storeEntity) PushEntity(ent entity.Entity) {
	val := newStoreEntityValue(ent)
	s.Values = append(s.Values, val)
}

// DelEntity removes an entity from the store value.
// Returns if found.
func (s *storeEntity) DelEntity(ent entity.Entity) bool {
	for i := range s.Values {
		if s.Values[i].Entity == ent {
			s.Values[i] = s.Values[len(s.Values)-1]
			s.Values = s.Values[:len(s.Values)-1]
			return true
		}
	}

	return false
}

// storeEntityValue is a value representing an entity.
type storeEntityValue struct {
	// Entity is the entity object.
	Entity entity.Entity
}

// newStoreEntityValue constructs a new storeEntityValue
func newStoreEntityValue(ent entity.Entity) storeEntityValue {
	return storeEntityValue{Entity: ent}
}
