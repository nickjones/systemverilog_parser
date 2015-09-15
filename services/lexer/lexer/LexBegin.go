package lexer

import (
	"log"
	"strings"

	// "github.com/nickjones/systemverilog_parser/services/lexer/errors"
	"github.com/nickjones/systemverilog_parser/services/lexer/lexertoken"
)

/*
This lexer function starts everything off. It determines if we are
beginning with a key/value assignment or a section.
*/
func LexBegin(lexer *Lexer) LexFn {
	lexer.SkipWhitespace()

	log.Println("LexBegin")
	// TODO: Abstract classes
	if strings.HasPrefix(lexer.InputToEnd(), lexertoken.SINGLE_LINE_COMMENT) {
		return LexCommentLine
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.MULTI_LINE_COMMENT_START) {
		return LexMultilineComment
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.PACKAGE) {
		lexer.Pos += len(lexertoken.PACKAGE)
		return LexPackage
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.IMPORT) {
		lexer.Pos += len(lexertoken.IMPORT)
		return LexImport
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.CLASS) {
		lexer.Pos += len(lexertoken.CLASS)
		return LexClass
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.MODULE) {
		lexer.Pos += len(lexertoken.MODULE)
		return LexModule
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.DEFINE) {
		return LexDefineBegin
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.FUNCTION) {
		lexer.Pos += len(lexertoken.FUNCTION)
		return LexFunction
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDFUNCTION) {
		lexer.Pos += len(lexertoken.ENDFUNCTION)
		return LexEndFunc
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDCLASS) {
		lexer.Pos += len(lexertoken.ENDCLASS)
		return LexEndClass
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDTASK) {
		lexer.Pos += len(lexertoken.ENDTASK)
		return LexEndTask
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDMODULE) {
		lexer.Pos += len(lexertoken.ENDMODULE)
		return LexEndModule
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDPACKAGE) {
		lexer.Pos += len(lexertoken.ENDPACKAGE)
		return LexEndPackage
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.TASK) {
		lexer.Pos += len(lexertoken.TASK)
		return LexTask
	} else {
		return LexStatement
		// return lexer.Errorf(errors.LEXER_ERROR_UNPARSABLE_STATEMENT)
	}
}
