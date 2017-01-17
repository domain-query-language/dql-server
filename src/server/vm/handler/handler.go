package handler

import (
	"github.com/domain-query-language/dql-server/src/server/vm/aggregate"
	"github.com/domain-query-language/dql-server/src/server/vm"
)

type Handler interface {

	Handle(command vm.Command) ([]vm.Event, error)
}

type SimpleHandler struct {

	context_map map[vm.Identifier][]vm.Identifier

	players map[vm.Identifier]Player

	repository aggregate.Repository
}

func (self *SimpleHandler) Handle(command vm.Command) ([]vm.Event, error) {

	// Get Aggregate from Repository
	agg, _ := self.repository.Get(
		aggregate.CreateIdentifier(
			command.AggregateId(),
			command.AggregateTypeId(),
		),
	)

	// Handle Command
	events, error := agg.Handle(command)

	// Save Aggregate to Repository
	self.repository.Save(agg)

	/*
		Play Domain Projectors
	 */
	players_index := self.context_map[command.AggregateTypeId()]

	for player_id := range players_index {
		self.players[player_id].Play()
	}

	return events, error
}
