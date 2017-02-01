package tokenizer

import (
	tok "github.com/domain-query-language/dql-server/src/server/adapter/parser/token"
)

type Tokenizer interface {
	Tokens() []tok.Token
	Next() *tok.Token
}

type tokeniser struct {
	l *lexer
	index int
}

func (t *tokeniser) Tokens () []tok.Token {

	return t.l.tokens;
}

func (t *tokeniser) Next() (*tok.Token) {

	if (t.index >= len(t.l.tokens)) {
		return nil
	}
	token := t.l.tokens[t.index];
	t.index++;

	return &token;
}

func NewTokenizer(dql string) Tokenizer {

	l := lex("DQL", dql);

	return &tokeniser{l, 0};
}


