package command

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/aggregate"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/player"
	"github.com/davecgh/go-spew/spew"
)

type Handler interface {

	Handle(command vm.Command) ([]vm.Event, error)
}

/*
	SimpleHandler Implementation
 */

type Handler_ struct {

	repository_aggregates aggregate.Repository
	repository_players player.Repository
}

func (self *Handler_) Handle(command vm.Command) ([]vm.Event, error) {

	// Get Aggregate from Repository
	agg, _ := self.repository_aggregates.Get(command.AggregateId())

	spew.Dump(agg)

	// Handle Command
	events, handling_error := agg.Handle(command)

	if handling_error != nil {
		return nil, handling_error
	}

	self.repository_aggregates.Save(agg)

	spew.Dump(self.repository_aggregates)

	/*
	players_index := self.context_map[command.AggregateTypeId()]

	for player_id := range players_index {
		player, _ := self.repository_players.Get(player_id)

		player.Play(1000)
	}

	*/

	return events, nil
}

func NewHandler(
	repository_aggregates aggregate.Repository,
	repository_players player.Repository,
) *Handler_ {

	return &Handler_{
		repository_aggregates: repository_aggregates,
		repository_players: repository_players,
	}
}
