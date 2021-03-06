package lexer

import (
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
)

/*
This lexer function emits a TOKEN_LEFT_BRACKET then returns
the lexer for a section header.
*/
func LexLeftBracket(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.LEFT_BRACKET)
	lexer.Emit(lexertoken.TOKEN_LEFT_BRACKET)
	// FIXME: This is just flat wrong
	return LexBegin
}
