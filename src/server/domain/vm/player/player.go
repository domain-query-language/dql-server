package player

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"time"
	"github.com/domain-query-language/dql-server/src/server/domain/store"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/projection"
)

type Player struct {

	id vm.Identifier
	context_id vm.Identifier

	last_updated_at time.Time

	version int

	stream store.Stream
	projector projection.Projector
}

func (self *Player) Id() vm.Identifier {
	return self.id
}

func (self *Player) ContextId() vm.Identifier {
	return self.context_id
}

func (self *Player) Reset() {

	self.projector.Reset()
	self.stream.Reset()
}

func (self *Player) Load(snapshot Snapshot) {
	self.stream.Seek(snapshot.LastId)
}

func (self *Player) Play(limit int) error {

	unlimited := false

	if(limit == 0) {
		unlimited = true
	}

	play_count := 0

	for self.stream.Next() {

		event := self.stream.Value().(vm.Event)

		proj_err := self.projector.Apply(
			event,
		)

		self.version++
		self.last_updated_at = time.Now()

		if proj_err != nil {
			return proj_err
		}

		if(!unlimited && play_count >= limit) {
			return nil
		}
	}

	return nil
}

func (self *Player) Snapshot() *Snapshot {

	return &Snapshot {
		Id: self.id,
		ContextId: self.context_id,
		LastId: self.stream.LastId(),
		OccurredAt: self.last_updated_at,
		Version: self.version,
	}
}

func NewPlayer(id vm.Identifier, context_id vm.Identifier, stream store.Stream, projector projection.Projector) *Player {

	return &Player {
		id: id,
		context_id: context_id,
		last_updated_at: time.Now(),
		version: 0,
		stream: stream,
		projector: projector,
	}
}

func FromSnapshot(snapshot *Snapshot, stream store.Stream, projector projection.Projector) *Player {

	player := &Player {
		id: snapshot.Id,
		context_id: snapshot.ContextId,
		last_updated_at: snapshot.OccurredAt,
		version: snapshot.Version,
		stream: stream,
		projector: projector,
	}

	player.stream.Seek(snapshot.LastId)

	return player
}
