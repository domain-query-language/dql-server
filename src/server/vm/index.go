package vm

import "errors"

var (
	INDEX_ENTITY_NOT_EXISTS = errors.New("Entity does not exist.")
	INDEX_ENTITY_ALREAD_EXISTS = errors.New("Entity already exists in index.")
)

type Index interface {

	Add(entity Entity) error

	Replace(entity Entity) error

	Remove(identifier Identifier) error
}
