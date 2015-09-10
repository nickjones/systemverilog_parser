package lexer

import (
	"github.com/nickjones/systemverilog_parser/services/lexer/errors"
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
	"strings"
)

/*
This lexer function emits a TOKEN_MODULE then returns
the lexer for another statement.
*/
func LexModule(lexer *Lexer) LexFn {
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		if strings.HasPrefix(lexer.InputToEnd(), ";") {
			lexer.Pos++
			lexer.Emit(lexertoken.TOKEN_MODULE)
			return LexBegin
		}

		lexer.Inc()
	}
}
