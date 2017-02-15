package player

/*
import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/player"
	"github.com/boltdb/bolt"
)

type BoltRepositoryCached struct {

	cache map[vm.Identifier]*player.Player

	repository BoltRepository
}

func (self *BoltRepositoryCached) Reset() error {
	return self.repository.Reset()
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

			snapshot := &player.Snapshot{}

			snapshot.Decode(v)

			self.cache[k] = player.CreatePlayer(
				snapshot.Id,

			)

			return nil
		})
	})
}

func (self *BoltRepositoryCached) Close() error {
	return self.repository.Close()
}

func (self *BoltRepositoryCached) Get(id vm.Identifier) (player.Player, error) {
	return self.cache[id], nil
}

func (self *BoltRepositoryCached) Save(snapshot player.Snapshot) error {

	self.cache[snapshot.Id] = snapshot

	return self.repository.Save(snapshot)
}
*/