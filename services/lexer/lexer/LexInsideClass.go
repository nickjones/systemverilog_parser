package lexer

import (
	"strings"

	"github.com/nickjones/systemverilog_parser/services/lexer/errors"
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
)

/*
This lexer function starts everything off. It determines if we are
beginning with a key/value assignment or a section.
*/
func LexInsideClass(lexer *Lexer) LexFn {
	lexer.SkipWhitespace()
	lexer.PrevState = LexInsideClass

	if strings.HasPrefix(lexer.InputToEnd(), lexertoken.IMPORT) {
		return LexImport
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.DEFINE) {
		return LexDefineBegin
	} else {
		return lexer.Errorf(errors.LEXER_ERROR_UNPARSABLE_STATEMENT)
	}
}
