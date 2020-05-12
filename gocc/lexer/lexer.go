// Code generated by gocc; DO NOT EDIT.

package lexer

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/Loptt/lambdish-compiler/gocc/token"
)

const (
	NoState    = -1
	NumStates  = 47
	NumSymbols = 57
)

type Lexer struct {
	src    []byte
	pos    int
	line   int
	column int
}

func NewLexer(src []byte) *Lexer {
	lexer := &Lexer{
		src:    src,
		pos:    0,
		line:   1,
		column: 1,
	}
	return lexer
}

func NewLexerFile(fpath string) (*Lexer, error) {
	src, err := ioutil.ReadFile(fpath)
	if err != nil {
		return nil, err
	}
	return NewLexer(src), nil
}

func (l *Lexer) Scan() (tok *token.Token) {
	tok = new(token.Token)
	if l.pos >= len(l.src) {
		tok.Type = token.EOF
		tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = l.pos, l.line, l.column
		return
	}
	start, startLine, startColumn, end := l.pos, l.line, l.column, 0
	tok.Type = token.INVALID
	state, rune1, size := 0, rune(-1), 0
	for state != -1 {
		if l.pos >= len(l.src) {
			rune1 = -1
		} else {
			rune1, size = utf8.DecodeRune(l.src[l.pos:])
			l.pos += size
		}

		nextState := -1
		if rune1 != -1 {
			nextState = TransTab[state](rune1)
		}
		state = nextState

		if state != -1 {

			switch rune1 {
			case '\n':
				l.line++
				l.column = 1
			case '\r':
				l.column = 1
			case '\t':
				l.column += 4
			default:
				l.column++
			}

			switch {
			case ActTab[state].Accept != -1:
				tok.Type = ActTab[state].Accept
				end = l.pos
			case ActTab[state].Ignore != "":
				start, startLine, startColumn = l.pos, l.line, l.column
				state = 0
				if start >= len(l.src) {
					tok.Type = token.EOF
				}

			}
		} else {
			if tok.Type == token.INVALID {
				end = l.pos
			}
		}
	}
	if end > start {
		l.pos = end
		tok.Lit = l.src[start:end]
	} else {
		tok.Lit = []byte{}
	}
	tok.Pos.Offset, tok.Pos.Line, tok.Pos.Column = start, startLine, startColumn

	return
}

func (l *Lexer) Reset() {
	l.pos = 0
}

/*
Lexer symbols:
0: '+'
1: '-'
2: '*'
3: '/'
4: '%'
5: '<'
6: '>'
7: '!'
8: '-'
9: '''
10: '''
11: 'f'
12: 'u'
13: 'n'
14: 'c'
15: ':'
16: ':'
17: '='
18: '>'
19: '('
20: ')'
21: ','
22: '['
23: ']'
24: 'n'
25: 'u'
26: 'm'
27: 'b'
28: 'o'
29: 'o'
30: 'l'
31: 'c'
32: 'h'
33: 'a'
34: 'r'
35: '#'
36: '.'
37: '"'
38: ' '
39: '"'
40: 't'
41: 'r'
42: 'u'
43: 'e'
44: 'f'
45: 'a'
46: 'l'
47: 's'
48: 'e'
49: ' '
50: '\t'
51: '\n'
52: '\r'
53: '0'-'9'
54: 'a'-'z'
55: 'A'-'Z'
56: .
*/
