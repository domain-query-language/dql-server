package player

import (
	"github.com/domain-query-language/dql-server/src/server/vm"
	"github.com/domain-query-language/dql-server/src/server/vm/player"
	"github.com/boltdb/bolt"
)

type BoltRepository struct {

	path string
	db *bolt.DB
}

func (self *BoltRepository) Reset() error {

	error := self.db.Update(func(tx *bolt.Tx) error {

		tx.DeleteBucket([]byte("players"))
		tx.CreateBucket([]byte("players"))

		return nil
	})

	return error
}

func (self *BoltRepository) Open() error {

	bolt, err := bolt.Open(self.path, 0600, nil)

	self.db = bolt

	return err
}

func (self *BoltRepository) Close() error {
	return self.db.Close()
}

func (self *BoltRepository) Get(id vm.Identifier) (player.Player, error) {

	snapshot := player.Snapshot{}

	err := self.db.View(func(tx *bolt.Tx) error {

		players_bucket := tx.Bucket([]byte("players"))

		encoded_snapshot := players_bucket.Get(id.Bytes())

		snapshot.Decode(encoded_snapshot)

		return nil
	})

	return snapshot, err
}

func (self *BoltRepository) Save(snapshot player.Snapshot) error {

	err := self.db.Update(func(tx *bolt.Tx) error {

		players_bucket := tx.Bucket([]byte("players"))

		players_bucket.Put(snapshot.Id.Bytes(), snapshot.Encode())

		return nil
	})

	return err
}
