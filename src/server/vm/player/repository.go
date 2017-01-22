package player

import "github.com/domain-query-language/dql-server/src/server/vm"

type Repository interface {

	Get(id vm.Identifier) (Player, error)

	Save(player Player) error
}
