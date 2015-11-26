package lexer

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/nickjones/systemverilog_parser/services/lexer/errors"
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
)

// LexFunction emits a TOKEN_FUNCTION then returns
// the lexer for another statement.
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

// LexEndFunc marks the end of a function ('endfunction')
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

// LexFuncArgs parses function arguments
func LexFuncArgs(lexer *Lexer) LexFn {
	log.Println("LexFuncArgs")
	lexer.Ignore()
	for {
		// log.Debugf("Lex pos: %d\n", lexer.Pos)
		// log.Debugf("inputtoend: %s\n", lexer.InputToEnd())

		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		if length := lexer.HasKeyword(); length != 0 {
			log.Debugf("Found keyword inside function arguments: length %d\n", length)
			log.Debugf("InputToEnd: %s\n", lexer.InputToEnd())
			lexer.Pos += length
			lexer.Emit(lexertoken.TOKEN_KEYWORD)
			lexer.SkipWhitespace()
			lexer.Ignore()
			continue
		} else if strings.HasSuffix(lexer.CurrentInput(), ",") {
			if strings.Contains(lexer.CurrentInput(), " ") {
				log.Debugln("Found space before comma, user defined type")
				log.Debugf("CurrentInput after finding space: %s\n", lexer.CurrentInput())
				lexer.Dec() // Exclude space from token
				lexer.Emit(lexertoken.TOKEN_USER_TYPE)
				lexer.SkipWhitespace()
				lexer.Ignore()
				continue
			} else {
				log.Debugln("No user type, found identifier")
				lexer.Dec() // Exclude comma from token
				lexer.Emit(lexertoken.TOKEN_IDENTIFIER)
				lexer.Inc()
				lexer.SkipWhitespace()
				lexer.Ignore()
				return LexFuncArgs
			}
		} else if strings.HasPrefix(lexer.InputToEnd(), ");") {
			semiColonPos := lexer.Pos
			currInput := lexer.CurrentInput()
			if strings.Contains(currInput, " ") {
				log.Debugln("Found space before ');', user defined type")
				log.Debugf("CurrentInput after finding space: %s\n", currInput)
				// splits := strings.Split(currInput, " ")
				// log.Debugf("first: %s second: %s\n", splits[0], splits[1])
				eol := lexer.Pos // Save position of ');'
				// Walk backward to the start of this section
				// lexer.Pos -= len(currInput)
				// log.Debugf("curr: %s\n", currInput)
				// Move forward the length of the first portion
				// lexer.Pos += len(splits[0])

				// Move the lexer back to the beginning of this pair
				lexer.Pos = lexer.Start
				// Walk forward until we hit the space
				for {
					if strings.HasPrefix(lexer.InputToEnd(), " ") {
						break
					}
					lexer.Inc()
				}
				log.Debugf("Up until space: %s\n", lexer.CurrentInput())
				lexer.Emit(lexertoken.TOKEN_USER_TYPE)
				lexer.SkipWhitespace()
				lexer.Ignore()
				lexer.Pos = eol
				lexer.Emit(lexertoken.TOKEN_IDENTIFIER)
				lexer.Pos = semiColonPos + 2 // ');'
				lexer.Ignore()
				return LexBegin
			} else {
				log.Debugln("No user type, found identifier")
				lexer.Emit(lexertoken.TOKEN_IDENTIFIER)
				lexer.Inc()
				lexer.SkipWhitespace()
				lexer.Ignore()
				return LexFuncArgs
			}
		}

		lexer.Inc()
	}
}
