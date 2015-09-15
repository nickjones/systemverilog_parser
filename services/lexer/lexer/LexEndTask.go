package lexer

import (
	"github.com/nickjones/systemverilog_parser/services/lexer/errors"
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
	"log"
	"strings"
)

func LexEndTask(lexer *Lexer) LexFn {
	log.Println("LexEndTask")
	lexer.Ignore()
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		// Function with no args
		if strings.HasPrefix(lexer.InputToEnd(), "\n") {
			lexer.Emit(lexertoken.TOKEN_END_TASK)
			lexer.Pos++
			return LexBegin
		} else if strings.HasPrefix(lexer.InputToEnd(), ":") {
			lexer.Emit(lexertoken.TOKEN_END_TASK)
			lexer.Pos++
			return LexLabel
		}

		lexer.Inc()
	}
}
