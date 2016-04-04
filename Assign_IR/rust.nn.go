package main

import (
	"fmt"
	"os"
	"strconv"
)
import (
	"bufio"
	"io"
	"strings"
)

type frame struct {
	i            int
	s            string
	line, column int
}
type Lexer struct {
	// The lexer runs in its own goroutine, and communicates via channel 'ch'.
	ch chan frame
	// We record the level of nesting because the action could return, and a
	// subsequent call expects to pick up where it left off. In other words,
	// we're simulating a coroutine.
	// TODO: Support a channel-based variant that compatible with Go's yacc.
	stack []frame
	stale bool

	// The 'l' and 'c' fields were added for
	// https://github.com/wagerlabs/docker/blob/65694e801a7b80930961d70c69cba9f2465459be/buildfile.nex
	// Since then, I introduced the built-in Line() and Column() functions.
	l, c int

	parseResult interface{}

	// The following line makes it easy for scripts to insert fields in the
	// generated code.
	// [NEX_END_OF_LEXER_STRUCT]
}

// NewLexerWithInit creates a new Lexer object, runs the given callback on it,
// then returns it.
func NewLexerWithInit(in io.Reader, initFun func(*Lexer)) *Lexer {
	type dfa struct {
		acc          []bool           // Accepting states.
		f            []func(rune) int // Transitions.
		startf, endf []int            // Transitions at start and end of input.
		nest         []dfa
	}
	yylex := new(Lexer)
	if initFun != nil {
		initFun(yylex)
	}
	yylex.ch = make(chan frame)
	var scan func(in *bufio.Reader, ch chan frame, family []dfa, line, column int)
	scan = func(in *bufio.Reader, ch chan frame, family []dfa, line, column int) {
		// Index of DFA and length of highest-precedence match so far.
		matchi, matchn := 0, -1
		var buf []rune
		n := 0
		checkAccept := func(i int, st int) bool {
			// Higher precedence match? DFAs are run in parallel, so matchn is at most len(buf), hence we may omit the length equality check.
			if family[i].acc[st] && (matchn < n || matchi > i) {
				matchi, matchn = i, n
				return true
			}
			return false
		}
		var state [][2]int
		for i := 0; i < len(family); i++ {
			mark := make([]bool, len(family[i].startf))
			// Every DFA starts at state 0.
			st := 0
			for {
				state = append(state, [2]int{i, st})
				mark[st] = true
				// As we're at the start of input, follow all ^ transitions and append to our list of start states.
				st = family[i].startf[st]
				if -1 == st || mark[st] {
					break
				}
				// We only check for a match after at least one transition.
				checkAccept(i, st)
			}
		}
		atEOF := false
		for {
			if n == len(buf) && !atEOF {
				r, _, err := in.ReadRune()
				switch err {
				case io.EOF:
					atEOF = true
				case nil:
					buf = append(buf, r)
				default:
					panic(err)
				}
			}
			if !atEOF {
				r := buf[n]
				n++
				var nextState [][2]int
				for _, x := range state {
					x[1] = family[x[0]].f[x[1]](r)
					if -1 == x[1] {
						continue
					}
					nextState = append(nextState, x)
					checkAccept(x[0], x[1])
				}
				state = nextState
			} else {
			dollar: // Handle $.
				for _, x := range state {
					mark := make([]bool, len(family[x[0]].endf))
					for {
						mark[x[1]] = true
						x[1] = family[x[0]].endf[x[1]]
						if -1 == x[1] || mark[x[1]] {
							break
						}
						if checkAccept(x[0], x[1]) {
							// Unlike before, we can break off the search. Now that we're at the end, there's no need to maintain the state of each DFA.
							break dollar
						}
					}
				}
				state = nil
			}

			if state == nil {
				lcUpdate := func(r rune) {
					if r == '\n' {
						line++
						column = 0
					} else {
						column++
					}
				}
				// All DFAs stuck. Return last match if it exists, otherwise advance by one rune and restart all DFAs.
				if matchn == -1 {
					if len(buf) == 0 { // This can only happen at the end of input.
						break
					}
					lcUpdate(buf[0])
					buf = buf[1:]
				} else {
					text := string(buf[:matchn])
					buf = buf[matchn:]
					matchn = -1
					ch <- frame{matchi, text, line, column}
					if len(family[matchi].nest) > 0 {
						scan(bufio.NewReader(strings.NewReader(text)), ch, family[matchi].nest, line, column)
					}
					if atEOF {
						break
					}
					for _, r := range text {
						lcUpdate(r)
					}
				}
				n = 0
				for i := 0; i < len(family); i++ {
					state = append(state, [2]int{i, 0})
				}
			}
		}
		ch <- frame{-1, "", line, column}
	}
	go scan(bufio.NewReader(in), yylex.ch, []dfa{
		// abstract
		{[]bool{false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return 1
				case 98:
					return -1
				case 99:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return 2
				case 99:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 114:
					return -1
				case 115:
					return 3
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 114:
					return 5
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 6
				case 98:
					return -1
				case 99:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return 7
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 8
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// alignof
		{[]bool{false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return 1
				case 102:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 108:
					return 2
				case 110:
					return -1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 105:
					return 3
				case 108:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 102:
					return -1
				case 103:
					return 4
				case 105:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 110:
					return 5
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 111:
					return 6
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 102:
					return 7
				case 103:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// as
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return 1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 115:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 115:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// become
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 98:
					return 1
				case 99:
					return -1
				case 101:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return 2
				case 109:
					return -1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 99:
					return 3
				case 101:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 109:
					return -1
				case 111:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 109:
					return 5
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return 6
				case 109:
					return -1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// box
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 98:
					return 1
				case 111:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 111:
					return 2
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 111:
					return -1
				case 120:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 111:
					return -1
				case 120:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// break
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return 1
				case 101:
					return -1
				case 107:
					return -1
				case 114:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 101:
					return -1
				case 107:
					return -1
				case 114:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 101:
					return 3
				case 107:
					return -1
				case 114:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 4
				case 98:
					return -1
				case 101:
					return -1
				case 107:
					return -1
				case 114:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 101:
					return -1
				case 107:
					return 5
				case 114:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 98:
					return -1
				case 101:
					return -1
				case 107:
					return -1
				case 114:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// const
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 99:
					return 1
				case 110:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 110:
					return -1
				case 111:
					return 2
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 110:
					return 3
				case 111:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 115:
					return 4
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				case 116:
					return 5
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// continue
		{[]bool{false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 99:
					return 1
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return 2
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return 3
				case 111:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 116:
					return 4
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 101:
					return -1
				case 105:
					return 5
				case 110:
					return -1
				case 111:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return 6
				case 111:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 116:
					return -1
				case 117:
					return 7
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 101:
					return 8
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// crate
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return 1
				case 101:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 114:
					return 2
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 3
				case 99:
					return -1
				case 101:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 114:
					return -1
				case 116:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return 5
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// do
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 100:
					return 1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 111:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 111:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// else
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return 1
				case 108:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 108:
					return 2
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 108:
					return -1
				case 115:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 4
				case 108:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 108:
					return -1
				case 115:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// enum
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return 1
				case 109:
					return -1
				case 110:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 109:
					return -1
				case 110:
					return 2
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 117:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 109:
					return 4
				case 110:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// extern
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return 1
				case 110:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 120:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 116:
					return 3
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 4
				case 110:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 114:
					return 5
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return 6
				case 114:
					return -1
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 120:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// false
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 102:
					return 1
				case 108:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 2
				case 101:
					return -1
				case 102:
					return -1
				case 108:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 108:
					return 3
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 108:
					return -1
				case 115:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return 5
				case 102:
					return -1
				case 108:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 108:
					return -1
				case 115:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// final
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 102:
					return 1
				case 105:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 102:
					return -1
				case 105:
					return 2
				case 108:
					return -1
				case 110:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 102:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 110:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 4
				case 102:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 102:
					return -1
				case 105:
					return -1
				case 108:
					return 5
				case 110:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 102:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 110:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// fn
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 102:
					return 1
				case 110:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 102:
					return -1
				case 110:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 102:
					return -1
				case 110:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// for
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 102:
					return 1
				case 111:
					return -1
				case 114:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 102:
					return -1
				case 111:
					return 2
				case 114:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 102:
					return -1
				case 111:
					return -1
				case 114:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 102:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// if
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 102:
					return -1
				case 105:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 102:
					return 2
				case 105:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 102:
					return -1
				case 105:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// impl
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 105:
					return 1
				case 108:
					return -1
				case 109:
					return -1
				case 112:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return 2
				case 112:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 112:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 108:
					return 4
				case 109:
					return -1
				case 112:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 112:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// in
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 105:
					return 1
				case 110:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 110:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 110:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// let
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 108:
					return 1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 2
				case 108:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 108:
					return -1
				case 116:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 108:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// loop
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 108:
					return 1
				case 111:
					return -1
				case 112:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 108:
					return -1
				case 111:
					return 2
				case 112:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 108:
					return -1
				case 111:
					return 3
				case 112:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 108:
					return -1
				case 111:
					return -1
				case 112:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 108:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// macro
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 109:
					return 1
				case 111:
					return -1
				case 114:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 2
				case 99:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return 3
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 109:
					return -1
				case 111:
					return 5
				case 114:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// match
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 104:
					return -1
				case 109:
					return 1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 2
				case 99:
					return -1
				case 104:
					return -1
				case 109:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 104:
					return -1
				case 109:
					return -1
				case 116:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return 4
				case 104:
					return -1
				case 109:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 104:
					return 5
				case 109:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 104:
					return -1
				case 109:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// mod
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 109:
					return 1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 109:
					return -1
				case 111:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return 3
				case 109:
					return -1
				case 111:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// move
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 109:
					return 1
				case 111:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 109:
					return -1
				case 111:
					return 2
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 118:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 4
				case 109:
					return -1
				case 111:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 118:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// mut
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 109:
					return 1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 109:
					return -1
				case 116:
					return -1
				case 117:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 109:
					return -1
				case 116:
					return 3
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 109:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// offsetof
		{[]bool{false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 111:
					return 1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return 2
				case 111:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return 3
				case 111:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 111:
					return -1
				case 115:
					return 4
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 5
				case 102:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				case 116:
					return 6
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 111:
					return 7
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return 8
				case 111:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// override
		{[]bool{false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 111:
					return 1
				case 114:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 118:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return 3
				case 105:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 111:
					return -1
				case 114:
					return 4
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 111:
					return -1
				case 114:
					return 5
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 105:
					return 6
				case 111:
					return -1
				case 114:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return 7
				case 101:
					return -1
				case 105:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return 8
				case 105:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 118:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// priv
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 112:
					return 1
				case 114:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 112:
					return -1
				case 114:
					return 2
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return 3
				case 112:
					return -1
				case 114:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 118:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 105:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 118:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// proc
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 111:
					return -1
				case 112:
					return 1
				case 114:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 111:
					return 3
				case 112:
					return -1
				case 114:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return 4
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// pub
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 112:
					return 1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 112:
					return -1
				case 117:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return 3
				case 112:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 98:
					return -1
				case 112:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// pure
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 112:
					return 1
				case 114:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 117:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 112:
					return -1
				case 114:
					return 3
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 4
				case 112:
					return -1
				case 114:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// ref
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 114:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 2
				case 102:
					return -1
				case 114:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return 3
				case 114:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 114:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// return
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 114:
					return 1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 2
				case 110:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 116:
					return 3
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 114:
					return 5
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return 6
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// Self
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 83:
					return 1
				case 101:
					return -1
				case 102:
					return -1
				case 108:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 101:
					return 2
				case 102:
					return -1
				case 108:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 108:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 101:
					return -1
				case 102:
					return 4
				case 108:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 108:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// self
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 108:
					return -1
				case 115:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 2
				case 102:
					return -1
				case 108:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 108:
					return 3
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return 4
				case 108:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 108:
					return -1
				case 115:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// sizeof
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 105:
					return -1
				case 111:
					return -1
				case 115:
					return 1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 105:
					return 2
				case 111:
					return -1
				case 115:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 105:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				case 122:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 4
				case 102:
					return -1
				case 105:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 105:
					return -1
				case 111:
					return 5
				case 115:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return 6
				case 105:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 105:
					return -1
				case 111:
					return -1
				case 115:
					return -1
				case 122:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// static
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 105:
					return -1
				case 115:
					return 1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 105:
					return -1
				case 115:
					return -1
				case 116:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 3
				case 99:
					return -1
				case 105:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 105:
					return -1
				case 115:
					return -1
				case 116:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 105:
					return 5
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return 6
				case 105:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 99:
					return -1
				case 105:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// struct
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 114:
					return -1
				case 115:
					return 1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 2
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 114:
					return 3
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return 5
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 6
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 99:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// super
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return 1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 112:
					return 3
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 4
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 112:
					return -1
				case 114:
					return 5
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// trait
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 114:
					return -1
				case 116:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 114:
					return 2
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 3
				case 105:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return 4
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 114:
					return -1
				case 116:
					return 5
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// true
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 114:
					return -1
				case 116:
					return 1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 114:
					return 2
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 4
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// type
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 112:
					return -1
				case 116:
					return 1
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 112:
					return -1
				case 116:
					return -1
				case 121:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 112:
					return 3
				case 116:
					return -1
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 4
				case 112:
					return -1
				case 116:
					return -1
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 112:
					return -1
				case 116:
					return -1
				case 121:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// typeof
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 116:
					return 1
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 116:
					return -1
				case 121:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 111:
					return -1
				case 112:
					return 3
				case 116:
					return -1
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 4
				case 102:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 116:
					return -1
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 111:
					return 5
				case 112:
					return -1
				case 116:
					return -1
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return 6
				case 111:
					return -1
				case 112:
					return -1
				case 116:
					return -1
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 102:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 116:
					return -1
				case 121:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// unsafe
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				case 117:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 110:
					return 2
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 110:
					return -1
				case 115:
					return 3
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 4
				case 101:
					return -1
				case 102:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 102:
					return 5
				case 110:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return 6
				case 102:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// unsized
		{[]bool{false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				case 117:
					return 1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return 2
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 115:
					return 3
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 105:
					return 4
				case 110:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return 5
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return 6
				case 105:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return 7
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// use
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 115:
					return -1
				case 117:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 115:
					return 2
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 3
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// virtual
		{[]bool{false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return 2
				case 108:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 114:
					return 3
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 114:
					return -1
				case 116:
					return 4
				case 117:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return 5
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return 6
				case 105:
					return -1
				case 108:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 108:
					return 7
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 97:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 114:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// where
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 104:
					return -1
				case 114:
					return -1
				case 119:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 104:
					return 2
				case 114:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 3
				case 104:
					return -1
				case 114:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 104:
					return -1
				case 114:
					return 4
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 5
				case 104:
					return -1
				case 114:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 104:
					return -1
				case 114:
					return -1
				case 119:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// while
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 119:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 104:
					return 2
				case 105:
					return -1
				case 108:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 104:
					return -1
				case 105:
					return 3
				case 108:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return 4
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 5
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 119:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 119:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// yield
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 121:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 105:
					return 2
				case 108:
					return -1
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return 3
				case 105:
					return -1
				case 108:
					return -1
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 108:
					return 4
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return 5
				case 101:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 121:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 100:
					return -1
				case 101:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 121:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// macro_rules!
		{[]bool{false, false, false, false, false, false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 95:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 109:
					return 1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 95:
					return -1
				case 97:
					return 2
				case 99:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 95:
					return -1
				case 97:
					return -1
				case 99:
					return 3
				case 101:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 95:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return 4
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 95:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 111:
					return 5
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 95:
					return 6
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 95:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return 7
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 95:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return 8
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 95:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 108:
					return 9
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 95:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return 10
				case 108:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 95:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return 11
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return 12
				case 95:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 95:
					return -1
				case 97:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// finish!
		{[]bool{false, false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 102:
					return 1
				case 104:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return 2
				case 110:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 110:
					return 3
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return 4
				case 110:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 115:
					return 5
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 102:
					return -1
				case 104:
					return 6
				case 105:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return 7
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// i8|i16|i32|i64|isize|u8|u16|u32|u64|usize|f32|f64|char|bool
		{[]bool{false, false, false, false, false, false, false, false, false, true, false, false, false, true, true, true, true, false, false, false, true, false, false, false, true, true, true, true, false, false, true, true, false, false, true, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return 1
				case 99:
					return 2
				case 101:
					return -1
				case 102:
					return 3
				case 104:
					return -1
				case 105:
					return 4
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return 5
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return 35
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return 32
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return 28
				case 52:
					return -1
				case 54:
					return 29
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return 17
				case 50:
					return -1
				case 51:
					return 18
				case 52:
					return -1
				case 54:
					return 19
				case 56:
					return 20
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return 21
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return 6
				case 50:
					return -1
				case 51:
					return 7
				case 52:
					return -1
				case 54:
					return 8
				case 56:
					return 9
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return 10
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return 16
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return 15
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return 14
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return 11
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return 12
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return 13
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return 27
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return 26
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return 25
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return 22
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return 23
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return 24
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return 31
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return 30
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return 33
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return 34
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return 36
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return 37
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 49:
					return -1
				case 50:
					return -1
				case 51:
					return -1
				case 52:
					return -1
				case 54:
					return -1
				case 56:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 108:
					return -1
				case 111:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 117:
					return -1
				case 122:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// =[ ]*[\-\+]?[0-9]+
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 32:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 61:
					return 1
				}
				switch {
				case 48 <= r && r <= 57:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 32:
					return 2
				case 43:
					return 3
				case 45:
					return 3
				case 61:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 32:
					return 2
				case 43:
					return 3
				case 45:
					return 3
				case 61:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 32:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 61:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 32:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 61:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 4
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// [0-9]+
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch {
				case 48 <= r && r <= 57:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch {
				case 48 <= r && r <= 57:
					return 1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// finish
		{[]bool{false, false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 102:
					return 1
				case 104:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return 2
				case 110:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 110:
					return 3
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return 4
				case 110:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 115:
					return 5
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 102:
					return -1
				case 104:
					return 6
				case 105:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 102:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 110:
					return -1
				case 115:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

		// [\n]
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 10:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 10:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// [ \t]+
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 9:
					return 1
				case 32:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 9:
					return 1
				case 32:
					return 1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// (\/\*([^\*]|[\r\n]|(\*+([^\*\/]|[\r\n])))*\*+\/)|(\/\/[^\n]*)
		{[]bool{false, false, false, true, true, false, false, false, false, true, false}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 10:
					return -1
				case 13:
					return -1
				case 42:
					return -1
				case 47:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 10:
					return -1
				case 13:
					return -1
				case 42:
					return 2
				case 47:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 10:
					return 5
				case 13:
					return 5
				case 42:
					return 6
				case 47:
					return 7
				}
				return 7
			},
			func(r rune) int {
				switch r {
				case 10:
					return -1
				case 13:
					return 4
				case 42:
					return 4
				case 47:
					return 4
				}
				return 4
			},
			func(r rune) int {
				switch r {
				case 10:
					return -1
				case 13:
					return 4
				case 42:
					return 4
				case 47:
					return 4
				}
				return 4
			},
			func(r rune) int {
				switch r {
				case 10:
					return 5
				case 13:
					return 5
				case 42:
					return 6
				case 47:
					return 7
				}
				return 7
			},
			func(r rune) int {
				switch r {
				case 10:
					return 8
				case 13:
					return 8
				case 42:
					return 6
				case 47:
					return 9
				}
				return 10
			},
			func(r rune) int {
				switch r {
				case 10:
					return 5
				case 13:
					return 5
				case 42:
					return 6
				case 47:
					return 7
				}
				return 7
			},
			func(r rune) int {
				switch r {
				case 10:
					return 5
				case 13:
					return 5
				case 42:
					return 6
				case 47:
					return 7
				}
				return 7
			},
			func(r rune) int {
				switch r {
				case 10:
					return -1
				case 13:
					return -1
				case 42:
					return -1
				case 47:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 10:
					return 5
				case 13:
					return 5
				case 42:
					return 6
				case 47:
					return 7
				}
				return 7
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// (\"([^\"])*\")
		{[]bool{false, false, true, false}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 34:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 34:
					return 2
				}
				return 3
			},
			func(r rune) int {
				switch r {
				case 34:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 34:
					return 2
				}
				return 3
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// (\'([^\'])\')
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 39:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 39:
					return -1
				}
				return 2
			},
			func(r rune) int {
				switch r {
				case 39:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 39:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// \>\>
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 62:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 62:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 62:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \<\<
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 60:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \+\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 43:
					return 1
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 61:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \-\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 45:
					return 1
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 61:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \<\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 60:
					return 1
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return -1
				case 61:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return -1
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \>\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 61:
					return -1
				case 62:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return 2
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \*\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 42:
					return 1
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 42:
					return -1
				case 61:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 42:
					return -1
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \/\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 47:
					return 1
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 47:
					return -1
				case 61:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 47:
					return -1
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \%\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 37:
					return 1
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 37:
					return -1
				case 61:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 37:
					return -1
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \&\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 38:
					return 1
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 38:
					return -1
				case 61:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 38:
					return -1
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \<\<\=
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 60:
					return 1
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return 2
				case 61:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return -1
				case 61:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return -1
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// \>\>\=
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 61:
					return -1
				case 62:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return -1
				case 62:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return 3
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// \|\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 61:
					return -1
				case 124:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return 2
				case 124:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return -1
				case 124:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \^\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 61:
					return -1
				case 94:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return 2
				case 94:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return -1
				case 94:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \-\>
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 45:
					return 1
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 62:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				case 62:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \=\>
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 61:
					return 1
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return -1
				case 62:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \=\=
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 61:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// noteq
		{[]bool{false, false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return 1
				case 111:
					return -1
				case 113:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 111:
					return 2
				case 113:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 113:
					return -1
				case 116:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return 4
				case 110:
					return -1
				case 111:
					return -1
				case 113:
					return -1
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 113:
					return 5
				case 116:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 101:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 113:
					return -1
				case 116:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1}, nil},

		// \&\&
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 38:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 38:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 38:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \&mut
		{[]bool{false, false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 38:
					return 1
				case 109:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 38:
					return -1
				case 109:
					return 2
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 38:
					return -1
				case 109:
					return -1
				case 116:
					return -1
				case 117:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 38:
					return -1
				case 109:
					return -1
				case 116:
					return 4
				case 117:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 38:
					return -1
				case 109:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// \|\|
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 124:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 124:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 124:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \*\*
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 42:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 42:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 42:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \.\.
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 46:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 46:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 46:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \.\.\.
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 46:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 46:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 46:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 46:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

		// \-
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 45:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 45:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \+
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 43:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \&
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 38:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 38:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \|
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 124:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 124:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \^
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 94:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 94:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \/
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 47:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 47:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \!
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 33:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \:
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 58:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 58:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \*
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 42:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 42:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \>
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 62:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 62:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \<
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 60:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 60:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \%
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 37:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 37:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \=
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 61:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 61:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \.
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 46:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 46:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \'
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 39:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 39:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// ::
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 58:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 58:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 58:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \#
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 35:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 35:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \[
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 91:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 91:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \]
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 93:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 93:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \(
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 40:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 40:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \)
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 41:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 41:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \{
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 123:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 123:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \}
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 125:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 125:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \,
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 44:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 44:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// \;
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 59:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 59:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// [\_a-zA-Z][\_a-zA-Z0-9]*
		{[]bool{false, true, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 95:
					return 1
				}
				switch {
				case 48 <= r && r <= 57:
					return -1
				case 65 <= r && r <= 90:
					return 1
				case 97 <= r && r <= 122:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 95:
					return 2
				}
				switch {
				case 48 <= r && r <= 57:
					return 2
				case 65 <= r && r <= 90:
					return 2
				case 97 <= r && r <= 122:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 95:
					return 2
				}
				switch {
				case 48 <= r && r <= 57:
					return 2
				case 65 <= r && r <= 90:
					return 2
				case 97 <= r && r <= 122:
					return 2
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// \*\/
		{[]bool{false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 42:
					return 1
				case 47:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 42:
					return -1
				case 47:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 42:
					return -1
				case 47:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1}, nil},

		// [0-9]+[\_a-zA-Z]+[\_a-zA-Z0-9]*
		{[]bool{false, false, true, true, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 95:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 1
				case 65 <= r && r <= 90:
					return -1
				case 97 <= r && r <= 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 95:
					return 2
				}
				switch {
				case 48 <= r && r <= 57:
					return 1
				case 65 <= r && r <= 90:
					return 2
				case 97 <= r && r <= 122:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 95:
					return 3
				}
				switch {
				case 48 <= r && r <= 57:
					return 4
				case 65 <= r && r <= 90:
					return 3
				case 97 <= r && r <= 122:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 95:
					return 3
				}
				switch {
				case 48 <= r && r <= 57:
					return 4
				case 65 <= r && r <= 90:
					return 3
				case 97 <= r && r <= 122:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 95:
					return 4
				}
				switch {
				case 48 <= r && r <= 57:
					return 4
				case 65 <= r && r <= 90:
					return 4
				case 97 <= r && r <= 122:
					return 4
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1}, nil},

		// .
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				return 1
			},
			func(r rune) int {
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},
	}, 0, 0)
	return yylex
}

func NewLexer(in io.Reader) *Lexer {
	return NewLexerWithInit(in, nil)
}

// Text returns the matched text.
func (yylex *Lexer) Text() string {
	return yylex.stack[len(yylex.stack)-1].s
}

// Line returns the current line number.
// The first line is 0.
func (yylex *Lexer) Line() int {
	if len(yylex.stack) == 0 {
		return 0
	}
	return yylex.stack[len(yylex.stack)-1].line
}

// Column returns the current column number.
// The first column is 0.
func (yylex *Lexer) Column() int {
	if len(yylex.stack) == 0 {
		return 0
	}
	return yylex.stack[len(yylex.stack)-1].column
}

func (yylex *Lexer) next(lvl int) int {
	if lvl == len(yylex.stack) {
		l, c := 0, 0
		if lvl > 0 {
			l, c = yylex.stack[lvl-1].line, yylex.stack[lvl-1].column
		}
		yylex.stack = append(yylex.stack, frame{0, "", l, c})
	}
	if lvl == len(yylex.stack)-1 {
		p := &yylex.stack[lvl]
		*p = <-yylex.ch
		yylex.stale = false
	} else {
		yylex.stale = true
	}
	return yylex.stack[lvl].i
}
func (yylex *Lexer) pop() {
	yylex.stack = yylex.stack[:len(yylex.stack)-1]
}
func (yylex Lexer) Error(e string) {
	panic(e)
}

// Lex runs the lexer. Always returns 0.
// When the -s option is given, this function is not generated;
// instead, the NN_FUN macro runs the lexer.
func (yylex *Lexer) Lex(lval *yySymType) int {
OUTER0:
	for {
		switch yylex.next(0) {
		case 0:
			{
				return ABSTRACT
			}
		case 1:
			{
				return ALIGNOF
			}
		case 2:
			{
				return AS
			}
		case 3:
			{
				return BECOME
			}
		case 4:
			{
				return BOX
			}
		case 5:
			{
				return BREAK
			}
		case 6:
			{
				return CONST
			}
		case 7:
			{
				return CONTINUE
			}
		case 8:
			{
				return CRATE
			}
		case 9:
			{
				return DO
			}
		case 10:
			{
				return ELSE
			}
		case 11:
			{
				return ENUM
			}
		case 12:
			{
				return EXTERN
			}
		case 13:
			{
				return FALSE
			}
		case 14:
			{
				return FINAL
			}
		case 15:
			{
				return FN
			}
		case 16:
			{
				return FOR
			}
		case 17:
			{
				return IF
			}
		case 18:
			{
				return IMPL
			}
		case 19:
			{
				return IN
			}
		case 20:
			{
				return LET
			}
		case 21:
			{
				return LOOP
			}
		case 22:
			{
				return MACRO
			}
		case 23:
			{
				return MATCH
			}
		case 24:
			{
				return MOD
			}
		case 25:
			{
				return MOVE
			}
		case 26:
			{
				return MUT
			}
		case 27:
			{
				return OFFSETOF
			}
		case 28:
			{
				return OVERRIDE
			}
		case 29:
			{
				return PRIV
			}
		case 30:
			{
				return PROC
			}
		case 31:
			{
				return PUB
			}
		case 32:
			{
				return PURE
			}
		case 33:
			{
				return REF
			}
		case 34:
			{
				return RETURN
			}
		case 35:
			{
				return SELF
			}
		case 36:
			{
				return SELF
			}
		case 37:
			{
				return SIZEOF
			}
		case 38:
			{
				return STATIC
			}
		case 39:
			{
				return STRUCT
			}
		case 40:
			{
				return SUPER
			}
		case 41:
			{
				return TRAIT
			}
		case 42:
			{
				return TRUE
			}
		case 43:
			{
				return TYPE
			}
		case 44:
			{
				return TYPEOF
			}
		case 45:
			{
				return UNSAFE
			}
		case 46:
			{
				return UNSIZED
			}
		case 47:
			{
				return USE
			}
		case 48:
			{
				return VIRTUAL
			}
		case 49:
			{
				return WHERE
			}
		case 50:
			{
				return WHILE
			}
		case 51:
			{
				return YIELD
			}
		case 52:
			{
				return MACRO_RULES
			}
		case 53:
			{
				return FINISH
			}
		case 54:
			{
				lval.s = yylex.Text()
				return VAR_TYPE
			}
		case 55:
			{
				lval.n, _ = strconv.Atoi(yylex.Text()[space(yylex.Text(), 1):])
				return OPEQ_INT
			}
		case 56:
			{
				lval.n, _ = strconv.Atoi(yylex.Text())
				return LIT_INT
			}
		case 57:
			{
				return FINISH
			}
		case 58:
			{
				line++
			}
		case 59:
			{ /* eat up whitespace */
			}
		case 60:
			{ /* eat up comments */
			}
		case 61:
			{
				lval.s = yylex.Text()
				return LITERAL_STR
			}
		case 62:
			{
				lval.s = yylex.Text()
				return LITERAL_CHAR
			}
		case 63:
			{
				return OP_RSHIFT
			}
		case 64:
			{
				return OP_LSHIFT
			}
		case 65:
			{
				return OP_ADDEQ
			}
		case 66:
			{
				return OP_SUBEQ
			}
		case 67:
			{
				return OP_LEQ
			}
		case 68:
			{
				return OP_GEQ
			}
		case 69:
			{
				return OP_MULEQ
			}
		case 70:
			{
				return OP_DIVEQ
			}
		case 71:
			{
				return OP_MODEQ
			}
		case 72:
			{
				return OP_ANDEQ
			}
		case 73:
			{
				return OP_SHLEQ
			}
		case 74:
			{
				return OP_SHREQ
			}
		case 75:
			{
				return OP_OREQ
			}
		case 76:
			{
				return OP_XOREQ
			}
		case 77:
			{
				return OP_INSIDE
			}
		case 78:
			{
				return OP_FAT_ARROW
			}
		case 79:
			{
				return OP_EQEQ
			}
		case 80:
			{
				return OP_NOTEQ
			}
		case 81:
			{
				return OP_ANDAND
			}
		case 82:
			{
				return OP_ANDMUT
			}
		case 83:
			{
				return OP_OROR
			}
		case 84:
			{
				return OP_POWER
			}
		case 85:
			{
				return OP_DOTDOT
			}
		case 86:
			{
				return OP_DOTDOTDOT
			}
		case 87:
			{
				return int(yylex.Text()[0])
			}
		case 88:
			{
				return int(yylex.Text()[0])
			}
		case 89:
			{
				return int(yylex.Text()[0])
			}
		case 90:
			{
				return int(yylex.Text()[0])
			}
		case 91:
			{
				return int(yylex.Text()[0])
			}
		case 92:
			{
				return int(yylex.Text()[0])
			}
		case 93:
			{
				return int(yylex.Text()[0])
			}
		case 94:
			{
				return int(yylex.Text()[0])
			}
		case 95:
			{
				return int(yylex.Text()[0])
			}
		case 96:
			{
				return int(yylex.Text()[0])
			}
		case 97:
			{
				return int(yylex.Text()[0])
			}
		case 98:
			{
				return int(yylex.Text()[0])
			}
		case 99:
			{
				return int(yylex.Text()[0])
			}
		case 100:
			{
				return int(yylex.Text()[0])
			}
		case 101:
			{
				return int(yylex.Text()[0])
			}
		case 102:
			{
				return SYM_COLCOL
			}
		case 103:
			{
				return int(yylex.Text()[0])
			}
		case 104:
			{
				return SYM_OPEN_SQ
			}
		case 105:
			{
				return SYM_CLOSE_SQ
			}
		case 106:
			{
				return SYM_OPEN_ROUND
			}
		case 107:
			{
				return SYM_CLOSE_ROUND
			}
		case 108:
			{
				return SYM_OPEN_CURLY
			}
		case 109:
			{
				return SYM_CLOSE_CURLY
			}
		case 110:
			{
				return int(yylex.Text()[0])
			}
		case 111:
			{
				return int(yylex.Text()[0])
			}
		case 112:
			{
				lval.s = yylex.Text()
				return IDENTIFIER
			}
		case 113:
			{
				fmt.Println("Syntax Error \n", "at line number\n", yylex.Text(), "is not a valid syntax\n", line+1)
			}
		case 114:
			{
				fmt.Println("Syntax Error \n", "at line number\n", yylex.Text(), "is not a valid syntax\n", line+1)
			}
		case 115:
			{
				fmt.Println("Syntax Error \n", "at line number\n", yylex.Text(), "is not a valid syntax\n", line+1)
			}
		default:
			break OUTER0
		}
		continue
	}
	yylex.pop()

	return 0
}
func main() {

	/*  in,err := os.Open(os.Args[1])
	    if err != nil {
	            log.Fatal(err)
	    }
	*/
	yyParse(NewLexer(os.Stdin))

}
