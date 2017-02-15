package player

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/player"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

type MemoryRepository struct {

	cache map[vm.Identifier][]*player.Player
}

func (self *MemoryRepository) Add(player *player.Player) {
	self.cache[player.ContextId()] = append(self.cache[player.ContextId()], player)
}

func (self *MemoryRepository) Reset() error {

	self.cache = map[vm.Identifier][]*player.Player{}

	return nil
}

func (self *MemoryRepository) Get(id vm.Identifier) ([]*player.Player, error) {
	return self.cache[id], nil
}

func (self *MemoryRepository) Save(player *player.Player) error {

	// Do nothing, cause, ya know, pointers are great.

	return nil
}

func CreateMemoryRepository() *MemoryRepository {

	return &MemoryRepository {
		cache: map[vm.Identifier][]*player.Player{},
	}
}
