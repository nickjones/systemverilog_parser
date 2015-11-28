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

	log.Println("Starting lexer and parser for file", fileName, "...")

	l := lexer.BeginLexing(fileName, input)

	for {
		token = l.NextToken()

		if isEOF(token) {
			break
		}

		switch token.Type {
		// case lexertoken.TOKEN_PACKAGE:
		// 	fmt.Printf("\nFound package: %#v\n\n", token)
		default:
			fmt.Printf("Token: %d value: %s\n", token.Type, token.String())
		}
	}

	log.Println("Parser has been shutdown")
	return buffer.String()
}
