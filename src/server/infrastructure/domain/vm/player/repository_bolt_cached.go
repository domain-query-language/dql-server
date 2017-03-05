package player

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/player"
	"github.com/boltdb/bolt"
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/src/server/domain/store"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
	"github.com/davecgh/go-spew/spew"
)

type BoltRepositoryCached struct {

	cache map[vm.Identifier]*player.Player
	context_cache map[vm.Identifier][]*player.Player

	repository *BoltRepository
}

func (self *BoltRepositoryCached) Reset() error {

	self.cache = map[vm.Identifier]*player.Player{}
	self.context_cache = map[vm.Identifier][]*player.Player{}

	return self.repository.Reset()
}

func (self *BoltRepositoryCached) Get(id vm.Identifier) (*player.Player, error) {

	spew.Dump(self.cache[id].Snapshot())

	return self.cache[id], nil
}

func (self *BoltRepositoryCached) GetByContext(id vm.Identifier) ([]*player.Player, error) {



	return self.context_cache[id], nil
}

func (self *BoltRepositoryCached) Save(player *player.Player) error {

	// If it doesn't exist in cache, add it.
	if _, exists := self.cache[player.Id()]; !exists {
		self.cache[player.Id()] = player
		self.context_cache[player.ContextId()] = append(self.context_cache[player.ContextId()], player)
	}

	return self.repository.Save(player)
}

func (self *BoltRepositoryCached) Add(identifier vm.Identifier, projector projection.Projector) {
	self.repository.Add(identifier, projector)
}

func (self *BoltRepositoryCached) Open() error {

	err := self.repository.Open()

	if err != nil {
		return err
	}

	// Priming Cache
	return self.repository.db.View(func(tx *bolt.Tx) error {

		players_bucket := tx.Bucket([]byte("players"))

		return players_bucket.ForEach(func(k, v []byte) error {

			p, _ := self.repository.Get(
				vm.Identifier(uuid.FromBytesOrNil(k)),
			)

			self.cache[p.Id()] = p

			return nil
		})
	})
}

func (self *BoltRepositoryCached) Close() error {
	return self.repository.Close()
}

func CreateBoltRespositoryCached(path string, log store.Log) *BoltRepositoryCached {

	return &BoltRepositoryCached {
		cache: map[vm.Identifier]*player.Player{},
		context_cache: map[vm.Identifier][]*player.Player{},

		repository: CreateBoltRepository(path, log),
	}
}
