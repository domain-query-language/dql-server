package parser

import (
	"github.com/domain-query-language/dql-server/src/server/adapter/parser/tokenizer"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser/token"
	"fmt"
	"errors"
)

type tokenStream struct {

	t tokenizer.Tokenizer
	error error

	curToken  *token.Token
	peekToken *token.Token
}

func NewTokenStreamFromFreshTokenizer(t tokenizer.Tokenizer) *tokenStream {

	tokenStream := &tokenStream{}
	tokenStream.t = t

	tokenStream.NextToken()
	tokenStream.NextToken()

	return tokenStream
}

func NewTokenStreamFromExistingTokenizer(t tokenizer.Tokenizer, curToken *token.Token, peekToken *token.Token) *tokenStream {

	tokenStream := &tokenStream{}
	tokenStream.t = t

	tokenStream.curToken = curToken
	tokenStream.peekToken = peekToken

	return tokenStream
}

func (p *tokenStream) CurToken() *token.Token {

	return p.curToken
}

func (p *tokenStream) PeekToken() *token.Token {

	return p.peekToken
}

func (p *tokenStream) NextToken() {

	p.curToken = p.peekToken
	p.peekToken = p.t.Next();
}

func (p *tokenStream) CurTokenIs(t token.TokenType) bool {

	return p.curToken.Type == t
}

func (p *tokenStream) PeekTokenIs(t token.TokenType) bool {

	if p.peekToken == nil {
		return false
	}

	return p.peekToken.Type == t
}

func (p *tokenStream) ExpectPeek(t token.TokenType) bool {

	if p.PeekTokenIs(t) {

		p.NextToken()
		return true
	} else {

		p.PeekError(t)
		return false
	}
}

func (p *tokenStream) ExpectCurrent(t token.TokenType) bool {

	if p.CurTokenIs(t) {

		p.NextToken()
		return true
	} else {

		p.PeekError(t)
		return false
	}
}

func (p *tokenStream) PeekError(t token.TokenType) {

	if (p.peekToken == nil) {
		p.LogError("Expected next token to be '%s', got EOF instead", t)
		return;
	}
	p.LogError("Expected next token to be '%s', got '%s' instead", t, p.peekToken.Val)
}

func (p *tokenStream) LogError(format string, a...interface{}) {

	msg := fmt.Sprintf(format, a...)
	p.error = errors.New(msg)
}

func (p *tokenStream) Error() error {

	return p.error
}
