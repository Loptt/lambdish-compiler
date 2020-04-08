// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 37: // ['%','%']
			return 2
		case r == 40: // ['(','(']
			return 3
		case r == 41: // [')',')']
			return 4
		case r == 42: // ['*','*']
			return 2
		case r == 43: // ['+','+']
			return 5
		case r == 44: // [',',',']
			return 6
		case r == 45: // ['-','-']
			return 5
		case r == 47: // ['/','/']
			return 2
		case 49 <= r && r <= 57: // ['1','9']
			return 7
		case r == 58: // [':',':']
			return 8
		case r == 61: // ['=','=']
			return 9
		case r == 64: // ['@','@']
			return 10
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case r == 91: // ['[','[']
			return 12
		case r == 92: // ['\','\']
			return 13
		case r == 93: // [']',']']
			return 14
		case r == 97: // ['a','a']
			return 11
		case r == 98: // ['b','b']
			return 15
		case r == 99: // ['c','c']
			return 16
		case 100 <= r && r <= 101: // ['d','e']
			return 11
		case r == 102: // ['f','f']
			return 17
		case 103 <= r && r <= 109: // ['g','m']
			return 11
		case r == 110: // ['n','n']
			return 18
		case 111 <= r && r <= 115: // ['o','s']
			return 11
		case r == 116: // ['t','t']
			return 19
		case 117 <= r && r <= 122: // ['u','z']
			return 11
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S4
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 20
		case 49 <= r && r <= 57: // ['1','9']
			return 7
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		case r == 58: // [':',':']
			return 21
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		case r == 62: // ['>','>']
			return 22
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 23
		case 65 <= r && r <= 90: // ['A','Z']
			return 24
		case 97 <= r && r <= 122: // ['a','z']
			return 24
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 122: // ['a','z']
			return 11
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		case r == 92: // ['\','\']
			return 26
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S15
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 110: // ['a','n']
			return 11
		case r == 111: // ['o','o']
			return 27
		case 112 <= r && r <= 122: // ['p','z']
			return 11
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 103: // ['a','g']
			return 11
		case r == 104: // ['h','h']
			return 28
		case 105 <= r && r <= 122: // ['i','z']
			return 11
		}
		return NoState
	},
	// S17
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case r == 97: // ['a','a']
			return 29
		case 98 <= r && r <= 122: // ['b','z']
			return 11
		}
		return NoState
	},
	// S18
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 116: // ['a','t']
			return 11
		case r == 117: // ['u','u']
			return 30
		case 118 <= r && r <= 122: // ['v','z']
			return 11
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 113: // ['a','q']
			return 11
		case r == 114: // ['r','r']
			return 31
		case 115 <= r && r <= 122: // ['s','z']
			return 11
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 32
		}
		return NoState
	},
	// S21
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S22
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		case r == 64: // ['@','@']
			return 33
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		case r == 64: // ['@','@']
			return 33
		}
		return NoState
	},
	// S25
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 122: // ['a','z']
			return 11
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S27
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 110: // ['a','n']
			return 11
		case r == 111: // ['o','o']
			return 34
		case 112 <= r && r <= 122: // ['p','z']
			return 11
		}
		return NoState
	},
	// S28
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case r == 97: // ['a','a']
			return 35
		case 98 <= r && r <= 122: // ['b','z']
			return 11
		}
		return NoState
	},
	// S29
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 107: // ['a','k']
			return 11
		case r == 108: // ['l','l']
			return 36
		case 109 <= r && r <= 122: // ['m','z']
			return 11
		}
		return NoState
	},
	// S30
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 108: // ['a','l']
			return 11
		case r == 109: // ['m','m']
			return 37
		case 110 <= r && r <= 122: // ['n','z']
			return 11
		}
		return NoState
	},
	// S31
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 116: // ['a','t']
			return 11
		case r == 117: // ['u','u']
			return 38
		case 118 <= r && r <= 122: // ['v','z']
			return 11
		}
		return NoState
	},
	// S32
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 32
		}
		return NoState
	},
	// S33
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S34
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 107: // ['a','k']
			return 11
		case r == 108: // ['l','l']
			return 39
		case 109 <= r && r <= 122: // ['m','z']
			return 11
		}
		return NoState
	},
	// S35
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 113: // ['a','q']
			return 11
		case r == 114: // ['r','r']
			return 40
		case 115 <= r && r <= 122: // ['s','z']
			return 11
		}
		return NoState
	},
	// S36
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 114: // ['a','r']
			return 11
		case r == 115: // ['s','s']
			return 41
		case 116 <= r && r <= 122: // ['t','z']
			return 11
		}
		return NoState
	},
	// S37
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 122: // ['a','z']
			return 11
		}
		return NoState
	},
	// S38
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 100: // ['a','d']
			return 11
		case r == 101: // ['e','e']
			return 42
		case 102 <= r && r <= 122: // ['f','z']
			return 11
		}
		return NoState
	},
	// S39
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 122: // ['a','z']
			return 11
		}
		return NoState
	},
	// S40
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 122: // ['a','z']
			return 11
		}
		return NoState
	},
	// S41
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 100: // ['a','d']
			return 11
		case r == 101: // ['e','e']
			return 43
		case 102 <= r && r <= 122: // ['f','z']
			return 11
		}
		return NoState
	},
	// S42
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 122: // ['a','z']
			return 11
		}
		return NoState
	},
	// S43
	func(r rune) int {
		switch {
		case 49 <= r && r <= 57: // ['1','9']
			return 25
		case 65 <= r && r <= 90: // ['A','Z']
			return 11
		case 97 <= r && r <= 122: // ['a','z']
			return 11
		}
		return NoState
	},
}
