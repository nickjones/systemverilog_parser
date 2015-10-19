package lexer

import (
	"github.com/nickjones/systemverilog_parser/services/lexer/errors"
	// "github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
	"log"
	// "strings"
)

func LexKeyword(lexer *Lexer) LexFn {
	log.Println("LexKeyword")
	lexer.SkipWhitespace()
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		return LexBegin
	}
}
