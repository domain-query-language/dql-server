package player

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/player"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

type MemoryRepository struct {

	context_cache map[vm.Identifier][]*player.Player
	cache map[vm.Identifier]*player.Player
}

func (self *MemoryRepository) Add(player *player.Player) {

	self.context_cache[player.ContextId()] = append(self.context_cache[player.ContextId()], player)
	self.cache[player.Id()] = player
}

func (self *MemoryRepository) Reset() error {

	self.context_cache = map[vm.Identifier][]*player.Player{}
	self.cache = map[vm.Identifier]*player.Player{}

	return nil
}

func (self *MemoryRepository) Get(id vm.Identifier) (*player.Player, error) {
	return self.cache[id], nil
}

func (self *MemoryRepository) GetByContext(id vm.Identifier) ([]*player.Player, error) {
	return self.context_cache[id], nil
}

func (self *MemoryRepository) Save(player *player.Player) error {

	// Do nothing, cause, ya know, pointers are great.

	return nil
}

func CreateMemoryRepository() *MemoryRepository {

	return &MemoryRepository {
		context_cache: map[vm.Identifier][]*player.Player{},
		cache: map[vm.Identifier]*player.Player{},
	}
}
