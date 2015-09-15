package lexer

import (
	"github.com/nickjones/systemverilog_parser/services/lexer/errors"
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
	"log"
	"strings"
)

/*
This lexer function emits a TOKEN_TASK then returns
the lexer for another statement.
*/
func LexTask(lexer *Lexer) LexFn {
	log.Println("LexTask")
	lexer.SkipWhitespace()
	lexer.Ignore()
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		// Function with no args
		if strings.HasPrefix(lexer.InputToEnd(), "();") {
			lexer.Emit(lexertoken.TOKEN_TASK)
			lexer.Pos += 3
			return LexBegin
		} else if strings.HasPrefix(lexer.InputToEnd(), "(") {
			lexer.Emit(lexertoken.TOKEN_TASK)
			lexer.Pos++
			return LexFuncArgs
		}

		lexer.Inc()
	}
}
