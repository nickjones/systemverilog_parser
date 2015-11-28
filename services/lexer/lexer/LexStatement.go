package lexer

import (
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/nickjones/systemverilog_parser/services/lexer/errors"
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
)

/*
This lexer function emits a TOKEN_STATEMENT then returns
the lexer for another statement.
*/
func LexStatement(lexer *Lexer) LexFn {
	var lastChar string
	log.Println("LexStatement")
	lexer.SkipWhitespace()
	lexer.Ignore()
	for {
		if lexer.IsEOF() {
			return lexer.Errorf(errors.LEXER_ERROR_UNEXPECTED_EOF)
		}

		keywordLength := lexer.HasKeyword()
		// End of the line
		if strings.HasPrefix(lexer.InputToEnd(), ";") {
			lexer.Pos++
			lexer.Ignore()
			return LexBegin
		} else if (lastChar == " " || lastChar == "\n") && keywordLength != 0 {
			log.Debugf("Found keyword: length %d\n", keywordLength)
			lexer.Pos += keywordLength
			lexer.Emit(lexertoken.TOKEN_KEYWORD)
			continue
		} else if strings.HasPrefix(lexer.InputToEnd(), ".") {
			// Method call on something
			lexer.Inc()
			lexer.Emit(lexertoken.TOKEN_PERIOD)
			continue
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
		} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.EQUAL_SIGN) {
			lexer.Inc()
			lexer.Emit(lexertoken.TOKEN_EQUAL_SIGN)
			continue
		}
		lastChar = lexer.CurrentInput()[1:]
		lexer.Inc()
	}
}
