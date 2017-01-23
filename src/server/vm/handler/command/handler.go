package command

import (
	"github.com/domain-query-language/dql-server/src/server/vm/aggregate"
	"github.com/domain-query-language/dql-server/src/server/vm"
	"github.com/domain-query-language/dql-server/src/server/vm/player"
)

type Handler interface {

	Handle(command Command) ([]vm.Event, error)
}

/*
	SimpleHandler Implementation
 */

type SimpleHandler struct {

	context_map map[vm.Identifier][]vm.Identifier

	repository_aggregates aggregate.Repository
	repository_players player.Repository
}

func (self *SimpleHandler) Handle(command Command) ([]vm.Event, error) {

	// Get Aggregate from Repository
	agg, _ := self.repository_aggregates.Get(
		aggregate.CreateIdentifier(
			command.Id(),
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
		player, err := self.repository_players.Get(player_id)

		player.Play(1000)
	}

	return events, nil
}
