package tokenizer

import (
	"strings"
	"unicode/utf8"
	tok "github.com/domain-query-language/dql-server/src/server/adapter/parser/token"
	"fmt"
)

// lexer holds the state of the scanner.
type lexer struct {
	name  string    // used only for error reports.
	input string    // the string being scanned.
	length int	// length of the input
	start int       // start position of this item.
	pos   int       // current position in the input.
	width int       // width of last rune read from input.
	tokens []tok.Token
}

type stateFn func(*lexer) stateFn

var easyLexKeywords = []tok.TokenType{}
var easyLexTokens = []tok.TokenType{}

const EOF = -1

func lex(name, input string) (*lexer) {

	l := &lexer {
		name:  name,
		input: input,
		length: len(input),
	}

	easyLexKeywords = []tok.TokenType {
		tok.CREATE,
		tok.LIST,
		tok.USING,
		tok.FOR,
		tok.IN,
		tok.WITHIN,
		tok.ASSERT,
		tok.RUN,
		tok.APPLY,
		tok.WHEN,
		tok.NOT,
		tok.AND,
		tok.OR,
		tok.PROPERTIES,
		tok.CHECK,
		tok.HANDLER,
		tok.FUNCTION,
		tok.IF,
		tok.ELSEIF,
		tok.ELSE,
		tok.FOREACH,
		tok.RETURN,
		tok.AS,
		tok.ON,
	}

	easyLexTokens = []tok.TokenType{
		tok.STRONGARROW,
		tok.EQ,
		tok.SEMICOLON,
		tok.ASSIGN,
		tok.CLASSCLOSE,
		tok.COLON,
		tok.LBRACE,
		tok.RBRACE,
		tok.LPAREN,
		tok.RPAREN,
		tok.LBRACKET,
		tok.RBRACKET,
		tok.NOTEQ,
		tok.COMMA,
		tok.ARROW,
		tok.PLUS,
		tok.MINUS,
		tok.BANG,
		tok.ASTERISK,
		tok.SLASH,
		tok.LTOREQ,
		tok.GTOREQ,
		tok.LT,
		tok.GT,
		tok.REMAINDER,
	}

	l.run()
	return l
}

func (l *lexer) run() {

	for state := lexToken; state != nil; {
		state = state(l)
	}

	return
}

func (l *lexer) emit(t tok.TokenType) {

	l.tokens = append(l.tokens, tok.Token{t, l.input[l.start:l.pos], l.start});
	l.start = l.pos
}

func (l *lexer) isNextPrefix(prefix string) bool {

	var unlexed string;

	if (l.pos + len(prefix) < l.length) {
		unlexed = l.input[l.pos: l.pos + len(prefix)];
	} else {
		unlexed = l.input[l.pos:];
	}

	return strings.HasPrefix(strings.ToLower(unlexed), strings.ToLower(prefix));
}

//Check if the prefix matches and is not followed immediately by another identifier character
func (l *lexer) isKeyWordAndNotIdentifier(prefix string) bool {

	hasPrefix := l.isNextPrefix(prefix)

	if (!hasPrefix) {
		return false
	}

	nextRune, _ := l.runeAtPos(l.pos + len(prefix));

	if (isDigit(nextRune) || isLetter(nextRune)) {
		return false
	}

	return true
}

//Tries to match the current prefix against a series of strings, if it doesn't get a match, it logs the error
func (l *lexer) lexMatchingPrefix(prefixes []tok.TokenType) stateFn {

	for _, prefix := range prefixes {

		if (l.isKeyWordAndNotIdentifier(string(prefix))) {
			l.pos += len(prefix)
			l.emit(prefix)
			return lexToken
		}
	}

	expected := []string{}

	for _, prefix := range prefixes {
		expected = append(expected, string(prefix))
	}

	found := l.scanWord()
	l.err(strings.Join(expected, ", "), found)

	return nil
}

var whitespace = []int{' ', '\n', '\r', '\t'}

func contains(s []int, e int) bool {

	for _, a := range s {

		if a == e {
			return true
		}
	}

	return false
}

//Just skip over whitespace
func (l *lexer) skipWS() {

	for {
		switch r := l.next(); {

		case contains(whitespace, r) :
			l.ignore()

		default :
			l.backup()
			return
		}
	}
}

//Move past the white space but don't skip it
func (l *lexer) consumeWS() {

	for {
		if !contains(whitespace, l.next()) {
			l.backup()
			break;
		}
	}
}

//Get the next rune in the input.
func (l *lexer) next() int {

	r, width := l.runeAtPos(l.pos)
	if r == EOF {
		l.width = 0
		return r
	}
	l.width = width
	l.pos += l.width
	return r
}

//Get the rune at this position
func (l *lexer) runeAtPos(pos int) (rn int, width int) {

	if pos >= len(l.input) {
		rn = EOF
		width = 0
		return
	}
	var r rune;
	r, width = utf8.DecodeRuneInString(l.input[pos:])
	rn = int(r)
	return
}

//Skip the next token
func (l *lexer) skip() {

	l.next()
	l.ignore()
}

//Skip a string of tokens
func (l *lexer) skipStr(string string) {

	l.pos += len(string)
	l.skipWS()
}

func (l *lexer) ignore() {

	l.start = l.pos
}

