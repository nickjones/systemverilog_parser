package lexer

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
)

/*
Lexer object contains the state of our parser and provides
a stream for accepting tokens.

Based on work by Rob Pike
http://cuddle.googlecode.com/hg/talk/lex.html#landing-slide
*/
type Lexer struct {
	Name      string
	Input     string
	Tokens    chan lexertoken.Token
	State     LexFn
	PrevState LexFn

	Start int
	Pos   int
	Width int
}

// Backup to the beginning of the last read token.
func (l *Lexer) Backup() {
	l.Pos -= l.Width
}

// CurrentInput returns a slice of the current input from the current lexer start position
// to the current position.
func (l *Lexer) CurrentInput() string {
	return l.Input[l.Start:l.Pos]
}

// Dec decrements the position
func (l *Lexer) Dec() {
	l.Pos--
}

// Emit puts a token onto the token channel. The value of l token is
// read from the input based on the current lexer position.
func (l *Lexer) Emit(tokenType lexertoken.TokenType) {
	l.Tokens <- lexertoken.Token{Type: tokenType, Value: l.Input[l.Start:l.Pos]}
	l.Start = l.Pos
}

// Errorf returns a token with error information.
func (l *Lexer) Errorf(format string, args ...interface{}) LexFn {
	l.Tokens <- lexertoken.Token{
		Type:  lexertoken.TOKEN_ERROR,
		Value: fmt.Sprintf(format, args...),
	}

	return nil
}

// Ignore ignores the current token by setting the lexer's start
// position to the current reading position.
func (l *Lexer) Ignore() {
	l.Start = l.Pos
}

// Inc increments the position
func (l *Lexer) Inc() {
	l.Pos++
	if l.Pos >= utf8.RuneCountInString(l.Input) {
		l.Emit(lexertoken.TOKEN_EOF)
	}
}

// InputToEnd returns a slice of the input from the current lexer position
// to the end of the input string.
func (l *Lexer) InputToEnd() string {
	return l.Input[l.Pos:]
}

// IsEOF returns the true/false if the lexer is at the end of the
// input stream.
func (l *Lexer) IsEOF() bool {
	return l.Pos >= len(l.Input)
}

// IsWhitespace returns true/false if then next character is whitespace
func (l *Lexer) IsWhitespace() bool {
	ch, _ := utf8.DecodeRuneInString(l.Input[l.Pos:])
	return unicode.IsSpace(ch)
}

// Next reads the next rune (character) from the input stream
// and advances the lexer position.
func (l *Lexer) Next() rune {
	if l.Pos >= utf8.RuneCountInString(l.Input) {
		l.Width = 0
		return lexertoken.EOF
	}

	result, width := utf8.DecodeRuneInString(l.Input[l.Pos:])

	l.Width = width
	l.Pos += l.Width
	return result
}

// NextToken returns the next token from the channel
func (l *Lexer) NextToken() lexertoken.Token {
	for {
		select {
		case token := <-l.Tokens:
			return token
		default:
			l.State = l.State(l)
		}
	}
}

// Peek returns the next rune in the stream, then puts the lexer
// position back. Basically reads the next rune without consuming
// it.
func (l *Lexer) Peek() rune {
	rune := l.Next()
	l.Backup()
	return rune
}

// Run starts the lexical analysis and feeding tokens into the
// token channel.
func (l *Lexer) Run() {
	for state := LexBegin; state != nil; {
		state = state(l)
	}

	l.Shutdown()
}

// Shutdown shuts down the token stream
func (l *Lexer) Shutdown() {
	close(l.Tokens)
}

// SkipWhitespace skips whitespace until we get something meaningful.
func (l *Lexer) SkipWhitespace() {
	for {
		ch := l.Next()

		if !unicode.IsSpace(ch) {
			l.Dec()
			break
		}

		if ch == lexertoken.EOF {
			l.Emit(lexertoken.TOKEN_EOF)
			break
		}
	}
}

// HasKeyword iterates through reserved keywords and return true
// if the input matches a keyword.
func (l *Lexer) HasKeyword() bool {
	for _, value := range lexertoken.KEYWORDS {
		if strings.HasPrefix(l.InputToEnd(), value) {
			return true
		}
	}
	return false
}
