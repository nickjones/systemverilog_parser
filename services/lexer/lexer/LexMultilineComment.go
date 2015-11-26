package lexer

import (
	"strings"

	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
)

/*
This lexer function emits a TOKEN_COMMENT
*/
func LexMultilineComment(lexer *Lexer) LexFn {
	lexer.PrevState = LexMultilineComment
	lexer.Pos += len(lexertoken.MULTI_LINE_COMMENT_START)
	lexer.SkipWhitespace()
	lexer.Ignore()
	for {
		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.MULTI_LINE_COMMENT_STOP) {
			lexer.Emit(lexertoken.TOKEN_COMMENT)
			lexer.Pos += len(lexertoken.MULTI_LINE_COMMENT_STOP)
			lexer.Ignore()
			return LexBegin
		}

		lexer.Inc()
	}
}
