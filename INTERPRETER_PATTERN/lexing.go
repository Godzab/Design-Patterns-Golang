package main

import (
	"fmt"
	"strings"
	"unicode"
)

type TokenType int

const(
	Int TokenType = iota
	Plus
	Minus
	Lparen
	Rparen
)

type Token struct{
	Type TokenType
	Text string
}

func (t Token) String() string {
	return fmt.Sprintf("`%s`", t.Text)
}

func Lex(input string)[]Token{
	var res []Token

	for i:=0; i<len(input); i++{
		switch input[i]{
		case '+':
			res = append(res, Token{Plus, "+"})
		case '-':
			res = append(res, Token{Minus, "-"})
		case '(':
			res = append(res, Token{Lparen, "("})
		case ')':
			res = append(res, Token{Rparen, ")"})
		default:
			sb:= strings.Builder{}
			for j:=i; j<len(input); j++{
				if unicode.IsDigit(rune(input[j])){
					sb.WriteRune(rune(input[j]))
					i++
				} else{
					res = append(res, Token{Int, sb.String()})
					i--
					break
				}
			}
		}

	}
	return res
}

func main(){
	input := "(13+4)-(12+1)"
	tokens := Lex(input)
	fmt.Println(tokens)
}
