// generated by Textmapper; DO NOT EDIT

package test

import (
	"context"
	"fmt"

	"github.com/inspirer/textmapper/parsers/test/token"
)

// Parser is a table-driven LALR parser for test.
type Parser struct {
	listener Listener

	next symbol

	// Tokens to be reported with the next shift. Only non-empty when next.symbol != noToken.
	pending []symbol
}

type SyntaxError struct {
	Offset    int
	Endoffset int
}

func (e SyntaxError) Error() string {
	return "syntax error"
}

type symbol struct {
	symbol    int32
	offset    int
	endoffset int
}

type stackEntry struct {
	sym   symbol
	state int16
	value interface{}
}

func (p *Parser) Init(l Listener) {
	p.listener = l
	if cap(p.pending) < startTokenBufferSize {
		p.pending = make([]symbol, 0, startTokenBufferSize)
	}
}

const (
	startStackSize       = 256
	startTokenBufferSize = 16
	noToken              = int32(token.UNAVAILABLE)
	eoiToken             = int32(token.EOI)
	debugSyntax          = false
)

func (p *Parser) ParseTest(ctx context.Context, lexer *Lexer) error {
	_, err := p.parse(ctx, 1, 139, lexer)
	return err
}

func (p *Parser) ParseDecl1(ctx context.Context, lexer *Lexer) (int, error) {
	v, err := p.parse(ctx, 2, 140, lexer)
	val, _ := v.(int)
	return val, err
}

type session struct {
	shiftCounter int32
	cache        map[uint64]bool
}

func (p *Parser) parse(ctx context.Context, start, end int16, lexer *Lexer) (interface{}, error) {
	p.pending = p.pending[:0]
	var s session
	s.cache = make(map[uint64]bool)
	state := start

	var alloc [startStackSize]stackEntry
	stack := append(alloc[:0], stackEntry{state: state})
	p.fetchNext(lexer, stack)

	for state != end {
		action := tmAction[state]
		if action < -2 {
			// Lookahead is needed.
			if p.next.symbol == noToken {
				p.fetchNext(lexer, stack)
			}
			action = lalr(action, p.next.symbol)
		}

		if action >= 0 {
			// Reduce.
			rule := action
			ln := int(tmRuleLen[rule])

			var entry stackEntry
			entry.sym.symbol = tmRuleSymbol[rule]
			rhs := stack[len(stack)-ln:]
			stack = stack[:len(stack)-ln]
			if ln == 0 {
				if p.next.symbol == noToken {
					p.fetchNext(lexer, stack)
				}
				entry.sym.offset, entry.sym.endoffset = p.next.offset, p.next.offset
			} else {
				entry.sym.offset = rhs[0].sym.offset
				entry.sym.endoffset = rhs[ln-1].sym.endoffset
			}
			if err := p.applyRule(ctx, rule, &entry, rhs, lexer, &s); err != nil {
				return nil, err
			}
			if debugSyntax {
				fmt.Printf("reduced to: %v\n", symbolName(entry.sym.symbol))
			}
			state = gotoState(stack[len(stack)-1].state, entry.sym.symbol)
			entry.state = state
			stack = append(stack, entry)

		} else if action == -1 {
			if s.shiftCounter++; s.shiftCounter&0x1ff == 0 {
				// Note: checking for context cancellation is expensive so we do it from time to time.
				select {
				case <-ctx.Done():
					return nil, ctx.Err()
				default:
				}
			}

			// Shift.
			if p.next.symbol == noToken {
				p.fetchNext(lexer, stack)
			}
			state = gotoState(state, p.next.symbol)
			if state >= 0 {
				stack = append(stack, stackEntry{
					sym:   p.next,
					state: state,
					value: lexer.Value(),
				})
				if debugSyntax {
					fmt.Printf("shift: %v (%s)\n", symbolName(p.next.symbol), lexer.Text())
				}
				if len(p.pending) > 0 {
					for _, tok := range p.pending {
						p.reportIgnoredToken(ctx, tok)
					}
					p.pending = p.pending[:0]
				}
				if p.next.symbol != eoiToken {
					switch token.Token(p.next.symbol) {
					case token.IDENTIFIER:
						p.listener(Identifier, 0, p.next.offset, p.next.endoffset)
					}
					p.next.symbol = noToken
				}
			}
		}

		if action == -2 || state == -1 {
			break
		}
	}

	if state != end {
		if p.next.symbol == noToken {
			p.fetchNext(lexer, stack)
		}
		err := SyntaxError{
			Offset:    p.next.offset,
			Endoffset: p.next.endoffset,
		}
		return nil, err
	}

	return stack[len(stack)-2].value, nil
}

