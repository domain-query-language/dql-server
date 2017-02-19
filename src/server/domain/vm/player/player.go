package player

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"time"
	"github.com/domain-query-language/dql-server/src/server/domain/store"
)

type Player struct {

	id vm.Identifier
	context_id vm.Identifier

	stream store.Stream
	projector Projector
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
		LastId: self.stream.LastId(),
		OccurredAt: time.Now(),
		Version: self.projector.Version(),
	}
}

func NewPlayer(id vm.Identifier, context_id vm.Identifier, stream store.Stream, projector Projector) *Player {

	return &Player {
		id: id,
		context_id: context_id,
		stream: stream,
		projector: projector,
	}
}
