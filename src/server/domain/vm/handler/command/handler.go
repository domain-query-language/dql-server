package command

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm/aggregate"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/src/server/domain/vm/player"
)

type Handler interface {

	Handle(command vm.Command) ([]vm.Event, error)
}

/*
	SimpleHandler Implementation
 */

type Handler_ struct {

	context_map map[vm.Identifier][]vm.Identifier

	repository_aggregates aggregate.Repository
	repository_players player.Repository
}

func (self *Handler_) Handle(command vm.Command) ([]vm.Event, error) {

	// Get Aggregate from Repository
	agg, _ := self.repository_aggregates.Get(
		aggregate.CreateIdentifier(
			command.AggregateId(),
			command.AggregateTypeId(),
		),
	)

	// Handle Command
	events, handling_error := agg.Handle(command)

	if(handling_error) {
		return nil, handling_error
	}

	// Save Aggregate to Repository
	self.repository_aggregates.Save(agg)

	/*
		Play Domain Projectors
	 */
	players_index := self.context_map[command.AggregateTypeId()]

	for player_id := range players_index {
		player, _ := self.repository_players.Get(player_id)

		player.Play(1000)
	}

	return events, nil
}

func NewHandler(
	context_map map[vm.Identifier]vm.Identifier,
	repository_aggregates aggregate.Repository,
	repository_players player.Repository,
) *Handler_ {

	return &Handler_{
		context_map: context_map,
		repository_aggregates: repository_aggregates,
		repository_players: repository_players,
	}
}
