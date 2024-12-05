package day03

import (
	"fmt"
	"io"
	"iter"
	"strconv"
)

type Parser struct {
	l *Lexer

	state parseStateFunc
	err   error
}

func Parse(r io.Reader) *Parser {
	l := Lex(r)

	return &Parser{
		l:     l,
		state: parseStart,
	}
}

func (p *Parser) Err() error {
	if p.err != nil {
		return p.err
	}

	return p.l.Err()
}

func (p *Parser) Ops() iter.Seq[Op] {
	return func(yield func(Op) bool) {
		running := true

		nextToken, done := iter.Pull(p.l.Tokens())
		defer done()

		next := func() Token {
			tok, ok := nextToken()
			if !ok {
				return Token{Kind: TokenEOF}
			}
			return tok
		}

		emit := func(op Op) {
			running = yield(op)
		}

		for running && p.state != nil {
			p.state = p.state(p, next, emit)
		}
	}
}

type Op any

type parseStateFunc func(*Parser, func() Token, func(Op)) parseStateFunc

func parseStart(p *Parser, next func() Token, emit func(Op)) parseStateFunc {
	t := next()

	switch t.Kind {
	case TokenEOF:
		return nil

	case TokenInvalid,
		TokenComma,
		TokenNumber,
		TokenOther,
		TokenParenLeft,
		TokenParenRight:
		// Ignore invalid tokens for now
		return parseStart

	case TokenSymbol:
		switch t.Value {
		case "mul":
			return parseMul
		case "do":
			return parseDo
		case "don't":
			return parseDont
		default:
			p.err = fmt.Errorf("unknown symbol: %q", t.Value)
			return nil
		}

	default:
		panic(fmt.Sprintf("unexpected TokenKind: %#v", t.Kind))
	}
}

type (
	OpDo   struct{}
	OpDont struct{}
)

type OpMul struct {
	A, B int
}

func parseMul(p *Parser, next func() Token, emit func(Op)) parseStateFunc {
	if t := next(); t.Kind != TokenParenLeft {
		return parseStart
	}

	t := next()
	if t.Kind != TokenNumber {
		return parseStart
	}

	a, err := strconv.Atoi(t.Value)
	if err != nil {
		p.err = fmt.Errorf("failed to parse number: %w", err)
	}

	if t := next(); t.Kind != TokenComma {
		return parseStart
	}

	t = next()
	if t.Kind != TokenNumber {
		return parseStart
	}

	b, err := strconv.Atoi(t.Value)
	if err != nil {
		p.err = fmt.Errorf("failed to parse number: %w", err)
	}

	if t := next(); t.Kind != TokenParenRight {
		return parseStart
	}

	emit(OpMul{A: a, B: b})

	return parseStart
}

func parseDo(p *Parser, next func() Token, emit func(Op)) parseStateFunc {
	if t := next(); t.Kind != TokenParenLeft {
		return parseStart
	}
	if t := next(); t.Kind != TokenParenRight {
		return parseStart
	}

	emit(OpDo{})

	return parseStart
}

func parseDont(p *Parser, next func() Token, emit func(Op)) parseStateFunc {
	if t := next(); t.Kind != TokenParenLeft {
		return parseStart
	}
	if t := next(); t.Kind != TokenParenRight {
		return parseStart
	}

	emit(OpDont{})

	return parseStart
}
