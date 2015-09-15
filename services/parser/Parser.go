package parser

import (
	"bytes"
	"fmt"
	"log"

	"github.com/nickjones/systemverilog_parser/services/lexer"
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
)

func isEOF(token lexertoken.Token) bool {
	return token.Type == lexertoken.TOKEN_EOF
}

func Parse(fileName, input string) string {
	var buffer bytes.Buffer
	// output := sv.SvFile{
	// 	FileName: fileName,
	// 	SvCode:   make([]sv.SvCode, 0),
	// }

	var token lexertoken.Token

	/* State variables */
	// section := ini.IniSection{}
	// key := ""

	log.Println("Starting lexer and parser for file", fileName, "...")

	l := lexer.BeginLexing(fileName, input)

	for {
		token = l.NextToken()

		// if token.Type != lexertoken.TOKEN_VALUE {
		// 	tokenValue = strings.TrimSpace(token.Value)
		// } else {
		// 	tokenValue = token.Value
		// }

		if isEOF(token) {
			// output.Sections = append(output.Sections, section)
			break
		}

		switch token.Type {
		// case lexertoken.TOKEN_PACKAGE:
		// 	fmt.Printf("\nFound package: %#v\n\n", token)
		// case lexertoken.TOKEN_SECTION:
		// 	/*
		// 	 * Reset tracking variables
		// 	 */
		// 	if len(section.KeyValuePairs) > 0 {
		// 		output.Sections = append(output.Sections, section)
		// 	}

		// 	key = ""

		// 	section.Name = tokenValue
		// 	section.KeyValuePairs = make([]ini.IniKeyValue, 0)

		// case lexertoken.TOKEN_KEY:
		// 	key = tokenValue

		// case lexertoken.TOKEN_VALUE:
		// 	section.KeyValuePairs = append(section.KeyValuePairs, ini.IniKeyValue{Key: key, Value: tokenValue})
		// 	key = ""
		// }
		default:
			fmt.Printf("%s\n", token.String())
		}
	}

	log.Println("Parser has been shutdown")
	return buffer.String()
}
