package lexer

import (
	"github.com/nickjones/systemverilog_parser/services/lexer/errors"
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
	"log"
	"strings"
)

/*
This lexer function emits a TOKEN_LABEL then returns
the lexer for another statement.
*/
func LexLabel(lexer *Lexer) LexFn {
	log.Println("LexLabel")
	lexer.Ignore()
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		// Function with no args
		if strings.HasPrefix(lexer.InputToEnd(), "\n") {
			lexer.Emit(lexertoken.TOKEN_LABEL)
			return LexBegin
		}

		lexer.Inc()
	}
}
