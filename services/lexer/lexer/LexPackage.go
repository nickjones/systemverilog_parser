package lexer

import (
	"github.com/nickjones/systemverilog_parser/services/lexer/errors"
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
	"log"
	"strings"
)

/*
This lexer function emits a TOKEN_PACKAGE then returns
the lexer for another statement.
*/
func LexPackage(lexer *Lexer) LexFn {
	log.Println("LexPackage")
	lexer.SkipWhitespace()
	lexer.Ignore()
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		if strings.HasPrefix(lexer.InputToEnd(), ";") {
			lexer.Emit(lexertoken.TOKEN_PACKAGE)
			lexer.Pos++
			return LexBegin
		}

		lexer.Inc()
	}
}

func LexEndPackage(lexer *Lexer) LexFn {
	log.Println("LexEndPackage")
	lexer.Ignore()
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		// Function with no args
		if strings.HasPrefix(lexer.InputToEnd(), "\n") {
			lexer.Emit(lexertoken.TOKEN_ENDPACKAGE)
			lexer.Pos++
			return LexBegin
		} else if strings.HasPrefix(lexer.InputToEnd(), ":") {
			lexer.Emit(lexertoken.TOKEN_ENDPACKAGE)
			lexer.Pos++
			return LexLabel
		}

		lexer.Inc()
	}
}
