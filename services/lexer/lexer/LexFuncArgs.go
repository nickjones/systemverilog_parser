package lexer

import (
	"github.com/nickjones/systemverilog_parser/services/lexer/errors"
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
	"log"
	"strings"
)

/*
This lexer function emits a TOKEN_FUNCTION then returns
the lexer for another statement.
*/
func LexFuncArgs(lexer *Lexer) LexFn {
	log.Println("LexFuncArgs")
	lexer.SkipWhitespace()
	lexer.Ignore()
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		// Function with no args
		if strings.HasPrefix(lexer.InputToEnd(), ",") {
			lexer.Emit(lexertoken.TOKEN_FUNCTION_ARGUMENT)
			lexer.Pos++
			return LexFuncArgs
		} else if strings.HasPrefix(lexer.InputToEnd(), ");") {
			lexer.Emit(lexertoken.TOKEN_FUNCTION_ARGUMENT)
			lexer.Pos += 2
			return LexBegin
		}

		lexer.Inc()
	}
}
