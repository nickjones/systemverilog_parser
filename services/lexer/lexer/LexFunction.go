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
func LexFunction(lexer *Lexer) LexFn {
	log.Println("LexFunction")
	lexer.SkipWhitespace()
	lexer.Ignore()
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		// Function with no args
		if strings.HasPrefix(lexer.InputToEnd(), "();") {
			lexer.Emit(lexertoken.TOKEN_FUNCTION)
			lexer.Pos += 3
			return LexBegin
		} else if strings.HasPrefix(lexer.InputToEnd(), "(") {
			lexer.Emit(lexertoken.TOKEN_FUNCTION)
			lexer.Pos++
			return LexFuncArgs
		}

		lexer.Inc()
	}
}
