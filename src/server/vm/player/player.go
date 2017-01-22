package player

import (
	"github.com/domain-query-language/dql-server/src/server/vm"
	"time"
)

type Player struct {

	id vm.Identifier

	stream Stream
	projector Projector
}

func (self *Player) Reset() {

	self.projector.Reset()
	self.stream.Reset()
}

func (self *Player) Load(snapshot Snapshot) {
	self.stream.Seek(snapshot.LastId)
}

func (self *Player) Play(limit int) error {

	play := true
	unlimited := false

	if(limit == 0) {
		unlimited = true
	}

	play_count := 0

	for play == true {

		event, stream_err := self.stream.Next()

		if stream_err != nil {
			return stream_err
		}

		proj_err := self.projector.Apply(
			event,
		)

		if proj_err != nil {
			return proj_err
		}

		if(!unlimited && play_count >= limit) {
			play = false
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

func CreatePlayer(id vm.Identifier, stream Stream, projector Projector) *Player {

	return &Player {
		id: id,
		stream: stream,
		projector: projector,
	}
}
