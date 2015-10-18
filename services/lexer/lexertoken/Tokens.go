package lexertoken

const EOF rune = 0

const LEFT_BRACKET string = "["
const RIGHT_BRACKET string = "]"
const EQUAL_SIGN string = "="
const NEWLINE string = "\n"
const LEFT_PAREN string = "("
const RIGHT_PAREN string = ")"
const NONBLOCKING_ASSIGNMENT string = "<="
const IFDEF string = "ifdef"
const IFNDEF string = "ifndef"
const DEFINE string = "define"
const ENDIFDEF string = "endif"
const ELSIF string = "elsif" // Used for defines
const DOUBLE_QUOTE string = "\""
const SEMICOLON string = ";"
const SINGLE_LINE_COMMENT string = "//"
const MULTI_LINE_COMMENT_START string = "/*"
const MULTI_LINE_COMMENT_STOP string = "*/"

// Reserved keywords
const ACCEPT_ON string = "accept_on"
const ALIAS string = "alias"
const ALWAYS string = "always"
const ALWAYS_COMB string = "always_comb"
const ALWAYS_FF string = "always_ff"
const ALWAYS_LATCH string = "always_latch"
const AND string = "and"
const ASSERT string = "assert"
const ASSIGN string = "assign"
const ASSUME string = "assume"
const AUTOMATIC string = "automatic"
const BEFORE string = "before"
const BEGIN string = "begin"
const BIND string = "bind"
const BINS string = "bins"
const BINSOF string = "binsof"
const BIT string = "bit"
const BREAK string = "break"
const BUF string = "buf"
const BUFIF0 string = "bufif0"
const BUFIF1 string = "bufif1"
const BYTE string = "byte"
const CASE string = "case"
const CASEX string = "casex"
const CASEZ string = "casez"
const CELL string = "cell"
const CHANDLE string = "chandle"
const CHECKER string = "checker"
const CLASS string = "class"
const CLOCKING string = "clocking"
const CMOS string = "cmos"
const CONFIG string = "config"
const CONST string = "const"
const CONSTRAINT string = "constraint"
const CONTEXT string = "context"
const CONTINUE string = "continue"
const COVER string = "cover"
const COVERGROUP string = "covergroup"
const COVERPOINT string = "coverpoint"
const CROSS string = "cross"
const DEASSIGN string = "deassign"
const DEFAULT string = "default"
const DEFPARAM string = "defparam"
const DESIGN string = "design"
const DISABLE string = "disable"
const DIST string = "dist"
const DO string = "do"
const EDGE string = "edge"
const ELSE string = "else"
const END string = "end"
const ENDCASE string = "endcase"
const ENDCHECKER string = "endchecker"
const ENDCLASS string = "endclass"
const ENDCLOCKING string = "endclocking"
const ENDCONFIG string = "endconfig"
const ENDFUNCTION string = "endfunction"
const ENDGENERATE string = "endgenerate"
const ENDGROUP string = "endgroup"
const ENDINTERFACE string = "endinterface"
const ENDMODULE string = "endmodule"
const ENDPACKAGE string = "endpackage"
const ENDPRIMITIVE string = "endprimitive"
const ENDPROGRAM string = "endprogram"
const ENDPROPERTY string = "endproperty"
const ENDSEQUENCE string = "endsequence"
const ENDSPECIFY string = "endspecify"
const ENDTABLE string = "endtable"
const ENDTASK string = "endtask"
const ENUM string = "enum"
const EVENT string = "event"
const EVENTUALLY string = "eventually"
const EXPECT string = "expect"
const EXPORT string = "export"
const EXTENDS string = "extends"
const EXTERN string = "extern"
const FINAL string = "final"
const FIRST_MATCH string = "first_match"
const FOR string = "for"
const FORCE string = "force"
const FOREACH string = "foreach"
const FOREVER string = "forever"
const FORK string = "fork"
const FORKJOIN string = "forkjoin"
const FUNCTION string = "function"
const GENERATE string = "generate"
const GENVAR string = "genvar"
const GLOBAL string = "global"
const HIGHZ0 string = "highz0"
const HIGHZ1 string = "highz1"
const IF string = "if"
const IFF string = "iff"
const IFNONE string = "ifnone"
const IGNORE_BINS string = "ignore_bins"
const ILLEGAL_BINS string = "illegal_bins"
const IMPLEMENTS string = "implements"
const IMPLIES string = "implies"
const IMPORT string = "import"
const INCDIR string = "incdir"
const INCLUDE string = "include"
const INITIAL string = "initial"
const INOUT string = "inout"
const INPUT string = "input"
const INSIDE string = "inside"
const INSTANCE string = "instance"
const INT string = "int"
const INTEGER string = "integer"
const INTERCONNECT string = "interconnect"
const INTERFACE string = "interface"
const INTERSECT string = "intersect"
const JOIN string = "join"
const JOIN_ANY string = "join_any"
const JOIN_NONE string = "join_none"
const LARGE string = "large"
const LET string = "let"
const LIBLIST string = "liblist"
const LIBRARY string = "library"
const LOCAL string = "local"
const LOCALPARAM string = "localparam"
const LOGIC string = "logic"
const LONGINT string = "longint"
const MACROMODULE string = "macromodule"
const MATCHES string = "matches"
const MEDIUM string = "medium"
const MODPORT string = "modport"
const MODULE string = "module"
const NAND string = "nand"
const NEGEDGE string = "negedge"
const NETTYPE string = "nettype"
const NEW string = "new"
const NEXTTIME string = "nexttime"
const NMOS string = "nmos"
const NOR string = "nor"
const NOSHOWCANCELLED string = "noshowcancelled"
const NOT string = "not"
const NOTIF0 string = "notif0"
const NOTIF1 string = "notif1"
const NULL string = "null"
const OR string = "or"
const OUTPUT string = "output"
const PACKAGE string = "package"
const PACKED string = "packed"
const PARAMETER string = "parameter"
const PMOS string = "pmos"
const POSEDGE string = "posedge"
const PRIMITIVE string = "primitive"
const PRIORITY string = "priority"
const PROGRAM string = "program"
const PROPERTY string = "property"
const PROTECTED string = "protected"
const PULL0 string = "pull0"
const PULL1 string = "pull1"
const PULLDOWN string = "pulldown"
const PULLUP string = "pullup"
const PULSESTYLE_ONDETECT string = "pulsestyle_ondetect"
const PULSESTYLE_ONEVENT string = "pulsestyle_onevent"
const PURE string = "pure"
const RAND string = "rand"
const RANDC string = "randc"
const RANDCASE string = "randcase"
const RANDSEQUENCE string = "randsequence"
const RCMOS string = "rcmos"
const REAL string = "real"
const REALTIME string = "realtime"
const REF string = "ref"
const REG string = "reg"
const REJECT_ON string = "reject_on"
const RELEASE string = "release"
const REPEAT string = "repeat"
const RESTRICT string = "restrict"
const RETURN string = "return"
const RNMOS string = "rnmos"
const RPMOS string = "rpmos"
const RTRAN string = "rtran"
const RTRANIF0 string = "rtranif0"
const RTRANIF1 string = "rtranif1"
const SCALARED string = "scalared"
const SEQUENCE string = "sequence"
const SHORTINT string = "shortint"
const SHORTREAL string = "shortreal"
const SHOWCANCELLED string = "showcancelled"
const SIGNED string = "signed"
const SMALL string = "small"
const SOFT string = "soft"
const SOLVE string = "solve"
const SPECIFY string = "specify"
const SPECPARAM string = "specparam"
const STATIC string = "static"
const STRING string = "string"
const STRONG string = "strong"
const STRONG0 string = "strong0"
const STRONG1 string = "strong1"
const STRUCT string = "struct"
const SUPER string = "super"
const SUPPLY0 string = "supply0"
const SUPPLY1 string = "supply1"
const SYNC_ACCEPT_ON string = "sync_accept_on"
const SYNC_REJECT_ON string = "sync_reject_on"
const S_ALWAYS string = "s_always"
const S_EVENTUALLY string = "s_eventually"
const S_NEXTTIME string = "s_nexttime"
const S_UNTIL string = "s_until"
const S_UNTIL_WITH string = "s_until_with"
const TABLE string = "table"
const TAGGED string = "tagged"
const TASK string = "task"
const THIS string = "this"
const THROUGHOUT string = "throughout"
const TIME string = "time"
const TIMEPRECISION string = "timeprecision"
const TIMEUNIT string = "timeunit"
const TRAN string = "tran"
const TRANIF0 string = "tranif0"
const TRANIF1 string = "tranif1"
const TRI string = "tri"
const TRI0 string = "tri0"
const TRI1 string = "tri1"
const TRIAND string = "triand"
const TRIOR string = "trior"
const TRIREG string = "trireg"
const TYPE string = "type"
const TYPEDEF string = "typedef"
const UNION string = "union"
const UNIQUE string = "unique"
const UNIQUE0 string = "unique0"
const UNSIGNED string = "unsigned"
const UNTIL string = "until"
const UNTIL_WITH string = "until_with"
const UNTYPED string = "untyped"
const USE string = "use"
const UWIRE string = "uwire"
const VAR string = "var"
const VECTORED string = "vectored"
const VIRTUAL string = "virtual"
const VOID string = "void"
const WAIT string = "wait"
const WAIT_ORDER string = "wait_order"
const WAND string = "wand"
const WEAK string = "weak"
const WEAK0 string = "weak0"
const WEAK1 string = "weak1"
const WHILE string = "while"
const WILDCARD string = "wildcard"
const WIRE string = "wire"
const WITH string = "with"
const WITHIN string = "within"
const WOR string = "wor"
const XNOR string = "xnor"
const XOR string = "xor"

// Operators
var ASSIGNMENT_OPERATOR []string = []string{"=", "+=", "-=", "*=", "/=", "%=", "&=", "|=", "^=", "<<=", ">>=", "<<<=", ">>>="}
var UNARY_OPERATOR []string = []string{"+", "-", "!", "~", "&", "~&", "|", "~|", "^", "~^", "^~"}
var BINARY_OPERATOR []string = []string{"+", "-", "*", "/", "%", "==", "!=", "===", "!==", "==?", "!=?", "&&", "||", "**", "<", "<=", ">", ">=", "&", "|", "^", "^~", "~^", ">>", "<<", ">>>", "<<<", "->", "<->"}
var INC_DEC_OPERATOR []string = []string{"++", "--"}
var STREAM_OPERATOR []string = []string{">>", "<<"}
