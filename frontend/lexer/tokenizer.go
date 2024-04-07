package lexer

import (
	"fmt"
	"regexp"
)
type RegexpHandler func (lex *lexer, regex *regexp.Regexp)
type RegexPatterns struct {
	pattern *regexp.Regexp
	handler RegexpHandler
}
type Lexer struct {
	source string
	position int 
	tokens []Token
	patterns []RegexPatterns
}