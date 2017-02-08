package parser

import (
	"fmt"
	"github.com/domain-query-language/dql-server/src/server/adapter"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser/token"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser/tokenizer"
	query "github.com/domain-query-language/dql-server/src/server/query/schema"
	command "github.com/domain-query-language/dql-server/src/server/command/schema"
	"github.com/domain-query-language/dql-server/src/server/domain/vm"
	"strings"
)

type UnexpectedTokenError struct {
	Expected string
	Actual *token.Token
}

func (e *UnexpectedTokenError) Error() string {

	return fmt.Sprintf("Error at char %d, expected [%s], got [%s] instead", e.Actual.Pos, e.Expected, e.Actual.Type)
}


type UnexpectedIdentifierError struct {
	Expected string
	Actual *token.Token
}

func (e *UnexpectedIdentifierError) Error() string {

	return fmt.Sprintf("Error at char %d, expected '%s', got '%s' instead", e.Actual.Pos, e.Expected, e.Actual.Val)
}

/** Implementation of the adapter, written using the tokenizer, parser pattern */
type parser struct {

	t tokenizer.Tokenizer
	error error

	curToken  *token.Token
	peekToken *token.Token
}

func New(statements string) adapter.Adapter {

	t := tokenizer.NewTokenizer(statements);

	a:= &parser{t, nil, nil, nil}

	a.nextToken()
	a.nextToken()

	return a;
}

func (a *parser) nextToken() {

	a.curToken = a.peekToken
	a.peekToken = a.t.Next();
}

func (a *parser) curTokenIs(t token.TokenType) bool {

	return a.curToken.Type == t
}

func (a *parser) peekTokenIs(t token.TokenType) bool {

	if a.peekToken == nil {
		return false
	}

	return a.peekToken.Type == t
}

func (a *parser) expectPeek(t token.TokenType) bool {

 	if a.peekTokenIs(t) {

		a.nextToken()
		return true
	} else {

		a.peekError(t)
		return false
	}
}

func (a *parser) expectPeekIdentifier(value string) bool {

	if (!a.expectPeek(token.IDENT)) {
		return false
	}

	if (strings.ToLower(a.curToken.Val) != value) {
		a.currValueError(value)
		return false
	}

	return true
}

func (a *parser) peekError(t token.TokenType) {

	actual := a.peekToken
	if (a.peekToken == nil) {
		actual = &token.Token{token.EOF, "eof", a.curToken.Pos + len(a.curToken.Val)}
	}

	a.error = &UnexpectedTokenError{string(t), actual}
}

func (a *parser) currValueError(value string) {

	a.error = &UnexpectedIdentifierError{value, a.curToken}
}

// Return the next handlable object
func (a *parser) Next() (*adapter.Handleable, error) {

	if (a.curToken == nil) {

		return nil, nil;
	}

	if (a.isQuery()) {

		qry := a.parseQuery();
		if (qry == nil) {
			return nil, a.error
		}
		return adapter.NewQuery(qry), nil;
	}

	if (a.isCommand()) {

		cmd := a.parseCommand();

		if (cmd == nil) {
			return nil, a.error
		}

		return adapter.NewCommand(cmd), nil;
	}

	return nil, &UnexpectedTokenError{string(token.CREATE+"/"+token.LIST), a.curToken}
}

func (a *parser) isQuery() bool {

	return a.curTokenIs(token.LIST);
}

func (a *parser) isCommand() bool {

	return a.curTokenIs(token.CREATE);
}

func (a *parser) parseQuery() vm.Query {

	return a.parseListQuery();
}

func (a *parser) parseListQuery() vm.Query {

	qry := &query.ListDatabases{};

	if (!a.expectPeekIdentifier("databases")) {

		return nil;
	}

	if (!a.expectPeek(token.SEMICOLON)) {

		return nil;
	}

	return qry;
}

func (a *parser) parseCommand() vm.Command {

	return a.parseCreateCommand()
}

func (a *parser) parseCreateCommand() vm.Command {

	cmd := &command.CreateDatabase{};

	if (!a.expectPeekIdentifier("database")) {

		return nil;
	}

	if (!a.expectPeek(token.OBJECTNAME)) {

		return nil;
	}

	cmd.Name = a.curToken.Val

	if (!a.expectPeek(token.SEMICOLON)) {

		return nil;
	}

	return cmd;
}