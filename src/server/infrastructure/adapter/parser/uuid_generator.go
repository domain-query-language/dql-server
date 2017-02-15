package parser

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/satori/go.uuid"
)

type UuidGenerator struct {

}

func (u *UuidGenerator) Generate() vm.Identifier {

	return uuid.NewV4()
}

func NewUuidGenerator() *UuidGenerator {

	return &UuidGenerator{}
}
