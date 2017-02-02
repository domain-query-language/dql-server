package main

import (
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"github.com/domain-query-language/dql-server/examples/e-commerce/domain/shopping/aggregate/cart/command"
	"github.com/satori/go.uuid"
	"github.com/domain-query-language/dql-server/examples/e-commerce/domain/shopping"
	"github.com/davecgh/go-spew/spew"
	"github.com/domain-query-language/dql-server/examples/e-commerce/application"
)

func main() {


	aggregate_id, _ := uuid.FromString("cf48952b-3c67-4339-9b4f-3846acd01d73")

	command := vm.NewCommand(
		command.TypeCreate,
		aggregate_id,
		shopping.ContextID,
		command.Create {
			ShopperId: uuid.NewV4(),
		},
	)

	spew.Dump(command)
	spew.Dump(application.CommandHandler)

}
