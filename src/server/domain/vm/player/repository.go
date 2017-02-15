package player

import "github.com/domain-query-language/dql-server/src/server/domain/vm"

type Repository interface {

	Get(id vm.Identifier) ([]*Player, error)

	Save(player *Player) error
}
