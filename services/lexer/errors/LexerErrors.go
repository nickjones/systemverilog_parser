package errors

const (
	LEXER_ERROR_MISSING_RIGHT_BRACKET string = "Missing a closing section bracket"
	LEXER_ERROR_UNPARSABLE_STATEMENT  string = "Unparsable statement"
	LEXER_ERROR_UNEXPECTED_EOF        string = "Unexpected end of file.  Missing a semicolon?"
	LEXER_ERROR_EXPECTED_STRING       string = "Expected a string following the token"
)
