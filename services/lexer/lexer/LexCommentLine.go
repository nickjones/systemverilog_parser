package lexer

import (
	"strings"

	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
)

/*
This lexer function emits a TOKEN_COMMENT
*/
func LexCommentLine(lexer *Lexer) LexFn {
	lexer.Pos += len(lexertoken.SINGLE_LINE_COMMENT)
	lexer.SkipWhitespace()
	lexer.Ignore()
	for {
		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.NEWLINE) {
			lexer.Emit(lexertoken.TOKEN_COMMENT)
			return LexBegin
		}

		lexer.Inc()
	}
}
