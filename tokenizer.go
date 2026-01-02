package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

const (
	TKN_BRACE_OPEN    = "BRACE_OPEN"
	TKN_BRACE_CLOSE   = "BRACE_CLOSE"
	TKN_STRING        = "STRING"
	TKN_NUMBER        = "NUMBER"
	TKN_COLON         = "COLON"
	TKN_COMMA         = "COMMA"
	TKN_TRUE          = "TRUE"
	TKN_FALSE         = "FALSE"
	TKN_NULL          = "NULL"
	TKN_BRACKET_OPEN  = "BRACKET_OPEN"
	TKN_BRACKET_CLOSE = "BRACKET_CLOSE"
)

type Token struct {
	Type  string
	Value string
}

func Tokenize(jsonString string) ([]Token, error) {
	current := 0
	stringLength := len(jsonString)

	tokens := []Token{}

	for current < stringLength {
		char := jsonString[current]

		if unicode.IsSpace(rune(char)) {
			current++
			continue
		}

		switch char {
		case '{':
			tokens = append(tokens, Token{Type: TKN_BRACE_OPEN, Value: "{"})
			current++
		case '}':
			tokens = append(tokens, Token{Type: TKN_BRACE_CLOSE, Value: "}"})
			current++
		case '[':
			tokens = append(tokens, Token{Type: TKN_BRACKET_OPEN, Value: "["})
			current++
		case ']':
			tokens = append(tokens, Token{Type: TKN_BRACKET_CLOSE, Value: "]"})
			current++
		case ':':
			tokens = append(tokens, Token{Type: TKN_COLON, Value: ":"})
			current++
		case ',':
			tokens = append(tokens, Token{Type: TKN_COMMA, Value: ","})
			current++
		case '"':
			current++
			start := current
			for current < stringLength && jsonString[current] != '"' {
				if jsonString[current] == '\\' {
					current++
				}
				current++
			}

			str := jsonString[start:current]

			if current >= stringLength {
				return nil, fmt.Errorf("undeterminated string: " + str)
			}

			tokens = append(tokens, Token{Type: TKN_STRING, Value: str})
			current++
		default:
			rest := jsonString[current:]
			if strings.HasPrefix(rest, "true") {
				tokens = append(tokens, Token{Type: TKN_TRUE, Value: "true"})
				current += 4
			} else if strings.HasPrefix(rest, "false") {
				tokens = append(tokens, Token{Type: TKN_FALSE, Value: "false"})
				current += 5
			} else if strings.HasPrefix(rest, "null") {
				tokens = append(tokens, Token{Type: TKN_NULL, Value: "null"})
				current += 4
			} else if unicode.IsNumber(rune(char)) || char == '-' {
				start := current
				current++

				hasDot := false
				hasExp := false
				expDigits := 0

				for current < stringLength {
					c := jsonString[current]

					if unicode.IsDigit(rune(c)) {
						current++
						if hasExp {
							expDigits++
						}
					} else if c == '.' {
						if hasDot || hasExp {
							return nil, fmt.Errorf("invalid number: multiple dots or dot after exponent at position %d", current)
						}
						hasDot = true
						current++
					} else if c == 'e' || c == 'E' {
						if hasExp {
							return nil, fmt.Errorf("invalid number: multiple exponents at position %d", current)
						}
						hasExp = true
						current++
						if current < stringLength && jsonString[current] == '+' || jsonString[current] == '-' {
							current++
						}
						expDigits = 0
					} else {
						break
					}
				}

				number := jsonString[start:current]
				isValidJSONNumber := isValidNumber(number)
				if !isValidJSONNumber {
					return nil, fmt.Errorf("invalid JSON number: %s", number)
				}

				if hasExp && expDigits == 0 {
					return nil, fmt.Errorf("invalid number: exponent missing digits in '%s'", number)
				}

				if _, err := strconv.ParseFloat(number, 64); err != nil {
					return nil, fmt.Errorf("invalid number: %s", number)
				}

				tokens = append(tokens, Token{Type: TKN_NUMBER, Value: number})
			} else {
				return nil, fmt.Errorf("unexpected character: %c, position %d", char, current)
			}
		}
	}

	return tokens, nil
}

func printTokens(tokens []Token) {
	fmt.Printf("%-14s | %s\n", "Type", "Value")
	fmt.Println(strings.Repeat("-", 50))

	for _, token := range tokens {
		fmt.Printf("%-14s | %s\n", token.Type, token.Value)
	}
}
