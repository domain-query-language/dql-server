package parser

import (
	"github.com/domain-query-language/dql-server/src/server/adapter/parser/tokenizer"
	"github.com/domain-query-language/dql-server/src/server/adapter/parser/token"
	"github.com/domain-query-language/dql-server/src/server/adapter/ast"
)

/** Implementation of the adapter, written using the tokenizer, parser pattern */
type objectComponentParser struct {

	tokenStream *tokenStream
}

func NewObjectComponent(function string) *objectComponentParser {

	t := tokenizer.NewTokenizer(function);

	tknStrm := NewTokenStreamFromFreshTokenizer(t)

	p := &objectComponentParser{tknStrm};

	return p;
}

func (p *objectComponentParser) ParseObjectComponent() (ast.ObjectComponent, error) {

	switch p.tokenStream.CurToken().Type {

	case token.FUNCTION:
		return p.parseFunction()

	case token.CHECK:
		return p.parseCheck()

	case token.HANDLER:
		return p.parseHandler()

	case token.WHEN:
		return p.parseWhen()

	case token.PROPERTIES:
		return p.parseProperties()

	default:
		p.tokenStream.LogError("Unexpected token '%s'", p.tokenStream.CurToken().Type);
		return nil, p.tokenStream.Error()
	}
}

func (p *objectComponentParser) parseFunction() (*ast.Function, error) {

	function := &ast.Function{Type: ast.FUNCTION}

	p.tokenStream.ExpectPeek(token.IDENT)

	function.Name = p.tokenStream.CurToken().Val

	p.tokenStream.ExpectPeek(token.LPAREN)

	p.tokenStream.NextToken()

	function.Parameters = p.parseParameters(token.RPAREN)

	function.Body = p.parseBlockStatement(token.LBRACE, token.RBRACE)

	return function, p.tokenStream.Error()
}

func (p *objectComponentParser) parseCheck() (*ast.Function, error) {

	check := &ast.Function{Type: ast.FUNCTION}

	check.Name = p.tokenStream.CurToken().Val

	p.tokenStream.ExpectPeek(token.LPAREN)

	check.Parameters = []*ast.Parameter{}

	check.Body = p.parseBlockStatement(token.LPAREN, token.RPAREN)

	return check, p.tokenStream.Error()
}

func (p *objectComponentParser) parseHandler() (*ast.Function, error) {

	check := &ast.Function{Type: ast.FUNCTION}

	check.Name = p.tokenStream.CurToken().Val

	p.tokenStream.ExpectPeek(token.LBRACE)

	check.Parameters = []*ast.Parameter{}

	check.Body = p.parseBlockStatement(token.LBRACE, token.RBRACE)

	return check, p.tokenStream.Error()
}

func (p *objectComponentParser) parseWhen() (*ast.When, error) {

	check := &ast.When{Type: ast.WHEN}

	p.tokenStream.ExpectPeek(token.IDENT)

	p.tokenStream.ExpectPeek(token.OBJECTNAME)

	check.Event = p.tokenStream.CurToken().Val

	p.tokenStream.ExpectPeek(token.LBRACE)

	check.Body = p.parseBlockStatement(token.LBRACE, token.RBRACE)

	return check, p.tokenStream.Error()
}

func (p *objectComponentParser) parseParameters(end token.TokenType) []*ast.Parameter {

	list := []*ast.Parameter{}

	if p.tokenStream.CurTokenIs(end) {
		p.tokenStream.NextToken()
		return list
	}

	list = append(list, p.parseParameter())

	for p.tokenStream.PeekTokenIs(token.COMMA) {
		p.tokenStream.NextToken()
		p.tokenStream.NextToken()
		list = append(list, p.parseParameter())
	}

	if !p.tokenStream.ExpectPeek(end) {
		return nil
	}

	p.tokenStream.NextToken()

	return list
}

func (p *objectComponentParser) parseParameter() *ast.Parameter{

	if (!p.tokenStream.CurTokenIs(token.OBJECTNAME)) {
		p.tokenStream.LogError("Expected object name");
		return nil
	}

	param := &ast.Parameter {
		Type: p.tokenStream.CurToken().Val,
	}
	p.tokenStream.ExpectPeek(token.IDENT)

	param.Name = p.tokenStream.CurToken().Val

	return param
}

func (p *objectComponentParser) parseBlockStatement(open token.TokenType, close token.TokenType) *ast.BlockStatement {

	statementParser := NewStatementFromTokenStream(p.tokenStream)

	blkStmnt, err := statementParser.ParseBlockStatementSurroundedBy(open, close)

	if (err != nil) {
		return nil
	}

	p.tokenStream = statementParser.TokenStream();

	p.tokenStream.NextToken()

	return blkStmnt
}

func (p *objectComponentParser) parseProperties() (*ast.Properties, error) {

	props := &ast.Properties{Type:ast.PROPERTIES}

	props.Properties = []*ast.Property{}

	p.tokenStream.ExpectPeek(token.LBRACE)
	p.tokenStream.NextToken()

	for ! p.tokenStream.CurTokenIs(token.RBRACE) && p.tokenStream.Error() == nil  {
		prop := p.parseProperty()
		if prop != nil {
			props.Properties = append(props.Properties, prop)
		}

		if p.tokenStream.CurToken() == nil {
			p.tokenStream.LogError("Unexpected EOF")
			break;
		}
	}


	return props, p.tokenStream.Error()
}

func (p *objectComponentParser) parseProperty() *ast.Property {

	prop := &ast.Property{Type:ast.PROPERTY}

	prop.ValueType = p.tokenStream.CurToken().Val

	p.tokenStream.ExpectPeek(token.IDENT)

	prop.Name = p.tokenStream.CurToken().Val

	if (p.tokenStream.PeekTokenIs(token.SEMICOLON)) {
		p.tokenStream.NextToken()
		p.tokenStream.NextToken()
		return prop
	}

	p.tokenStream.ExpectPeek(token.ASSIGN);
	p.tokenStream.NextToken()

	prop.Exp = p.parseObjectCreation()

	return prop
}

func (p *objectComponentParser) parseObjectCreation() ast.Expression {

	return nil
}