func lalr(action, next int32) int32 {
	a := -action - 3
	for ; tmLalr[a] >= 0; a += 2 {
		if tmLalr[a] == next {
			break
		}
	}
	return tmLalr[a+1]
}

func gotoState(state int16, symbol int32) int16 {
	min := tmGoto[symbol]
	max := tmGoto[symbol+1]

	if max-min < 32 {
		for i := min; i < max; i += 2 {
			if tmFromTo[i] == state {
				return tmFromTo[i+1]
			}
		}
	} else {
		for min < max {
			e := (min + max) >> 1 &^ int32(1)
			i := tmFromTo[e]
			if i == state {
				return tmFromTo[e+1]
			} else if i < state {
				min = e + 2
			} else {
				max = e
			}
		}
	}
	return -1
}

func (p *Parser) fetchNext(lexer *Lexer, stack []stackEntry) {
restart:
	tok := lexer.Next()
	switch tok {
	case token.MULTILINECOMMENT, token.SINGLELINECOMMENT, token.INVALID_TOKEN:
		s, e := lexer.Pos()
		tok := symbol{int32(tok), s, e}
		p.pending = append(p.pending, tok)
		goto restart
	}
	p.next.symbol = int32(tok)
	p.next.offset, p.next.endoffset = lexer.Pos()
}

func lookaheadNext(lexer *Lexer) int32 {
restart:
	tok := lexer.Next()
	switch tok {
	case token.MULTILINECOMMENT, token.SINGLELINECOMMENT, token.INVALID_TOKEN:
		goto restart
	}
	return int32(tok)
}

func lookaheadRule(ctx context.Context, lexer *Lexer, next, rule int32, s *session) (sym int32, err error) {
	switch rule {
	case 93:
		var ok bool
		if ok, err = lookahead(ctx, lexer, next, 0, 136, s); ok {
			sym = 42 /* lookahead_FooLookahead */
		} else {
			sym = 43 /* lookahead_notFooLookahead */
		}
		return
	}
	return 0, nil
}

func AtFooLookahead(ctx context.Context, lexer *Lexer, next int32, s *session) (bool, error) {
	if debugSyntax {
		fmt.Printf("lookahead FooLookahead, next: %v\n", symbolName(next))
	}
	return lookahead(ctx, lexer, next, 0, 136, s)
}

func lookahead(ctx context.Context, l *Lexer, next int32, start, end int16, s *session) (bool, error) {
	var lexer Lexer = *l

	// Use memoization for recursive lookaheads.
	if next == noToken {
		next = lookaheadNext(&lexer)
	}
	key := uint64(l.tokenOffset) + uint64(end)<<40
	if ret, ok := s.cache[key]; ok {
		return ret, nil
	}

	var allocated [64]stackEntry
	state := start
	stack := append(allocated[:0], stackEntry{state: state})

	for state != end {
		action := tmAction[state]
		if action < -2 {
			// Lookahead is needed.
			if next == noToken {
				next = lookaheadNext(&lexer)
			}
			action = lalr(action, next)
		}

		if action >= 0 {
			// Reduce.
			rule := action
			ln := int(tmRuleLen[rule])

			var entry stackEntry
			entry.sym.symbol = tmRuleSymbol[rule]
			stack = stack[:len(stack)-ln]
			sym, err := lookaheadRule(ctx, &lexer, next, rule, s)
			if err != nil {
				return false, err
			}
			if sym != 0 {
				entry.sym.symbol = sym
			}
			if debugSyntax {
				fmt.Printf("lookahead reduced to: %v\n", symbolName(entry.sym.symbol))
			}
			state = gotoState(stack[len(stack)-1].state, entry.sym.symbol)
			entry.state = state
			stack = append(stack, entry)

		} else if action == -1 {
			if s.shiftCounter++; s.shiftCounter&0x1ff == 0 {
				// Note: checking for context cancellation is expensive so we do it from time to time.
				select {
				case <-ctx.Done():
					return false, ctx.Err()
				default:
				}
			}

			// Shift.
			if next == noToken {
				next = lookaheadNext(&lexer)
			}
			state = gotoState(state, next)
			stack = append(stack, stackEntry{
				sym:   symbol{symbol: next},
				state: state,
			})
			if debugSyntax {
				fmt.Printf("lookahead shift: %v (%s)\n", symbolName(next), lexer.Text())
			}
			if state != -1 && next != eoiToken {
				next = noToken
			}
		}

		if action == -2 || state == -1 {
			break
		}
	}

	s.cache[key] = state == end
	if debugSyntax {
		fmt.Printf("lookahead done: %v\n", state == end)
	}
	return state == end, nil
}

