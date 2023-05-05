package src

import (
	"strconv"
)

type lexer struct {
	input   string
	pos     int
	current byte
}

func (l *lexer) advance() {
	l.pos++
	if l.pos < len(l.input) {
		l.current = l.input[l.pos]
	} else {
		l.current = 0
	}
}

func (l *lexer) skipWhitespace() {
	for l.current == ' ' || l.current == '\t' || l.current == '\n' || l.current == '\r' {
		l.advance()
	}
}

func (l *lexer) parseNumber() int {
	var numStr string
	for l.current >= '0' && l.current <= '9' {
		numStr += string(l.current)
		l.advance()
	}
	num, _ := strconv.Atoi(numStr)
	return num
}

func (l *lexer) parseFactor() int {
	l.skipWhitespace()
	if l.current >= '0' && l.current <= '9' {
		return l.parseNumber()
	} else if l.current == '(' {
		l.advance()
		result := l.parseExpr()
		l.skipWhitespace()
		if l.current != ')' {
			panic("Expected ')' but found " + string(l.current))
		}
		l.advance()
		return result
	} else if l.current == '-' {
		l.advance()
		return -l.parseFactor()
	}
	panic("Unexpected character: " + string(l.current))
}

func (l *lexer) parseTerm() int {
	result := l.parseFactor()
	for l.current == '*' || l.current == '/' {
		op := l.current
		l.advance()
		l.skipWhitespace()
		num := l.parseFactor()
		if op == '*' {
			result *= num
		} else {
			result /= num
		}
	}
	return result
}

func (l *lexer) parseExpr() int {
	result := l.parseTerm()
	for l.current == '+' || l.current == '-' {
		op := l.current
		l.advance()
		l.skipWhitespace()
		num := l.parseTerm()
		if op == '+' {
			result += num
		} else {
			result -= num
		}
	}
	return result
}

func evaluate(expr string) int {
	l := lexer{input: expr, pos: -1}
	l.advance()
	return l.parseExpr()
}

// func main() {
// 	fmt.Println(evaluate("3 + 4 * 5"))   // Output: 23
// 	fmt.Println(evaluate("(3 + 4) * 5")) // Output: 35
// 	fmt.Println(evaluate("2 * (3 + 4)")) // Output: 14
// }
