package lexer

import (
	"github.com/nickjones/systemverilog_parser/services/lexer/errors"
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
	"log"
	"strings"
)

/*
This lexer function emits a TOKEN_CLASS or TOKEN_SUPERCLASS then returns
the lexer for another statement.
*/
// TODO: Add parameterization support
func LexClass(lexer *Lexer) LexFn {
	log.Println("LexClass")
	lexer.SkipWhitespace()
	lexer.Ignore()
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		if strings.HasPrefix(lexer.InputToEnd(), " extends ") {
			lexer.Emit(lexertoken.TOKEN_CLASS)
			lexer.Pos += 9
			return LexSuperClass
		} else if strings.HasPrefix(lexer.InputToEnd(), ";") {
			lexer.Emit(lexertoken.TOKEN_CLASS)
			lexer.Pos++
			return LexInsideClass
		}

		lexer.Inc()
	}
}
