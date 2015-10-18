package lexer

import (
	"github.com/nickjones/systemverilog_parser/services/lexer/errors"
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
	"log"
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

func LexEndModule(lexer *Lexer) LexFn {
	log.Println("LexEndModule")
	lexer.Ignore()
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		// Function with no args
		if strings.HasPrefix(lexer.InputToEnd(), "\n") {
			lexer.Emit(lexertoken.TOKEN_ENDMODULE)
			lexer.Pos++
			return LexBegin
		} else if strings.HasPrefix(lexer.InputToEnd(), ":") {
			lexer.Emit(lexertoken.TOKEN_ENDMODULE)
			lexer.Pos++
			return LexLabel
		}

		lexer.Inc()
	}
}
