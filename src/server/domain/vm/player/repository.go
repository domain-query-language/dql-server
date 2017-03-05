package player

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"errors"
)

var (
	NOT_EXISTS = errors.New("The player does not exist.")
)

type Repository interface {

	Get(id vm.Identifier) (*Player, error)

	GetByContext(id vm.Identifier) ([]*Player, error)

	Save(player *Player) error
}
