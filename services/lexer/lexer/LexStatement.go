package lexer

import (
	"github.com/nickjones/systemverilog_parser/services/lexer/errors"
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
	"log"
	"strings"
)

/*
This lexer function emits a TOKEN_STATEMENT then returns
the lexer for another statement.
*/
func LexStatement(lexer *Lexer) LexFn {
	log.Println("LexStatement")
	lexer.SkipWhitespace()
	lexer.Ignore()
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		// Function with no args
		if strings.HasPrefix(lexer.InputToEnd(), ";") {
			lexer.Emit(lexertoken.TOKEN_STATEMENT)
			lexer.Pos++
			return LexBegin
		}

		lexer.Inc()
	}
}
