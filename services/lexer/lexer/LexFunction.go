package lexer

import (
	"github.com/nickjones/systemverilog_parser/services/lexer/errors"
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
	"log"
	"strings"
)

/*
This lexer function emits a TOKEN_FUNCTION then returns
the lexer for another statement.
*/
func LexFunction(lexer *Lexer) LexFn {
	log.Println("LexFunction")
	lexer.SkipWhitespace()
	lexer.Ignore()
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		// Function with no args
		if strings.HasPrefix(lexer.InputToEnd(), "();") {
			lexer.Emit(lexertoken.TOKEN_FUNCTION)
			lexer.Pos += 3
			return LexBegin
		} else if strings.HasPrefix(lexer.InputToEnd(), "(") {
			lexer.Emit(lexertoken.TOKEN_FUNCTION)
			lexer.Pos++
			return LexFuncArgs
		}

		lexer.Inc()
	}
}

func LexEndFunc(lexer *Lexer) LexFn {
	log.Println("LexEndFunc")
	lexer.Ignore()
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		// Function with no args
		if strings.HasPrefix(lexer.InputToEnd(), "\n") {
			lexer.Emit(lexertoken.TOKEN_ENDFUNCTION)
			lexer.Pos++
			return LexBegin
		} else if strings.HasPrefix(lexer.InputToEnd(), ":") {
			lexer.Emit(lexertoken.TOKEN_ENDFUNCTION)
			lexer.Pos++
			return LexLabel
		}

		lexer.Inc()
	}
}

func LexFuncArgs(lexer *Lexer) LexFn {
	log.Println("LexFuncArgs")
	lexer.SkipWhitespace()
	lexer.Ignore()
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		// Function with no args
		if strings.HasPrefix(lexer.InputToEnd(), ",") {
			lexer.Emit(lexertoken.TOKEN_FUNCTION_ARGUMENT)
			lexer.Pos++
			return LexFuncArgs
		} else if strings.HasPrefix(lexer.InputToEnd(), ");") {
			lexer.Emit(lexertoken.TOKEN_FUNCTION_ARGUMENT)
			lexer.Pos += 2
			return LexBegin
		}

		lexer.Inc()
	}
}
