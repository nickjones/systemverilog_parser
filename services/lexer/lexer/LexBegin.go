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

	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ACCEPT_ON) {
		lexer.Pos += len(lexertoken.ACCEPT_ON)
		lexer.Emit(lexertoken.TOKEN_ACCEPT_ON)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ALIAS) {
		lexer.Pos += len(lexertoken.ALIAS)
		lexer.Emit(lexertoken.TOKEN_ALIAS)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ALWAYS) {
		lexer.Pos += len(lexertoken.ALWAYS)
		lexer.Emit(lexertoken.TOKEN_ALWAYS)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ALWAYS_COMB) {
		lexer.Pos += len(lexertoken.ALWAYS_COMB)
		lexer.Emit(lexertoken.TOKEN_ALWAYS_COMB)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ALWAYS_FF) {
		lexer.Pos += len(lexertoken.ALWAYS_FF)
		lexer.Emit(lexertoken.TOKEN_ALWAYS_FF)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ALWAYS_LATCH) {
		lexer.Pos += len(lexertoken.ALWAYS_LATCH)
		lexer.Emit(lexertoken.TOKEN_ALWAYS_LATCH)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.AND) {
		lexer.Pos += len(lexertoken.AND)
		lexer.Emit(lexertoken.TOKEN_AND)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ASSERT) {
		lexer.Pos += len(lexertoken.ASSERT)
		lexer.Emit(lexertoken.TOKEN_ASSERT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ASSIGN) {
		lexer.Pos += len(lexertoken.ASSIGN)
		lexer.Emit(lexertoken.TOKEN_ASSIGN)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ASSUME) {
		lexer.Pos += len(lexertoken.ASSUME)
		lexer.Emit(lexertoken.TOKEN_ASSUME)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.AUTOMATIC) {
		lexer.Pos += len(lexertoken.AUTOMATIC)
		lexer.Emit(lexertoken.TOKEN_AUTOMATIC)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.BEFORE) {
		lexer.Pos += len(lexertoken.BEFORE)
		lexer.Emit(lexertoken.TOKEN_BEFORE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.BEGIN) {
		lexer.Pos += len(lexertoken.BEGIN)
		lexer.Emit(lexertoken.TOKEN_BEGIN)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.BIND) {
		lexer.Pos += len(lexertoken.BIND)
		lexer.Emit(lexertoken.TOKEN_BIND)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.BINS) {
		lexer.Pos += len(lexertoken.BINS)
		lexer.Emit(lexertoken.TOKEN_BINS)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.BINSOF) {
		lexer.Pos += len(lexertoken.BINSOF)
		lexer.Emit(lexertoken.TOKEN_BINSOF)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.BIT) {
		lexer.Pos += len(lexertoken.BIT)
		lexer.Emit(lexertoken.TOKEN_BIT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.BREAK) {
		lexer.Pos += len(lexertoken.BREAK)
		lexer.Emit(lexertoken.TOKEN_BREAK)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.BUF) {
		lexer.Pos += len(lexertoken.BUF)
		lexer.Emit(lexertoken.TOKEN_BUF)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.BUFIF0) {
		lexer.Pos += len(lexertoken.BUFIF0)
		lexer.Emit(lexertoken.TOKEN_BUFIF0)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.BUFIF1) {
		lexer.Pos += len(lexertoken.BUFIF1)
		lexer.Emit(lexertoken.TOKEN_BUFIF1)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.BYTE) {
		lexer.Pos += len(lexertoken.BYTE)
		lexer.Emit(lexertoken.TOKEN_BYTE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.CASE) {
		lexer.Pos += len(lexertoken.CASE)
		lexer.Emit(lexertoken.TOKEN_CASE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.CASEX) {
		lexer.Pos += len(lexertoken.CASEX)
		lexer.Emit(lexertoken.TOKEN_CASEX)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.CASEZ) {
		lexer.Pos += len(lexertoken.CASEZ)
		lexer.Emit(lexertoken.TOKEN_CASEZ)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.CELL) {
		lexer.Pos += len(lexertoken.CELL)
		lexer.Emit(lexertoken.TOKEN_CELL)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.CHANDLE) {
		lexer.Pos += len(lexertoken.CHANDLE)
		lexer.Emit(lexertoken.TOKEN_CHANDLE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.CHECKER) {
		lexer.Pos += len(lexertoken.CHECKER)
		lexer.Emit(lexertoken.TOKEN_CHECKER)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.CLASS) {
		lexer.Pos += len(lexertoken.CLASS)
		lexer.Emit(lexertoken.TOKEN_CLASS)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.CLOCKING) {
		lexer.Pos += len(lexertoken.CLOCKING)
		lexer.Emit(lexertoken.TOKEN_CLOCKING)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.CMOS) {
		lexer.Pos += len(lexertoken.CMOS)
		lexer.Emit(lexertoken.TOKEN_CMOS)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.CONFIG) {
		lexer.Pos += len(lexertoken.CONFIG)
		lexer.Emit(lexertoken.TOKEN_CONFIG)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.CONST) {
		lexer.Pos += len(lexertoken.CONST)
		lexer.Emit(lexertoken.TOKEN_CONST)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.CONSTRAINT) {
		lexer.Pos += len(lexertoken.CONSTRAINT)
		lexer.Emit(lexertoken.TOKEN_CONSTRAINT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.CONTEXT) {
		lexer.Pos += len(lexertoken.CONTEXT)
		lexer.Emit(lexertoken.TOKEN_CONTEXT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.CONTINUE) {
		lexer.Pos += len(lexertoken.CONTINUE)
		lexer.Emit(lexertoken.TOKEN_CONTINUE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.COVER) {
		lexer.Pos += len(lexertoken.COVER)
		lexer.Emit(lexertoken.TOKEN_COVER)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.COVERGROUP) {
		lexer.Pos += len(lexertoken.COVERGROUP)
		lexer.Emit(lexertoken.TOKEN_COVERGROUP)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.COVERPOINT) {
		lexer.Pos += len(lexertoken.COVERPOINT)
		lexer.Emit(lexertoken.TOKEN_COVERPOINT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.CROSS) {
		lexer.Pos += len(lexertoken.CROSS)
		lexer.Emit(lexertoken.TOKEN_CROSS)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.DEASSIGN) {
		lexer.Pos += len(lexertoken.DEASSIGN)
		lexer.Emit(lexertoken.TOKEN_DEASSIGN)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.DEFAULT) {
		lexer.Pos += len(lexertoken.DEFAULT)
		lexer.Emit(lexertoken.TOKEN_DEFAULT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.DEFPARAM) {
		lexer.Pos += len(lexertoken.DEFPARAM)
		lexer.Emit(lexertoken.TOKEN_DEFPARAM)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.DESIGN) {
		lexer.Pos += len(lexertoken.DESIGN)
		lexer.Emit(lexertoken.TOKEN_DESIGN)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.DISABLE) {
		lexer.Pos += len(lexertoken.DISABLE)
		lexer.Emit(lexertoken.TOKEN_DISABLE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.DIST) {
		lexer.Pos += len(lexertoken.DIST)
		lexer.Emit(lexertoken.TOKEN_DIST)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.DO) {
		lexer.Pos += len(lexertoken.DO)
		lexer.Emit(lexertoken.TOKEN_DO)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.EDGE) {
		lexer.Pos += len(lexertoken.EDGE)
		lexer.Emit(lexertoken.TOKEN_EDGE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ELSE) {
		lexer.Pos += len(lexertoken.ELSE)
		lexer.Emit(lexertoken.TOKEN_ELSE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.END) {
		lexer.Pos += len(lexertoken.END)
		lexer.Emit(lexertoken.TOKEN_END)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDCASE) {
		lexer.Pos += len(lexertoken.ENDCASE)
		lexer.Emit(lexertoken.TOKEN_ENDCASE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDCHECKER) {
		lexer.Pos += len(lexertoken.ENDCHECKER)
		lexer.Emit(lexertoken.TOKEN_ENDCHECKER)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDCLOCKING) {
		lexer.Pos += len(lexertoken.ENDCLOCKING)
		lexer.Emit(lexertoken.TOKEN_ENDCLOCKING)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDCONFIG) {
		lexer.Pos += len(lexertoken.ENDCONFIG)
		lexer.Emit(lexertoken.TOKEN_ENDCONFIG)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDGENERATE) {
		lexer.Pos += len(lexertoken.ENDGENERATE)
		lexer.Emit(lexertoken.TOKEN_ENDGENERATE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDGROUP) {
		lexer.Pos += len(lexertoken.ENDGROUP)
		lexer.Emit(lexertoken.TOKEN_ENDGROUP)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDINTERFACE) {
		lexer.Pos += len(lexertoken.ENDINTERFACE)
		lexer.Emit(lexertoken.TOKEN_ENDINTERFACE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDPRIMITIVE) {
		lexer.Pos += len(lexertoken.ENDPRIMITIVE)
		lexer.Emit(lexertoken.TOKEN_ENDPRIMITIVE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDPROGRAM) {
		lexer.Pos += len(lexertoken.ENDPROGRAM)
		lexer.Emit(lexertoken.TOKEN_ENDPROGRAM)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDPROPERTY) {
		lexer.Pos += len(lexertoken.ENDPROPERTY)
		lexer.Emit(lexertoken.TOKEN_ENDPROPERTY)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDSEQUENCE) {
		lexer.Pos += len(lexertoken.ENDSEQUENCE)
		lexer.Emit(lexertoken.TOKEN_ENDSEQUENCE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDSPECIFY) {
		lexer.Pos += len(lexertoken.ENDSPECIFY)
		lexer.Emit(lexertoken.TOKEN_ENDSPECIFY)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENDTABLE) {
		lexer.Pos += len(lexertoken.ENDTABLE)
		lexer.Emit(lexertoken.TOKEN_ENDTABLE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ENUM) {
		lexer.Pos += len(lexertoken.ENUM)
		lexer.Emit(lexertoken.TOKEN_ENUM)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.EVENT) {
		lexer.Pos += len(lexertoken.EVENT)
		lexer.Emit(lexertoken.TOKEN_EVENT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.EVENTUALLY) {
		lexer.Pos += len(lexertoken.EVENTUALLY)
		lexer.Emit(lexertoken.TOKEN_EVENTUALLY)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.EXPECT) {
		lexer.Pos += len(lexertoken.EXPECT)
		lexer.Emit(lexertoken.TOKEN_EXPECT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.EXPORT) {
		lexer.Pos += len(lexertoken.EXPORT)
		lexer.Emit(lexertoken.TOKEN_EXPORT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.EXTENDS) {
		lexer.Pos += len(lexertoken.EXTENDS)
		lexer.Emit(lexertoken.TOKEN_EXTENDS)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.EXTERN) {
		lexer.Pos += len(lexertoken.EXTERN)
		lexer.Emit(lexertoken.TOKEN_EXTERN)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.FINAL) {
		lexer.Pos += len(lexertoken.FINAL)
		lexer.Emit(lexertoken.TOKEN_FINAL)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.FIRST_MATCH) {
		lexer.Pos += len(lexertoken.FIRST_MATCH)
		lexer.Emit(lexertoken.TOKEN_FIRST_MATCH)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.FOR) {
		lexer.Pos += len(lexertoken.FOR)
		lexer.Emit(lexertoken.TOKEN_FOR)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.FORCE) {
		lexer.Pos += len(lexertoken.FORCE)
		lexer.Emit(lexertoken.TOKEN_FORCE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.FOREACH) {
		lexer.Pos += len(lexertoken.FOREACH)
		lexer.Emit(lexertoken.TOKEN_FOREACH)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.FOREVER) {
		lexer.Pos += len(lexertoken.FOREVER)
		lexer.Emit(lexertoken.TOKEN_FOREVER)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.FORK) {
		lexer.Pos += len(lexertoken.FORK)
		lexer.Emit(lexertoken.TOKEN_FORK)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.FORKJOIN) {
		lexer.Pos += len(lexertoken.FORKJOIN)
		lexer.Emit(lexertoken.TOKEN_FORKJOIN)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.FUNCTION) {
		lexer.Pos += len(lexertoken.FUNCTION)
		lexer.Emit(lexertoken.TOKEN_FUNCTION)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.GENERATE) {
		lexer.Pos += len(lexertoken.GENERATE)
		lexer.Emit(lexertoken.TOKEN_GENERATE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.GENVAR) {
		lexer.Pos += len(lexertoken.GENVAR)
		lexer.Emit(lexertoken.TOKEN_GENVAR)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.GLOBAL) {
		lexer.Pos += len(lexertoken.GLOBAL)
		lexer.Emit(lexertoken.TOKEN_GLOBAL)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.HIGHZ0) {
		lexer.Pos += len(lexertoken.HIGHZ0)
		lexer.Emit(lexertoken.TOKEN_HIGHZ0)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.HIGHZ1) {
		lexer.Pos += len(lexertoken.HIGHZ1)
		lexer.Emit(lexertoken.TOKEN_HIGHZ1)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.IF) {
		lexer.Pos += len(lexertoken.IF)
		lexer.Emit(lexertoken.TOKEN_IF)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.IFF) {
		lexer.Pos += len(lexertoken.IFF)
		lexer.Emit(lexertoken.TOKEN_IFF)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.IFNONE) {
		lexer.Pos += len(lexertoken.IFNONE)
		lexer.Emit(lexertoken.TOKEN_IFNONE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.IGNORE_BINS) {
		lexer.Pos += len(lexertoken.IGNORE_BINS)
		lexer.Emit(lexertoken.TOKEN_IGNORE_BINS)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.ILLEGAL_BINS) {
		lexer.Pos += len(lexertoken.ILLEGAL_BINS)
		lexer.Emit(lexertoken.TOKEN_ILLEGAL_BINS)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.IMPLEMENTS) {
		lexer.Pos += len(lexertoken.IMPLEMENTS)
		lexer.Emit(lexertoken.TOKEN_IMPLEMENTS)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.IMPLIES) {
		lexer.Pos += len(lexertoken.IMPLIES)
		lexer.Emit(lexertoken.TOKEN_IMPLIES)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.IMPORT) {
		lexer.Pos += len(lexertoken.IMPORT)
		lexer.Emit(lexertoken.TOKEN_IMPORT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.INCDIR) {
		lexer.Pos += len(lexertoken.INCDIR)
		lexer.Emit(lexertoken.TOKEN_INCDIR)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.INCLUDE) {
		lexer.Pos += len(lexertoken.INCLUDE)
		lexer.Emit(lexertoken.TOKEN_INCLUDE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.INITIAL) {
		lexer.Pos += len(lexertoken.INITIAL)
		lexer.Emit(lexertoken.TOKEN_INITIAL)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.INOUT) {
		lexer.Pos += len(lexertoken.INOUT)
		lexer.Emit(lexertoken.TOKEN_INOUT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.INPUT) {
		lexer.Pos += len(lexertoken.INPUT)
		lexer.Emit(lexertoken.TOKEN_INPUT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.INSIDE) {
		lexer.Pos += len(lexertoken.INSIDE)
		lexer.Emit(lexertoken.TOKEN_INSIDE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.INSTANCE) {
		lexer.Pos += len(lexertoken.INSTANCE)
		lexer.Emit(lexertoken.TOKEN_INSTANCE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.INT) {
		lexer.Pos += len(lexertoken.INT)
		lexer.Emit(lexertoken.TOKEN_INT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.INTEGER) {
		lexer.Pos += len(lexertoken.INTEGER)
		lexer.Emit(lexertoken.TOKEN_INTEGER)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.INTERCONNECT) {
		lexer.Pos += len(lexertoken.INTERCONNECT)
		lexer.Emit(lexertoken.TOKEN_INTERCONNECT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.INTERFACE) {
		lexer.Pos += len(lexertoken.INTERFACE)
		lexer.Emit(lexertoken.TOKEN_INTERFACE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.INTERSECT) {
		lexer.Pos += len(lexertoken.INTERSECT)
		lexer.Emit(lexertoken.TOKEN_INTERSECT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.JOIN) {
		lexer.Pos += len(lexertoken.JOIN)
		lexer.Emit(lexertoken.TOKEN_JOIN)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.JOIN_ANY) {
		lexer.Pos += len(lexertoken.JOIN_ANY)
		lexer.Emit(lexertoken.TOKEN_JOIN_ANY)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.JOIN_NONE) {
		lexer.Pos += len(lexertoken.JOIN_NONE)
		lexer.Emit(lexertoken.TOKEN_JOIN_NONE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LARGE) {
		lexer.Pos += len(lexertoken.LARGE)
		lexer.Emit(lexertoken.TOKEN_LARGE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LET) {
		lexer.Pos += len(lexertoken.LET)
		lexer.Emit(lexertoken.TOKEN_LET)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LIBLIST) {
		lexer.Pos += len(lexertoken.LIBLIST)
		lexer.Emit(lexertoken.TOKEN_LIBLIST)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LIBRARY) {
		lexer.Pos += len(lexertoken.LIBRARY)
		lexer.Emit(lexertoken.TOKEN_LIBRARY)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LOCAL) {
		lexer.Pos += len(lexertoken.LOCAL)
		lexer.Emit(lexertoken.TOKEN_LOCAL)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LOCALPARAM) {
		lexer.Pos += len(lexertoken.LOCALPARAM)
		lexer.Emit(lexertoken.TOKEN_LOCALPARAM)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LOGIC) {
		lexer.Pos += len(lexertoken.LOGIC)
		lexer.Emit(lexertoken.TOKEN_LOGIC)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.LONGINT) {
		lexer.Pos += len(lexertoken.LONGINT)
		lexer.Emit(lexertoken.TOKEN_LONGINT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.MACROMODULE) {
		lexer.Pos += len(lexertoken.MACROMODULE)
		lexer.Emit(lexertoken.TOKEN_MACROMODULE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.MATCHES) {
		lexer.Pos += len(lexertoken.MATCHES)
		lexer.Emit(lexertoken.TOKEN_MATCHES)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.MEDIUM) {
		lexer.Pos += len(lexertoken.MEDIUM)
		lexer.Emit(lexertoken.TOKEN_MEDIUM)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.MODPORT) {
		lexer.Pos += len(lexertoken.MODPORT)
		lexer.Emit(lexertoken.TOKEN_MODPORT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.MODULE) {
		lexer.Pos += len(lexertoken.MODULE)
		lexer.Emit(lexertoken.TOKEN_MODULE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.NAND) {
		lexer.Pos += len(lexertoken.NAND)
		lexer.Emit(lexertoken.TOKEN_NAND)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.NEGEDGE) {
		lexer.Pos += len(lexertoken.NEGEDGE)
		lexer.Emit(lexertoken.TOKEN_NEGEDGE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.NETTYPE) {
		lexer.Pos += len(lexertoken.NETTYPE)
		lexer.Emit(lexertoken.TOKEN_NETTYPE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.NEW) {
		lexer.Pos += len(lexertoken.NEW)
		lexer.Emit(lexertoken.TOKEN_NEW)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.NEXTTIME) {
		lexer.Pos += len(lexertoken.NEXTTIME)
		lexer.Emit(lexertoken.TOKEN_NEXTTIME)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.NMOS) {
		lexer.Pos += len(lexertoken.NMOS)
		lexer.Emit(lexertoken.TOKEN_NMOS)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.NOR) {
		lexer.Pos += len(lexertoken.NOR)
		lexer.Emit(lexertoken.TOKEN_NOR)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.NOSHOWCANCELLED) {
		lexer.Pos += len(lexertoken.NOSHOWCANCELLED)
		lexer.Emit(lexertoken.TOKEN_NOSHOWCANCELLED)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.NOT) {
		lexer.Pos += len(lexertoken.NOT)
		lexer.Emit(lexertoken.TOKEN_NOT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.NOTIF0) {
		lexer.Pos += len(lexertoken.NOTIF0)
		lexer.Emit(lexertoken.TOKEN_NOTIF0)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.NOTIF1) {
		lexer.Pos += len(lexertoken.NOTIF1)
		lexer.Emit(lexertoken.TOKEN_NOTIF1)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.NULL) {
		lexer.Pos += len(lexertoken.NULL)
		lexer.Emit(lexertoken.TOKEN_NULL)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.OR) {
		lexer.Pos += len(lexertoken.OR)
		lexer.Emit(lexertoken.TOKEN_OR)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.OUTPUT) {
		lexer.Pos += len(lexertoken.OUTPUT)
		lexer.Emit(lexertoken.TOKEN_OUTPUT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.PACKED) {
		lexer.Pos += len(lexertoken.PACKED)
		lexer.Emit(lexertoken.TOKEN_PACKED)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.PARAMETER) {
		lexer.Pos += len(lexertoken.PARAMETER)
		lexer.Emit(lexertoken.TOKEN_PARAMETER)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.PMOS) {
		lexer.Pos += len(lexertoken.PMOS)
		lexer.Emit(lexertoken.TOKEN_PMOS)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.POSEDGE) {
		lexer.Pos += len(lexertoken.POSEDGE)
		lexer.Emit(lexertoken.TOKEN_POSEDGE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.PRIMITIVE) {
		lexer.Pos += len(lexertoken.PRIMITIVE)
		lexer.Emit(lexertoken.TOKEN_PRIMITIVE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.PRIORITY) {
		lexer.Pos += len(lexertoken.PRIORITY)
		lexer.Emit(lexertoken.TOKEN_PRIORITY)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.PROGRAM) {
		lexer.Pos += len(lexertoken.PROGRAM)
		lexer.Emit(lexertoken.TOKEN_PROGRAM)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.PROPERTY) {
		lexer.Pos += len(lexertoken.PROPERTY)
		lexer.Emit(lexertoken.TOKEN_PROPERTY)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.PROTECTED) {
		lexer.Pos += len(lexertoken.PROTECTED)
		lexer.Emit(lexertoken.TOKEN_PROTECTED)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.PULL0) {
		lexer.Pos += len(lexertoken.PULL0)
		lexer.Emit(lexertoken.TOKEN_PULL0)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.PULL1) {
		lexer.Pos += len(lexertoken.PULL1)
		lexer.Emit(lexertoken.TOKEN_PULL1)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.PULLDOWN) {
		lexer.Pos += len(lexertoken.PULLDOWN)
		lexer.Emit(lexertoken.TOKEN_PULLDOWN)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.PULLUP) {
		lexer.Pos += len(lexertoken.PULLUP)
		lexer.Emit(lexertoken.TOKEN_PULLUP)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.PULSESTYLE_ONDETECT) {
		lexer.Pos += len(lexertoken.PULSESTYLE_ONDETECT)
		lexer.Emit(lexertoken.TOKEN_PULSESTYLE_ONDETECT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.PULSESTYLE_ONEVENT) {
		lexer.Pos += len(lexertoken.PULSESTYLE_ONEVENT)
		lexer.Emit(lexertoken.TOKEN_PULSESTYLE_ONEVENT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.PURE) {
		lexer.Pos += len(lexertoken.PURE)
		lexer.Emit(lexertoken.TOKEN_PURE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.RAND) {
		lexer.Pos += len(lexertoken.RAND)
		lexer.Emit(lexertoken.TOKEN_RAND)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.RANDC) {
		lexer.Pos += len(lexertoken.RANDC)
		lexer.Emit(lexertoken.TOKEN_RANDC)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.RANDCASE) {
		lexer.Pos += len(lexertoken.RANDCASE)
		lexer.Emit(lexertoken.TOKEN_RANDCASE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.RANDSEQUENCE) {
		lexer.Pos += len(lexertoken.RANDSEQUENCE)
		lexer.Emit(lexertoken.TOKEN_RANDSEQUENCE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.RCMOS) {
		lexer.Pos += len(lexertoken.RCMOS)
		lexer.Emit(lexertoken.TOKEN_RCMOS)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.REAL) {
		lexer.Pos += len(lexertoken.REAL)
		lexer.Emit(lexertoken.TOKEN_REAL)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.REALTIME) {
		lexer.Pos += len(lexertoken.REALTIME)
		lexer.Emit(lexertoken.TOKEN_REALTIME)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.REF) {
		lexer.Pos += len(lexertoken.REF)
		lexer.Emit(lexertoken.TOKEN_REF)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.REG) {
		lexer.Pos += len(lexertoken.REG)
		lexer.Emit(lexertoken.TOKEN_REG)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.REJECT_ON) {
		lexer.Pos += len(lexertoken.REJECT_ON)
		lexer.Emit(lexertoken.TOKEN_REJECT_ON)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.RELEASE) {
		lexer.Pos += len(lexertoken.RELEASE)
		lexer.Emit(lexertoken.TOKEN_RELEASE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.REPEAT) {
		lexer.Pos += len(lexertoken.REPEAT)
		lexer.Emit(lexertoken.TOKEN_REPEAT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.RESTRICT) {
		lexer.Pos += len(lexertoken.RESTRICT)
		lexer.Emit(lexertoken.TOKEN_RESTRICT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.RETURN) {
		lexer.Pos += len(lexertoken.RETURN)
		lexer.Emit(lexertoken.TOKEN_RETURN)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.RNMOS) {
		lexer.Pos += len(lexertoken.RNMOS)
		lexer.Emit(lexertoken.TOKEN_RNMOS)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.RPMOS) {
		lexer.Pos += len(lexertoken.RPMOS)
		lexer.Emit(lexertoken.TOKEN_RPMOS)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.RTRAN) {
		lexer.Pos += len(lexertoken.RTRAN)
		lexer.Emit(lexertoken.TOKEN_RTRAN)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.RTRANIF0) {
		lexer.Pos += len(lexertoken.RTRANIF0)
		lexer.Emit(lexertoken.TOKEN_RTRANIF0)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.RTRANIF1) {
		lexer.Pos += len(lexertoken.RTRANIF1)
		lexer.Emit(lexertoken.TOKEN_RTRANIF1)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.S_ALWAYS) {
		lexer.Pos += len(lexertoken.S_ALWAYS)
		lexer.Emit(lexertoken.TOKEN_S_ALWAYS)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.S_EVENTUALLY) {
		lexer.Pos += len(lexertoken.S_EVENTUALLY)
		lexer.Emit(lexertoken.TOKEN_S_EVENTUALLY)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.S_NEXTTIME) {
		lexer.Pos += len(lexertoken.S_NEXTTIME)
		lexer.Emit(lexertoken.TOKEN_S_NEXTTIME)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.S_UNTIL) {
		lexer.Pos += len(lexertoken.S_UNTIL)
		lexer.Emit(lexertoken.TOKEN_S_UNTIL)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.S_UNTIL_WITH) {
		lexer.Pos += len(lexertoken.S_UNTIL_WITH)
		lexer.Emit(lexertoken.TOKEN_S_UNTIL_WITH)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.SCALARED) {
		lexer.Pos += len(lexertoken.SCALARED)
		lexer.Emit(lexertoken.TOKEN_SCALARED)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.SEQUENCE) {
		lexer.Pos += len(lexertoken.SEQUENCE)
		lexer.Emit(lexertoken.TOKEN_SEQUENCE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.SHORTINT) {
		lexer.Pos += len(lexertoken.SHORTINT)
		lexer.Emit(lexertoken.TOKEN_SHORTINT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.SHORTREAL) {
		lexer.Pos += len(lexertoken.SHORTREAL)
		lexer.Emit(lexertoken.TOKEN_SHORTREAL)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.SHOWCANCELLED) {
		lexer.Pos += len(lexertoken.SHOWCANCELLED)
		lexer.Emit(lexertoken.TOKEN_SHOWCANCELLED)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.SIGNED) {
		lexer.Pos += len(lexertoken.SIGNED)
		lexer.Emit(lexertoken.TOKEN_SIGNED)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.SMALL) {
		lexer.Pos += len(lexertoken.SMALL)
		lexer.Emit(lexertoken.TOKEN_SMALL)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.SOFT) {
		lexer.Pos += len(lexertoken.SOFT)
		lexer.Emit(lexertoken.TOKEN_SOFT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.SOLVE) {
		lexer.Pos += len(lexertoken.SOLVE)
		lexer.Emit(lexertoken.TOKEN_SOLVE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.SPECIFY) {
		lexer.Pos += len(lexertoken.SPECIFY)
		lexer.Emit(lexertoken.TOKEN_SPECIFY)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.SPECPARAM) {
		lexer.Pos += len(lexertoken.SPECPARAM)
		lexer.Emit(lexertoken.TOKEN_SPECPARAM)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.STATIC) {
		lexer.Pos += len(lexertoken.STATIC)
		lexer.Emit(lexertoken.TOKEN_STATIC)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.STRING) {
		lexer.Pos += len(lexertoken.STRING)
		lexer.Emit(lexertoken.TOKEN_STRING)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.STRONG) {
		lexer.Pos += len(lexertoken.STRONG)
		lexer.Emit(lexertoken.TOKEN_STRONG)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.STRONG0) {
		lexer.Pos += len(lexertoken.STRONG0)
		lexer.Emit(lexertoken.TOKEN_STRONG0)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.STRONG1) {
		lexer.Pos += len(lexertoken.STRONG1)
		lexer.Emit(lexertoken.TOKEN_STRONG1)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.STRUCT) {
		lexer.Pos += len(lexertoken.STRUCT)
		lexer.Emit(lexertoken.TOKEN_STRUCT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.SUPER) {
		lexer.Pos += len(lexertoken.SUPER)
		lexer.Emit(lexertoken.TOKEN_SUPER)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.SUPPLY0) {
		lexer.Pos += len(lexertoken.SUPPLY0)
		lexer.Emit(lexertoken.TOKEN_SUPPLY0)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.SUPPLY1) {
		lexer.Pos += len(lexertoken.SUPPLY1)
		lexer.Emit(lexertoken.TOKEN_SUPPLY1)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.SYNC_ACCEPT_ON) {
		lexer.Pos += len(lexertoken.SYNC_ACCEPT_ON)
		lexer.Emit(lexertoken.TOKEN_SYNC_ACCEPT_ON)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.SYNC_REJECT_ON) {
		lexer.Pos += len(lexertoken.SYNC_REJECT_ON)
		lexer.Emit(lexertoken.TOKEN_SYNC_REJECT_ON)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.TABLE) {
		lexer.Pos += len(lexertoken.TABLE)
		lexer.Emit(lexertoken.TOKEN_TABLE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.TAGGED) {
		lexer.Pos += len(lexertoken.TAGGED)
		lexer.Emit(lexertoken.TOKEN_TAGGED)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.THIS) {
		lexer.Pos += len(lexertoken.THIS)
		lexer.Emit(lexertoken.TOKEN_THIS)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.THROUGHOUT) {
		lexer.Pos += len(lexertoken.THROUGHOUT)
		lexer.Emit(lexertoken.TOKEN_THROUGHOUT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.TIME) {
		lexer.Pos += len(lexertoken.TIME)
		lexer.Emit(lexertoken.TOKEN_TIME)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.TIMEPRECISION) {
		lexer.Pos += len(lexertoken.TIMEPRECISION)
		lexer.Emit(lexertoken.TOKEN_TIMEPRECISION)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.TIMEUNIT) {
		lexer.Pos += len(lexertoken.TIMEUNIT)
		lexer.Emit(lexertoken.TOKEN_TIMEUNIT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.TRAN) {
		lexer.Pos += len(lexertoken.TRAN)
		lexer.Emit(lexertoken.TOKEN_TRAN)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.TRANIF0) {
		lexer.Pos += len(lexertoken.TRANIF0)
		lexer.Emit(lexertoken.TOKEN_TRANIF0)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.TRANIF1) {
		lexer.Pos += len(lexertoken.TRANIF1)
		lexer.Emit(lexertoken.TOKEN_TRANIF1)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.TRI) {
		lexer.Pos += len(lexertoken.TRI)
		lexer.Emit(lexertoken.TOKEN_TRI)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.TRI0) {
		lexer.Pos += len(lexertoken.TRI0)
		lexer.Emit(lexertoken.TOKEN_TRI0)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.TRI1) {
		lexer.Pos += len(lexertoken.TRI1)
		lexer.Emit(lexertoken.TOKEN_TRI1)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.TRIAND) {
		lexer.Pos += len(lexertoken.TRIAND)
		lexer.Emit(lexertoken.TOKEN_TRIAND)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.TRIOR) {
		lexer.Pos += len(lexertoken.TRIOR)
		lexer.Emit(lexertoken.TOKEN_TRIOR)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.TRIREG) {
		lexer.Pos += len(lexertoken.TRIREG)
		lexer.Emit(lexertoken.TOKEN_TRIREG)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.TYPE) {
		lexer.Pos += len(lexertoken.TYPE)
		lexer.Emit(lexertoken.TOKEN_TYPE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.TYPEDEF) {
		lexer.Pos += len(lexertoken.TYPEDEF)
		lexer.Emit(lexertoken.TOKEN_TYPEDEF)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.UNION) {
		lexer.Pos += len(lexertoken.UNION)
		lexer.Emit(lexertoken.TOKEN_UNION)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.UNIQUE) {
		lexer.Pos += len(lexertoken.UNIQUE)
		lexer.Emit(lexertoken.TOKEN_UNIQUE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.UNIQUE0) {
		lexer.Pos += len(lexertoken.UNIQUE0)
		lexer.Emit(lexertoken.TOKEN_UNIQUE0)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.UNSIGNED) {
		lexer.Pos += len(lexertoken.UNSIGNED)
		lexer.Emit(lexertoken.TOKEN_UNSIGNED)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.UNTIL) {
		lexer.Pos += len(lexertoken.UNTIL)
		lexer.Emit(lexertoken.TOKEN_UNTIL)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.UNTIL_WITH) {
		lexer.Pos += len(lexertoken.UNTIL_WITH)
		lexer.Emit(lexertoken.TOKEN_UNTIL_WITH)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.UNTYPED) {
		lexer.Pos += len(lexertoken.UNTYPED)
		lexer.Emit(lexertoken.TOKEN_UNTYPED)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.USE) {
		lexer.Pos += len(lexertoken.USE)
		lexer.Emit(lexertoken.TOKEN_USE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.UWIRE) {
		lexer.Pos += len(lexertoken.UWIRE)
		lexer.Emit(lexertoken.TOKEN_UWIRE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.VAR) {
		lexer.Pos += len(lexertoken.VAR)
		lexer.Emit(lexertoken.TOKEN_VAR)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.VECTORED) {
		lexer.Pos += len(lexertoken.VECTORED)
		lexer.Emit(lexertoken.TOKEN_VECTORED)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.VIRTUAL) {
		lexer.Pos += len(lexertoken.VIRTUAL)
		lexer.Emit(lexertoken.TOKEN_VIRTUAL)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.VOID) {
		lexer.Pos += len(lexertoken.VOID)
		lexer.Emit(lexertoken.TOKEN_VOID)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.WAIT) {
		lexer.Pos += len(lexertoken.WAIT)
		lexer.Emit(lexertoken.TOKEN_WAIT)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.WAIT_ORDER) {
		lexer.Pos += len(lexertoken.WAIT_ORDER)
		lexer.Emit(lexertoken.TOKEN_WAIT_ORDER)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.WAND) {
		lexer.Pos += len(lexertoken.WAND)
		lexer.Emit(lexertoken.TOKEN_WAND)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.WEAK) {
		lexer.Pos += len(lexertoken.WEAK)
		lexer.Emit(lexertoken.TOKEN_WEAK)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.WEAK0) {
		lexer.Pos += len(lexertoken.WEAK0)
		lexer.Emit(lexertoken.TOKEN_WEAK0)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.WEAK1) {
		lexer.Pos += len(lexertoken.WEAK1)
		lexer.Emit(lexertoken.TOKEN_WEAK1)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.WHILE) {
		lexer.Pos += len(lexertoken.WHILE)
		lexer.Emit(lexertoken.TOKEN_WHILE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.WILDCARD) {
		lexer.Pos += len(lexertoken.WILDCARD)
		lexer.Emit(lexertoken.TOKEN_WILDCARD)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.WIRE) {
		lexer.Pos += len(lexertoken.WIRE)
		lexer.Emit(lexertoken.TOKEN_WIRE)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.WITH) {
		lexer.Pos += len(lexertoken.WITH)
		lexer.Emit(lexertoken.TOKEN_WITH)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.WITHIN) {
		lexer.Pos += len(lexertoken.WITHIN)
		lexer.Emit(lexertoken.TOKEN_WITHIN)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.WOR) {
		lexer.Pos += len(lexertoken.WOR)
		lexer.Emit(lexertoken.TOKEN_WOR)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.XNOR) {
		lexer.Pos += len(lexertoken.XNOR)
		lexer.Emit(lexertoken.TOKEN_XNOR)
		return LexKeyword
	} else if strings.HasPrefix(lexer.InputToEnd(), lexertoken.XOR) {
		lexer.Pos += len(lexertoken.XOR)
		lexer.Emit(lexertoken.TOKEN_XOR)
		return LexKeyword

	} else {
		return LexStatement
		// return lexer.Errorf(errors.LEXER_ERROR_UNPARSABLE_STATEMENT)
	}
}
