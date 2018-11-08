package store

import (
	"sync"

	"github.com/aperturerobotics/entitygraph/entity"
)

// Handler handles store events.
type Handler interface {
	// HandleEntityAdded handles a new entity being added to the store.
	HandleEntityAdded(ent entity.Entity)
	// HandleEntityRemoved handles a entity being removed from the store.
	HandleEntityRemoved(ent entity.Entity)
}

// Store tracks, refcounts, and de-duplicates known entities.
type Store struct {
	// mtx guards the store
	mtx sync.Mutex
	// entities is the entities map, keyed by ID
	entities map[EntityMapKey]*storeEntity
	// handler is the store handler
	handler Handler
}

// NewStore constructs a new Store.
func NewStore(handler Handler) *Store {
	return &Store{
		entities: make(map[EntityMapKey]*storeEntity),
		handler:  handler,
	}
}

// AddEntityObj adds an entity object to the store.
func (s *Store) AddEntityObj(ent entity.Entity) {
	key := NewEntityMapKey(ent)
	s.mtx.Lock()
	defer s.mtx.Unlock()

	sent := s.entities[key]
	if sent == nil {
		sent = newStoreEntity(ent)
		s.entities[key] = sent
		s.emitStoreEntityAdd(sent)
		return
	}

	sent.PushEntity(ent)
}

// RemoveEntityObj removes an entity object from the store.
// Returns if found.
func (s *Store) RemoveEntityObj(ent entity.Entity) bool {
	key := NewEntityMapKey(ent)
	s.mtx.Lock()
	defer s.mtx.Unlock()

	sent := s.entities[key]
	if sent == nil {
		return false
	}

	f := sent.DelEntity(ent)
	if len(sent.Values) == 0 {
		delete(s.entities, key)
		s.emitStoreEntityRm(sent)
	}
	return f
}

// emitStoreEntityAdd emits a new store entity when added.
// mtx is locked
func (s *Store) emitStoreEntityAdd(storeEnt *storeEntity) {
	if s.handler != nil {
		s.handler.HandleEntityAdded(storeEnt.Ent)
	}
}

// emitStoreEntityRm emits a old store entity when removed.
// mtx is locked
func (s *Store) emitStoreEntityRm(storeEnt *storeEntity) {
	if s.handler != nil {
		s.handler.HandleEntityRemoved(storeEnt.Ent)
	}
}