func (p *Parser) applyRule(ctx context.Context, rule int32, lhs *stackEntry, rhs []stackEntry, lexer *Lexer, s *session) (err error) {
	switch rule {
	case 5: // Declaration : '{' '-' '-' Declaration_list '}'
		p.listener(Negation, 0, rhs[1].sym.offset, rhs[2].sym.endoffset)
	case 6: // Declaration : '{' '-' '-' '}'
		p.listener(Negation, 0, rhs[1].sym.offset, rhs[2].sym.endoffset)
	case 7: // Declaration : '{' '-' Declaration_list '}'
		p.listener(Negation, 0, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 8: // Declaration : '{' '-' '}'
		p.listener(Negation, 0, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 11: // Declaration : lastInt
		{
			println("it works")
		}
	case 12: // Declaration : IntegerConstant '[' ']'
		nn0, _ := rhs[0].value.(int)
		{
			switch nn0 {
			case 7:
				p.listener(Int7, 0, rhs[0].sym.offset, rhs[2].sym.endoffset)
			case 9:
				p.listener(Int9, 0, rhs[0].sym.offset, rhs[2].sym.endoffset)
			}
		}
	case 13: // Declaration : IntegerConstant
		nn0, _ := rhs[0].value.(int)
		{
			switch nn0 {
			case 7:
				p.listener(Int7, 0, rhs[0].sym.offset, rhs[0].sym.endoffset)
			case 9:
				p.listener(Int9, 0, rhs[0].sym.offset, rhs[0].sym.endoffset)
			}
		}
	case 15: // Declaration : 'test' '(' empty1 ')'
		p.listener(Empty1, 0, rhs[2].sym.offset, rhs[2].sym.endoffset)
	case 17: // Declaration : 'test' IntegerConstant
		p.listener(Icon, InTest, rhs[1].sym.offset, rhs[1].sym.endoffset)
	case 18: // Declaration : 'eval' lookahead_notFooLookahead '(' expr ')' empty1
		fixTrailingWS(lhs, rhs)
	case 21: // Declaration : 'decl2' ':' QualifiedNameopt
		fixTrailingWS(lhs, rhs)
	case 87: // customPlus : '\\' primaryExpr '+' expr
		{
			p.listener(PlusExpr, 0, rhs[0].sym.offset, rhs[3].sym.endoffset)
		}
	case 93:
		var ok bool
		if ok, err = AtFooLookahead(ctx, lexer, p.next.symbol, s); ok {
			lhs.sym.symbol = 42 /* lookahead_FooLookahead */
		} else {
			lhs.sym.symbol = 43 /* lookahead_notFooLookahead */
		}
		return
	}
	if nt := tmRuleType[rule]; nt != 0 {
		p.listener(NodeType(nt&0xffff), NodeFlags(nt>>16), lhs.sym.offset, lhs.sym.endoffset)
	}
	return
}

func fixTrailingWS(lhs *stackEntry, rhs []stackEntry) {
	last := len(rhs) - 1
	if last < 0 {
		return
	}
	for last >= 0 && rhs[last].sym.offset == rhs[last].sym.endoffset {
		last--
	}
	if last >= 0 {
		lhs.sym.endoffset = rhs[last].sym.endoffset
	} else {
		lhs.sym.endoffset = lhs.sym.offset
	}
}

func (p *Parser) reportIgnoredToken(ctx context.Context, tok symbol) {
	var t NodeType
	switch token.Token(tok.symbol) {
	case token.MULTILINECOMMENT:
		t = MultiLineComment
	case token.SINGLELINECOMMENT:
		t = SingleLineComment
	case token.INVALID_TOKEN:
		t = InvalidToken
	default:
		return
	}
	if debugSyntax {
		fmt.Printf("ignored: %v as %v\n", token.Token(tok.symbol), t)
	}
	p.listener(t, 0, tok.offset, tok.endoffset)
}

func parserEnd() {}
