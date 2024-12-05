package day03

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"iter"
	"strings"
)

type Lexer struct {
	r io.ByteReader

	buf bytes.Buffer
	pos int

	state lexStateFunc
	err   error
}

func Lex(r io.Reader) *Lexer {
	br, ok := r.(io.ByteReader)
	if !ok {
		br = bufio.NewReader(r)
	}
	return &Lexer{
		r:     br,
		state: lexStart,
	}
}

func (l *Lexer) Err() error {
	if errors.Is(l.err, io.EOF) {
		return nil
	}

	return l.err
}

// much of this adapted from: https://go.dev/talks/2011/lex.slide

func (l *Lexer) Tokens() iter.Seq[Token] {
	return func(yield func(Token) bool) {
		running := true

		emit := func(kind TokenKind) {
			pos := l.pos
			l.pos = 0

			t := Token{
				Kind:  kind,
				Value: string(l.buf.Next(pos)),
			}

			running = yield(t)
		}

		for running && l.state != nil {
			l.state = l.state(l, emit)
		}
	}
}

func (l *Lexer) next() byte {
	if l.err != nil {
		return 0
	}

	if l.pos < l.buf.Len() {
		l.pos++
		return l.buf.Bytes()[l.pos-1]
	}

	b, err := l.r.ReadByte()
	if err != nil {
		if errors.Is(err, io.EOF) {
			l.err = err
		} else {
			l.err = fmt.Errorf("failed to read next byte: %w", err)
		}
		return 0
	}

	l.buf.WriteByte(b)
	l.pos++

	return b
}

func (l *Lexer) backup() {
	if l.err != nil {
		return
	}
	l.pos--
}

func (l *Lexer) peek() byte {
	b := l.next()
	l.backup()
	return b
}

// accept consumes the next byte if it's from the valid set.
func (l *Lexer) accept(valid string) bool {
	if strings.IndexByte(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of bytes from the valid set.
func (l *Lexer) acceptRun(valid string) {
	for strings.IndexByte(valid, l.next()) >= 0 {
	}
	l.backup()
}

// acceptLiteral consumes bytes if they sequentially match the literal string.
func (l *Lexer) acceptLiteral(literal string) bool {
	for i, c := range []byte(literal) {
		if l.next() != c {
			for range i {
				l.backup()
			}
			l.backup()
			return false
		}
	}

	return true
}

type Token struct {
	Kind  TokenKind
	Value string
}

func (t Token) String() string {
	switch t.Kind {
	case TokenInvalid:
		return fmt.Sprintf("INVALID(%v)", t.Value)
	case TokenEOF:
		return "EOF"

	case TokenNumber:
		return fmt.Sprintf("number(%v)", t.Value)
	case TokenSymbol:
		return fmt.Sprintf("symbol(%v)", t.Value)

	case TokenOther:
		return fmt.Sprintf("other(%v)", t.Value)

	case TokenComma, TokenParenLeft, TokenParenRight:
		return fmt.Sprintf("%q", t.Value)

	default:
		panic(fmt.Sprintf("unexpected day03.TokenKind: %#v", t.Kind))
	}
}

type TokenKind int

const (
	TokenInvalid TokenKind = iota
	TokenEOF

	TokenParenLeft
	TokenParenRight
	TokenComma

	TokenSymbol
	TokenNumber

	TokenOther
)

type lexStateFunc func(*Lexer, func(TokenKind)) lexStateFunc

func lexStart(l *Lexer, emit func(TokenKind)) lexStateFunc {
	b := l.next()

	switch b {
	case 0:
		emit(TokenEOF)
		return nil

	case '(':
		emit(TokenParenLeft)
		return lexStart
	case ')':
		emit(TokenParenRight)
		return lexStart
	case ',':
		emit(TokenComma)
		return lexStart

	case 'm', 'd':
		return lexSymbol
	}

	switch {
	case '0' <= b && b <= '9':
		return lexNumber

	default:
		return lexOther
	}
}

func lexSymbol(l *Lexer, emit func(TokenKind)) lexStateFunc {
	// back up to start from the beginning of the (potential) symbol
	l.backup()

	switch {
	case
		l.acceptLiteral("mul"),
		l.acceptLiteral("don't"), // consume "don't" first before do
		l.acceptLiteral("do"):
		emit(TokenSymbol)
		return lexStart

	default:
		// it's not a symbol, consume the byte as part of an "other" token
		l.next()
		return lexOther
	}
}

func lexNumber(l *Lexer, emit func(TokenKind)) lexStateFunc {
	l.acceptRun("0123456789")
	emit(TokenNumber)
	return lexStart
}

func lexOther(l *Lexer, emit func(TokenKind)) lexStateFunc {
	for {
		switch b := l.next(); b {
		case 0,
			'(', ')', ',', 'm', 'd',
			'0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			l.backup()
			emit(TokenOther)
			return lexStart
		}
	}
}
