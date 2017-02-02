package player

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/player"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

type MemoryRepository struct {

	cache map[[]byte]*player.Player
}

func (self *MemoryRepository) Reset() error {

	self.cache = map[[]byte]*player.Player{}

	return nil
}

func (self *MemoryRepository) Get(id vm.Identifier) (*player.Player, error) {
	return self.cache[id.Bytes()]
}

func (self *MemoryRepository) Save(player *player.Player) error {

	self.cache[player.] = player

	return nil
}

func CreateMemoryRepository() *MemoryRepository {

	return &MemoryRepository {
		cache: map[[]byte]*player.Player{},
	}
}
