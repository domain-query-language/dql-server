package player

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/player"
	"github.com/boltdb/bolt"
	"github.com/davecgh/go-spew/spew"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
	"github.com/domain-query-language/dql-server/src/server/domain/store"
	"encoding/gob"
	"github.com/satori/go.uuid"
	"bytes"
)

type BoltRepository struct {

	log store.Log
	projectors map[vm.Identifier]projection.Projector

	path string
	db *bolt.DB

	encoding_buffer bytes.Buffer
	encoder *gob.Encoder

	decoding_buffer bytes.Buffer
	decoder *gob.Decoder
}

func (self *BoltRepository) Reset() error {

	error := self.db.Update(func(tx *bolt.Tx) error {

		tx.DeleteBucket([]byte("players"))
		tx.CreateBucket([]byte("players"))

		tx.DeleteBucket([]byte("contexts"))
		tx.CreateBucket([]byte("contexts"))

		return nil
	})

	return error
}

func (self *BoltRepository) Get(id vm.Identifier) (*player.Player, error) {

	var p *player.Player

	err := self.db.View(func(tx *bolt.Tx) error {

		players_bucket := tx.Bucket([]byte("players"))

		self.decoding_buffer.Reset()
		self.decoding_buffer.Write(
			players_bucket.Get(id.Bytes()),
		)

		snapshot := &player.Snapshot{}

		self.decoder.Decode(snapshot)

		spew.Dump(snapshot)

		p = player.FromSnapshot(
			snapshot,
			self.log.Stream(),
			self.projectors[snapshot.Id],
		)

		return nil
	})

	return p, err
}

func (self *BoltRepository) GetByContext(id vm.Identifier) ([]*player.Player, error) {

	players := []*player.Player {}

	err := self.db.View(func(tx *bolt.Tx) error {

		players_bucket := tx.Bucket([]byte("players"))
		contexts_bucket := tx.Bucket([]byte("contexts"))
		context_bucket := contexts_bucket.Bucket(id.Bytes())

		context_bucket.ForEach(func(key, v []byte) error {

			self.decoding_buffer.Reset()
			self.decoding_buffer.Write(
				players_bucket.Get(key),
			)

			snapshot := &player.Snapshot{}

			self.decoder.Decode(snapshot)

			spew.Dump(snapshot)

			players = append(
				players,
				player.FromSnapshot(
					snapshot,
					self.log.Stream(),
					self.projectors[snapshot.Id],
				),
			)

			return nil
		})

		return nil
	})

	return players, err
}

func (self *BoltRepository) Save(player *player.Player) error {

	spew.Dump(player.Snapshot())

	err := self.db.Update(func(tx *bolt.Tx) error {

		players_bucket := tx.Bucket([]byte("players"))
		contexts_bucket := tx.Bucket([]byte("contexts"))

		self.encoding_buffer.Reset()
		self.encoder.Encode(
			player.Snapshot(),
		)

		players_bucket.Put(
			player.Id().Bytes(),
			self.encoding_buffer.Bytes(),
		)

		context_bucket, _ := contexts_bucket.CreateBucketIfNotExists(
			player.ContextId().Bytes(),
		)

		context_bucket.Put(
			player.Id().Bytes(),
			[]byte("1"),
		)

		return nil
	})

	return err
}

func (self *BoltRepository) Add(identifier vm.Identifier, projector projection.Projector) {
	self.projectors[identifier] = projector
}

func (self *BoltRepository) Open() error {

	bolt, err := bolt.Open(self.path, 0600, nil)

	self.db = bolt

	return err
}

func (self *BoltRepository) Close() error {
	return self.db.Close()
}

func CreateBoltRepository(path string, log store.Log) *BoltRepository {

	gob.Register(vm.Identifier(uuid.UUID{}))

	repository := &BoltRepository {
		path: path,
		log: log,
		projectors: map[vm.Identifier]projection.Projector {},

	}

	repository.encoder = gob.NewEncoder(&repository.encoding_buffer)
	repository.decoder = gob.NewDecoder(&repository.decoding_buffer)

	return repository
}
