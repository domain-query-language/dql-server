package adapter

import (
	"testing"
	"github.com/domain-query-language/dql-server/src/server/adapter"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
)

type FakeCommand struct {}

func (c *FakeCommand) Id() vm.Identifier {
	return nil
}

func (c *FakeCommand) TypeId() vm.Identifier {
	return nil
}

func (c *FakeCommand) AggregateId() *vm.AggregateIdentifier {
	return nil
}

func (c *FakeCommand) Payload() vm.Payload {
	return nil
}

func TestTypesAreCorrectForCommand(t *testing.T) {

	cmd := &FakeCommand{};
	hndlr := adapter.NewCommand( cmd );

	if (hndlr.Typ != adapter.CMD) {
		t.Error("Expected command, got query");
	}

	if (hndlr.Query != nil) {
		t.Error("Query should be nil")
	}

	if (hndlr.Command == nil) {
		t.Error("Command should not be nil")
	}

	if (hndlr.Command != cmd) {
		t.Error("Commands should be equal");
	}
}

type FakeQuery struct {}

func (c *FakeQuery) String() string {
	return "";
}

func (c *FakeQuery) Id() vm.Identifier {
	return nil;
}

func TestTypesAreCorrectForQuery(t *testing.T) {

	qry := &FakeQuery{};
	hndlr := adapter.NewQuery( qry );

	if (hndlr.Typ != adapter.QRY) {
		t.Error("Expected query, got command");
	}

	if (hndlr.Command != nil) {
		t.Error("Command should be nil")
	}

	if (hndlr.Query == nil) {
		t.Error("Query should not be nil")
	}

	if (hndlr.Query != qry) {
		t.Error("Queries should be equal");
	}
}