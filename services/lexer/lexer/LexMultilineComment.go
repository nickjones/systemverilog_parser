package lexer

import (
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
	"strings"
)

/*
This lexer function emits a TOKEN_COMMENT
*/
func LexMultilineComment(lexer *Lexer) LexFn {
	lexer.PrevState = LexMultilineComment
	for {
		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.MULTI_LINE_COMMENT_STOP) {
			lexer.Emit(lexertoken.TOKEN_COMMENT)
			return LexBegin
		}

		lexer.Inc()
	}
}
