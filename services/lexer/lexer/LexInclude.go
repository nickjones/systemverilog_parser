package lexer

import (
	"github.com/nickjones/systemverilog_parser/services/lexer/errors"
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
	"strings"
)

/*
This lexer function extracts the included file
*/
func LexInclude(lexer *Lexer) LexFn {
	if strings.HasPrefix(lexer.InputToEnd(), lexertoken.DOUBLE_QUOTE) {
		lexer.Inc()
	} else {
		// Thrown an error for a malformed quoted string
		return lexer.Errorf(errors.LEXER_ERROR_EXPECTED_STRING)
	}

	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.DOUBLE_QUOTE) {
			lexer.Pos += len(lexertoken.DOUBLE_QUOTE)
			lexer.Emit(lexertoken.TOKEN_INCLUDE)
			return LexBegin
		}

		lexer.Inc()
	}
}
