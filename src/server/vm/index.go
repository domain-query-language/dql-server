package vm

import "errors"

var (
	INDEX_ENTITY_NOT_EXISTS = errors.New("Entity does not exist.")
	INDEX_ENTITY_ALREADY_EXISTS = errors.New("Entity already exists in index.")
)

type Index interface {

	Add(entity Entity) error

	Replace(entity Entity) error

	Get(identifier Identifier) (Entity, error)

	Exists(identifier Identifier) bool

	Remove(identifier Identifier) error
}
