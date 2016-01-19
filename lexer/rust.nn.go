package main

import (
	"fmt"
	"log"
	"os"
	"sort"
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
		// abstract|alignof|as|become|box|break|const|continue|crate|do|else|enum|extern|false|final|fn|for|if|impl|in|let|loop|macro|match|mod|move|mut|offsetof|override|priv|proc|pub|pure|ref|return|Self|self|sizeof|static|struct|super|trait|true|type|typeof|unsafe|unsized|use|virtual|where|while|yield|println
		{[]bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, false, false, false, false, true, false, true, false, false, false, false, false, true, false, false, true, false, false, false, false, false, true, false, true, false, false, false, true, false, true, false, false, true, false, true, false, false, false, false, false, false, true, false, false, false, false, true, false, false, true, false, false, false, true, false, true, false, true, false, false, false, true, false, false, true, false, true, false, false, true, false, true, false, false, true, false, false, false, false, false, false, false, true, false, false, false, false, false, true, false, false, false, true, true, false, true, false, false, false, true, false, true, false, false, false, true, true, true, false, true, false, true, false, false, true, false, true, false, false, true, false, false, true, false, false, false, false, false, false, true, false, true, false, true, true, false, false, false, false, true, false, false, false, false, false, false, true, true, false, false, false, false, false, true, true, false, false, false, true, false, false, true, false, false, false, false, true, false, false, false, false, false, true, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 83:
					return 1
				case 97:
					return 2
				case 98:
					return 3
				case 99:
					return 4
				case 100:
					return 5
				case 101:
					return 6
				case 102:
					return 7
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return 8
				case 107:
					return -1
				case 108:
					return 9
				case 109:
					return 10
				case 110:
					return -1
				case 111:
					return 11
				case 112:
					return 12
				case 114:
					return 13
				case 115:
					return 14
				case 116:
					return 15
				case 117:
					return 16
				case 118:
					return 17
				case 119:
					return 18
				case 120:
					return -1
				case 121:
					return 19
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 197
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return 183
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return 184
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return 185
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 172
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return 173
				case 112:
					return -1
				case 114:
					return 174
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return 159
				case 112:
					return -1
				case 114:
					return 160
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return 158
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return 147
				case 109:
					return -1
				case 110:
					return 148
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return 149
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return 136
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return 137
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return 138
				case 111:
					return 139
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return 131
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return 132
				case 110:
					return 133
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 126
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return 127
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return 113
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return 114
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return 115
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return 99
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return 100
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return 86
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return 87
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 80
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 59
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return 60
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 61
				case 117:
					return 62
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return 48
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return 49
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return 37
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return 38
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return 31
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return 24
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return 20
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 21
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return 22
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return 23
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 25
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return 26
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return 29
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return 27
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 28
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 30
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return 32
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 33
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return 34
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return 35
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return 36
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return 40
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 39
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return 41
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return 42
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return 46
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return 43
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 44
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return 45
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 47
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return 54
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return 55
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return 50
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 51
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return 52
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return 53
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return 57
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 56
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 58
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return 78
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return 74
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return 66
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return 67
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return 63
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 64
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return 65
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 71
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return 68
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return 69
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 70
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return 72
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return 73
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 75
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return 76
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return 77
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return 79
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return 81
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 82
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return 83
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return 84
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return 85
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return 91
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return 92
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return 88
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return 89
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 90
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return 94
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return 95
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return 93
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 96
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return 97
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return 98
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return 107
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 101
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return 102
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return 103
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return 104
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return 105
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 106
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return 108
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 109
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 110
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return 111
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return 112
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return 120
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 121
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return 117
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return 118
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 116
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 119
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return 124
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return 122
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return 123
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return 125
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 130
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return 128
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return 129
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return 134
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return 135
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return 144
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return 141
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return 140
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return 142
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return 143
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return 145
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 146
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return 156
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return 154
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 150
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 151
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return 152
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return 153
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return 155
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 157
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return 164
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return 161
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 162
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 163
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return 165
				case 116:
					return 166
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 171
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return 167
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return 168
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return 169
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 170
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return 179
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return 178
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 175
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return 176
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return 177
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return 180
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return 181
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return 182
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return 191
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return 186
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return 187
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return 188
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return 189
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return 190
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 192
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return 193
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return 194
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return 195
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return 196
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return 198
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return 199
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 83:
					return -1
				case 97:
					return -1
				case 98:
					return -1
				case 99:
					return -1
				case 100:
					return -1
				case 101:
					return -1
				case 102:
					return -1
				case 103:
					return -1
				case 104:
					return -1
				case 105:
					return -1
				case 107:
					return -1
				case 108:
					return -1
				case 109:
					return -1
				case 110:
					return -1
				case 111:
					return -1
				case 112:
					return -1
				case 114:
					return -1
				case 115:
					return -1
				case 116:
					return -1
				case 117:
					return -1
				case 118:
					return -1
				case 119:
					return -1
				case 120:
					return -1
				case 121:
					return -1
				case 122:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

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

		// =[-+]?[0-9]+
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
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
				case 43:
					return 2
				case 45:
					return 2
				case 61:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 45:
					return -1
				case 61:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 3
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 45:
					return -1
				case 61:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 3
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

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

		// =[-+]?[0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?
		{[]bool{false, false, false, false, true, false, false, true, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 61:
					return 1
				case 69:
					return -1
				case 101:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return 2
				case 45:
					return 2
				case 46:
					return 3
				case 61:
					return -1
				case 69:
					return -1
				case 101:
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
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return 3
				case 61:
					return -1
				case 69:
					return -1
				case 101:
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
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 61:
					return -1
				case 69:
					return -1
				case 101:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 8
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return 3
				case 61:
					return -1
				case 69:
					return 5
				case 101:
					return 5
				}
				switch {
				case 48 <= r && r <= 57:
					return 4
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return 6
				case 45:
					return 6
				case 46:
					return -1
				case 61:
					return -1
				case 69:
					return -1
				case 101:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 7
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 61:
					return -1
				case 69:
					return -1
				case 101:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 7
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 61:
					return -1
				case 69:
					return -1
				case 101:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 7
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 61:
					return -1
				case 69:
					return 5
				case 101:
					return 5
				}
				switch {
				case 48 <= r && r <= 57:
					return 8
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// [0-9]*\.?[0-9]+([eE][-+]?[0-9]+)?
		{[]bool{false, false, true, false, false, true, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return 1
				case 69:
					return -1
				case 101:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 69:
					return -1
				case 101:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 6
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return 1
				case 69:
					return 3
				case 101:
					return 3
				}
				switch {
				case 48 <= r && r <= 57:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return 4
				case 45:
					return 4
				case 46:
					return -1
				case 69:
					return -1
				case 101:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 5
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 69:
					return -1
				case 101:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 5
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 69:
					return -1
				case 101:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 5
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 69:
					return 3
				case 101:
					return 3
				}
				switch {
				case 48 <= r && r <= 57:
					return 6
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

		// (\"([^\"])*\")|(\'([^\'])*\')
		{[]bool{false, false, false, false, true, true, false}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 34:
					return 1
				case 39:
					return 2
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 34:
					return 5
				case 39:
					return 6
				}
				return 6
			},
			func(r rune) int {
				switch r {
				case 34:
					return 3
				case 39:
					return 4
				}
				return 3
			},
			func(r rune) int {
				switch r {
				case 34:
					return 3
				case 39:
					return 4
				}
				return 3
			},
			func(r rune) int {
				switch r {
				case 34:
					return -1
				case 39:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 34:
					return -1
				case 39:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 34:
					return 5
				case 39:
					return 6
				}
				return 6
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1}, nil},

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

		// [\-\+\&\|\^\/\!\:\*\>\<\%\=\.]
		{[]bool{false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 33:
					return 1
				case 37:
					return 1
				case 38:
					return 1
				case 42:
					return 1
				case 43:
					return 1
				case 45:
					return 1
				case 46:
					return 1
				case 47:
					return 1
				case 58:
					return 1
				case 60:
					return 1
				case 61:
					return 1
				case 62:
					return 1
				case 94:
					return 1
				case 124:
					return 1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 38:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 58:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				case 94:
					return -1
				case 124:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1}, []int{ /* End-of-input transitions */ -1, -1}, nil},

		// (<<)|(>>)|(\+=)|(-=)|(\*=)|(\/=)|(%=)|(<<=)|(>>=)|(->)|(==)|(!=)|(\.\.)
		{[]bool{false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true, true, true, true, true, true, true, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 33:
					return 1
				case 37:
					return 2
				case 42:
					return 3
				case 43:
					return 4
				case 45:
					return 5
				case 46:
					return 6
				case 47:
					return 7
				case 60:
					return 8
				case 61:
					return 9
				case 62:
					return 10
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return 23
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return 22
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return 21
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return 20
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return 18
				case 62:
					return 19
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return 17
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return 16
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return 14
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return 13
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return 11
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return 12
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return 15
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 33:
					return -1
				case 37:
					return -1
				case 42:
					return -1
				case 43:
					return -1
				case 45:
					return -1
				case 46:
					return -1
				case 47:
					return -1
				case 60:
					return -1
				case 61:
					return -1
				case 62:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

		// (::)|\#|\[|\]|\(|\)|\{|\}|\,|\;
		{[]bool{false, true, true, true, true, false, true, true, true, true, true, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 35:
					return 1
				case 40:
					return 2
				case 41:
					return 3
				case 44:
					return 4
				case 58:
					return 5
				case 59:
					return 6
				case 91:
					return 7
				case 93:
					return 8
				case 123:
					return 9
				case 125:
					return 10
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 35:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 44:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 91:
					return -1
				case 93:
					return -1
				case 123:
					return -1
				case 125:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 35:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 44:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 91:
					return -1
				case 93:
					return -1
				case 123:
					return -1
				case 125:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 35:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 44:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 91:
					return -1
				case 93:
					return -1
				case 123:
					return -1
				case 125:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 35:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 44:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 91:
					return -1
				case 93:
					return -1
				case 123:
					return -1
				case 125:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 35:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 44:
					return -1
				case 58:
					return 11
				case 59:
					return -1
				case 91:
					return -1
				case 93:
					return -1
				case 123:
					return -1
				case 125:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 35:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 44:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 91:
					return -1
				case 93:
					return -1
				case 123:
					return -1
				case 125:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 35:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 44:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 91:
					return -1
				case 93:
					return -1
				case 123:
					return -1
				case 125:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 35:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 44:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 91:
					return -1
				case 93:
					return -1
				case 123:
					return -1
				case 125:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 35:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 44:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 91:
					return -1
				case 93:
					return -1
				case 123:
					return -1
				case 125:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 35:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 44:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 91:
					return -1
				case 93:
					return -1
				case 123:
					return -1
				case 125:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 35:
					return -1
				case 40:
					return -1
				case 41:
					return -1
				case 44:
					return -1
				case 58:
					return -1
				case 59:
					return -1
				case 91:
					return -1
				case 93:
					return -1
				case 123:
					return -1
				case 125:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}, nil},

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

		// [\_a-zA-Z]+\.[0-9]+
		{[]bool{false, false, false, true}, []func(rune) int{ // Transitions
			func(r rune) int {
				switch r {
				case 46:
					return -1
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
				case 46:
					return 2
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
				case 46:
					return -1
				case 95:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 3
				case 65 <= r && r <= 90:
					return -1
				case 97 <= r && r <= 122:
					return -1
				}
				return -1
			},
			func(r rune) int {
				switch r {
				case 46:
					return -1
				case 95:
					return -1
				}
				switch {
				case 48 <= r && r <= 57:
					return 3
				case 65 <= r && r <= 90:
					return -1
				case 97 <= r && r <= 122:
					return -1
				}
				return -1
			},
		}, []int{ /* Start-of-input transitions */ -1, -1, -1, -1}, []int{ /* End-of-input transitions */ -1, -1, -1, -1}, nil},

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
func main() {
	line := 0

	in, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	/*var WHITESPACE int
	  var COMMENT int*/
	tokens := make(map[string]string)
	tokens_count := make(map[string]int)
	lex := NewLexer(in)
	txt := func() string { return lex.Text() }
	func(yylex *Lexer) {
	OUTER0:
		for {
			switch yylex.next(0) {
			case 0:
				{
					tokens_count["KEYWORD"]++
					tokens[txt()] = "KEYWORD"
				}
			case 1:
				{
					tokens_count["VAR_TYPE"]++
					tokens[txt()] = "VAR_TYPE"
				}
			case 2:
				{
					tokens_count["INTEGER"]++
					tokens_count["OPERATOR"]++
					tokens[txt()[1:len(txt())]] = "INTEGER"
					tokens["="] = "OPERATOR"
				}
			case 3:
				{
					tokens_count["INTEGER"]++
					tokens[txt()] = "INTEGER"
				}
			case 4:
				{
					tokens_count["FLOAT"]++
					tokens_count["OPERATOR"]++
					tokens[txt()[1:len(txt())]] = "FLOAT"
					tokens["="] = "OPERATOR"
				}
			case 5:
				{
					tokens_count["FLOAT"]++
					tokens[txt()] = "FLOAT"
				}
			case 6:
				{
					line++
				}
			case 7:
				{ /* eat up whitespace */
				}
			case 8:
				{ /* eat up comments */
				}
			case 9:
				{
					tokens_count["LITERAL"]++
					tokens[txt()] = "LITERAL"
				}
			case 10:
				{
					fmt.Println("Syntax Error \n", txt(), "is not a valid syntax\n", "at line number", line)
					return
				}
			case 11:
				{
					tokens_count["OPERATOR"]++
					tokens[txt()] = "OPERATOR"
				}
			case 12:
				{
					tokens_count["OPERATOR"]++
					tokens[txt()] = "OPERATOR"
				}
			case 13:
				{
					tokens_count["SYMBOL"]++
					tokens[txt()] = "SYMBOL"
				}
			case 14:
				{
					tokens_count["IDENTIFIER"]++
					tokens[txt()] = "IDENTIFIER"
				}
			case 15:
				{
					tokens_count["IDENTIFIER"]++
					tokens[txt()] = "IDENTIFIER"
				}
			case 16:
				{
					fmt.Println("Syntax Error \n", txt(), "is not a valid syntax\n", "at line number", line)
					return
				}
			default:
				break OUTER0
			}
			continue
		}
		yylex.pop()

	}(lex)

	n := map[string][]string{}
	var a []string
	for k, v := range tokens {
		n[v] = append(n[v], k)
	}
	for k := range n {
		a = append(a, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(a)))

	first := true
	fmt.Printf("%15s %15s %15s\n", "Tokens", "Occurances", "Lexemes")
	for _, k := range a {
		first = true
		for _, s := range n[k] {
			if first {
				fmt.Printf("%15s %15d %15s\n", k, tokens_count[k], s)
				first = false
			} else {
				fmt.Printf("%15s %15s %15s\n", " ", " ", s)
			}
			//  fmt.Println(k,s)
		}
	}

}
