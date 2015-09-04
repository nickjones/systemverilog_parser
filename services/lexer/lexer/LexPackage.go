package lexer

import (
	"github.com/nickjones/systemverilog_parser/services/lexer/errors"
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
	"strings"
)

/*
This lexer function emits a TOKEN_PACKAGE then returns
the lexer for another statement.
*/
func LexPackage(lexer *Lexer) LexFn {
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		if strings.HasPrefix(lexer.InputToEnd(), ";") {
			lexer.Emit(lexertoken.TOKEN_PACKAGE)
			return LexBegin
		}

		lexer.Inc()
	}
}
