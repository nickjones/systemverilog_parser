package lexer

import (
	"github.com/nickjones/systemverilog_parser/services/lexer/errors"
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
	"strings"
)

/*
This lexer function determines the kind of define is used
and deferrs emitting until it is unknown
*/
func LexDefineBegin(lexer *Lexer) LexFn {
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		if strings.HasPrefix(lexer.InputToEnd(), lexertoken.INCLUDE) {
			lexer.Pos += len(lexertoken.INCLUDE)
			lexer.Emit(lexertoken.TOKEN_INCLUDE)
			return LexInclude
		} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.IFDEF) {
			lexer.Pos += len(lexertoken.IFDEF)
			lexer.Emit(lexertoken.TOKEN_IFDEF)
			return LexBegin
		} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.IFNDEF) {
			lexer.Pos += len(lexertoken.IFNDEF)
			lexer.Emit(lexertoken.TOKEN_IFNDEF)
			return LexBegin
		} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.DEFINE) {
			lexer.Pos += len(lexertoken.DEFINE)
			lexer.Emit(lexertoken.TOKEN_DEFINE)
			return LexBegin
		} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDIFDEF) {
			lexer.Pos += len(lexertoken.ENDIFDEF)
			lexer.Emit(lexertoken.TOKEN_ENDIFDEF)
			return LexBegin
		} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ELSIF) {
			lexer.Pos += len(lexertoken.ELSIF)
			lexer.Emit(lexertoken.TOKEN_ELSIF)
			return LexBegin
		} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ELSE) {
			lexer.Pos += len(lexertoken.ELSE)
			lexer.Emit(lexertoken.TOKEN_ELSE)
			return LexBegin
		} else {
		}
	}
}