func (l *lexer) backup() {

	l.pos -= l.width
}

func (l *lexer) peek() int {

	r := l.next()
	l.backup()
	return r
}

func (l *lexer) err(expected string, found string) stateFn {

	msg := "Parse error, expected "+expected+", found "+found;
	l.tokens = append(l.tokens, tok.Token{tok.ERROR, msg, l.start});
	l.start = l.pos

	return nil
}

//Scan until next space or non identifier character
func (l *lexer) scanWord() string {

	for {
		if (!isLetter(l.peek()) && !isDigit(l.peek())) {
			break;
		}

		l.next();
	}

	return l.input[ l.start : l.pos ]
}

//Scan until the quotedname is finished, or the EOF
func (l *lexer) scanQuotedName() (found string, isEOF bool) {

	for {
		if (l.peek() == '\'') {
			isEOF = false
			break;
		}

		if (l.peek() == EOF) {
			isEOF = true
			break;
		}

		l.next();
	}

	found = l.input[ l.start : l.pos ]

	return
}

var objectTypes = []string {
	"value",
	"entity",
	"event",
	"command",
	"invariant",
	"projection",
	"query",
}

func (l *lexer) isTypeRefence() bool {

	if (l.isKeyWordAndNotIdentifier(tok.STRING)) {
		return true;
	}
	if (l.isKeyWordAndNotIdentifier(tok.INTEGER)) {
		return true;
	}
	if (l.isKeyWordAndNotIdentifier(tok.FLOAT)) {
		return true;
	}
	if (l.isKeyWordAndNotIdentifier(tok.BOOLEAN)) {
		return true;
	}
	for _, objectType := range objectTypes {
		if (l.isNextPrefix(objectType+"\\")) {
			return true;
		}
	}
	return false
}

func lexToken(l *lexer) stateFn {

	l.skipWS();

	if (l.peek() == EOF) {
		return nil;
	}

	//Complex/partal check, special lexing rule
	if (l.isTypeRefence()) {
		return lexTypeRef
	}
	if (l.isNextPrefix("'")) {
		return lexNSObjectName
	}
	if (l.isNextPrefix("\"")) {
		return lexString
	}
	if (l.isNextPrefix(tok.CLASSOPEN)) {
		return lexClassOpen;
	}

	// If keyword, then lex
	for _, token := range easyLexKeywords {
		if l.isKeyWordAndNotIdentifier(string(token)) {
			return l.lexAsToken(token)
		}
	}

	// If match, then lex
	for _, token := range easyLexTokens {
		if l.isNextPrefix(string(token)) {
			return l.lexAsToken(token)
		}
	}

	if (isDigit(l.peek())) {
		return lexNumber;
	}
	if (isLetter(l.peek())) {
		return lexIdentifier;
	}

	return l.err("keyword", fmt.Sprintf("%c", rune(l.peek())));
}

func (l *lexer) lexQuotedStringAsToken(tokenType tok.TokenType) stateFn {

	l.skipWS()
	nxt := l.next()
	if (nxt == EOF) {
		l.err("'", "EOF")
		return nil
	}
	if (nxt != '\'') {
		l.err("'", l.scanWord())
		return nil
	}
	l.ignore();

	found, isEOF := l.scanQuotedName()
	if (isEOF) {
		l.err("'", "EOF")
		return nil
	}

	if (found == "") {
		l.err("value name", "empty name")
		return nil
	}

	l.emit(tokenType)
	l.skip()
	return lexToken;
}

func (l *lexer) lexAsToken(tokenType tok.TokenType) stateFn {

	l.pos += len(tokenType)
	l.emit(tokenType)
	return lexToken
}

func lexNSObjectName(l *lexer) stateFn {

	return l.lexQuotedStringAsToken(tok.OBJECTNAME)
}

func lexTypeRef(l *lexer) stateFn {

	for {
		if (contains(whitespace, l.peek()) || l.peek() == EOF) {
			break;
		}
		l.next();
	}

	l.emit(tok.OBJECTNAME)
	l.skipWS()
	return lexIdentifier
}

func lexIdentifier(l *lexer) stateFn {

	word := l.scanWord()
	if (word == "true" || word == "false") {
		l.emit(tok.BOOLEAN)
	} else if (word == "null") {
		l.emit(tok.NULL)
	} else {
		l.emit(tok.IDENT)
	}
	return lexToken;
}

func lexString(l *lexer) stateFn {

	l.skip()
	for {
		if l.next() == '"' {
			l.backup()
			break;
		}
	}

	l.emit(tok.STRING)
	l.skip()

	return lexToken
}

func lexNumber(l *lexer) stateFn {

	hasDot := false;
	for {
		if (l.peek() == '.') {
			hasDot = true;
		}
		if (!isDigit(l.peek()) && l.peek() != '.') {
			break;
		}
		l.next();
	}
	if (hasDot) {
		l.emit(tok.FLOAT)
	} else {
		l.emit(tok.INTEGER)
	}

	return lexToken;
}

func lexClassOpen(l *lexer) stateFn {

	l.lexAsToken(tok.CLASSOPEN);
	l.skipWS()
	return lexToken;
}

func isLetter(ch int) bool {

	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch int) bool {

	return '0' <= ch && ch <= '9'
}