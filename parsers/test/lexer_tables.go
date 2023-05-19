// generated by Textmapper; DO NOT EDIT

package test

import (
	"github.com/inspirer/textmapper/parsers/test/token"
)

const tmNumClasses = 42

type mapRange struct {
	lo         rune
	hi         rune
	defaultVal uint8
	val        []uint8
}

func mapRune(c rune) int {
	lo := 0
	hi := len(tmRuneRanges)
	for lo < hi {
		m := lo + (hi-lo)/2
		r := tmRuneRanges[m]
		if c < r.lo {
			hi = m
		} else if c >= r.hi {
			lo = m + 1
		} else {
			i := int(c - r.lo)
			if i < len(r.val) {
				return int(r.val[i])
			}
			return int(r.defaultVal)
		}
	}
	return 40
}

// Latin-1 characters.
var tmRuneClass = []uint8{
	1, 2, 2, 2, 2, 2, 2, 2, 2, 3, 4, 5, 5, 6, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 3, 2, 7, 2, 2, 8, 2, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18,
	18, 18, 18, 18, 18, 18, 18, 18, 18, 19, 2, 2, 2, 20, 2, 2, 21, 21, 21, 21,
	21, 21, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22, 22,
	22, 22, 23, 24, 25, 26, 27, 28, 2, 29, 29, 29, 29, 30, 31, 32, 32, 32, 32,
	32, 32, 32, 32, 33, 32, 34, 32, 35, 36, 37, 32, 32, 32, 32, 32, 38, 2, 39, 2,
	2, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
	40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40, 40,
}

const tmRuneClassLen = 256
const tmFirstRule = -5

var tmRuneRanges = []mapRange{
	{8232, 8234, 41, nil},
}

var tmStateMap = []int{
	0, 50,
}

var tmToken = []token.Token{
	1, 0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 1, 36, 37, 39, 1, 39,
	39, 2,
}

var tmLexerAction = []int8{
	-6, 49, -5, 49, 49, -5, 49, 48, 43, 42, 41, 40, -5, 39, 38, 36, 33, 30, 28,
	27, -5, 26, 26, 19, 18, 17, 16, 14, 13, 26, 26, 26, 26, 26, 26, 26, 3, 26, 2,
	1, -5, -5, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21,
	-21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21,
	-21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -21, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20,
	-20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -20, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -1, -9, -9, 26, -9, -9, 26, 26, 26,
	-9, -9, -9, -9, 26, 26, 4, 26, 26, 26, 26, 26, 26, 26, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -1, -9, -9, 26, -9,
	-9, 26, 26, 26, -9, -9, -9, -9, 26, 26, 26, 26, 26, 26, 26, 5, 26, 26, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -1,
	-9, -9, 26, -9, -9, 26, 26, 26, -9, -9, -9, -9, 26, 26, 26, 26, 26, 26, 26,
	26, 6, 26, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -2, -9, -9, 26, -9, -9, 26, 26, 26, -9, -9, -9, -9, 26, 26, 26,
	7, 26, 26, 26, 26, 26, 26, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -1, -9, -9, 26, -9, -9, 26, 26, 26, -9, -9, -9,
	-9, 26, 26, 26, 26, 26, 8, 26, 26, 26, 26, -9, -9, -9, -9, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -1, -9, -9, 26, -9, -9, 26, 26,
	26, -9, -9, -9, -9, 26, 26, 26, 26, 26, 9, 26, 26, 26, 26, -9, -9, -9, -9,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -2, -9, -9, 26,
	-9, -9, 26, 26, 26, -9, -9, -9, -9, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
	-9, -9, -9, -9, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	10, -5, -5, 26, -5, 11, 26, 26, 26, -5, -5, -5, -5, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 26, -5, -5, -5, -5, -43, -43, -43, -43, -43, -43, -43, -43, -43,
	-43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43,
	-43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43, -43,
	-43, -43, -43, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	12, -5, -5, 26, -5, -5, 26, 26, 26, -5, -5, -5, -5, 26, 26, 26, 26, 26, 26,
	26, 26, 26, 26, -5, -5, -5, -5, -34, -34, -34, -34, -34, -34, -34, -34, -34,
	-34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34,
	-34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34, -34,
	-34, -34, -34, -5, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, -5, -5, -5, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 15, -5, -5, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	-10, -10, -10, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25,
	-25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25,
	-25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25, -25,
	-33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33,
	-33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33,
	-33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -33, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24, -24,
	-24, -24, -24, -24, -24, -24, -24, -24, -24, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, 25, 25, 25, -5, 20, -5,
	-5, -5, 25, 25, 25, 25, 25, 25, 25, 25, 25, -5, -5, -5, -5, -41, -41, -41,
	-41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41,
	-41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41,
	-41, -41, -41, -41, 21, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41,
	-41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, 22, -41, -41, 22,
	-41, -41, -41, -41, -41, -41, -41, 22, 22, 22, -41, -41, -41, -41, -41, -41,
	-41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41,
	-41, -41, -41, -41, -41, -41, -41, 23, -41, -41, 23, -41, -41, -41, -41, -41,
	-41, -41, 23, 23, 23, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41,
	-41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41,
	-41, -41, 24, -41, -41, 24, -41, -41, -41, -41, -41, -41, -41, 24, 24, 24,
	-41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41,
	-41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, -41, 25, -41,
	-41, 25, -41, -41, -41, -41, -41, -41, -41, 25, 25, 25, -41, -41, -41, -41,
	-41, -41, -41, -41, -41, -41, -40, -40, -40, -40, -40, -40, -40, -40, -40,
	-40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, -40, 25, 25, 25, -40,
	20, -40, -40, -40, 25, 25, 25, 25, 25, 25, 25, 25, 25, -40, -40, -40, -40,
	-9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -9, -1, -9, -9, 26,
	-9, -9, 26, 26, 26, -9, -9, -9, -9, 26, 26, 26, 26, 26, 26, 26, 26, 26, 26,
	-9, -9, -9, -9, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29,
	-29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29,
	-29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29, -29,
	29, -11, -11, -11, 29, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, 28, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11,
	-11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -11, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, 32, -5, -5, -5, -5, 31, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -8, 31, 31, 31, -8,
	31, -8, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31,
	31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, 31, -8, -44,
	-44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44,
	-44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44,
	-44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -44, -26, -26, -26, -26,
	-26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -3, -26, -26,
	-26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26, -26,
	-26, -26, -26, -26, -26, -26, -26, -26, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, 35, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27, -27,
	-27, -27, -27, -27, -27, -27, -27, -27, -30, -30, -30, -30, -30, -30, -30,
	-30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, 37, -30,
	-30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30, -30,
	-30, -30, -30, -30, -30, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31,
	-31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31,
	-31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31, -31,
	-31, -31, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28,
	-28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28,
	-28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -28, -32,
	-32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32,
	-32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32,
	-32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -32, -23, -23, -23, -23,
	-23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23,
	-23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23, -23,
	-23, -23, -23, -23, -23, -23, -23, -23, -22, -22, -22, -22, -22, -22, -22,
	-22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22,
	-22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22, -22,
	-22, -22, -22, -22, -22, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39,
	-39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39,
	-39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39, -39,
	-39, -39, -5, -5, -5, 43, 43, 43, 43, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, 44, -5,
	-5, -5, -5, -5, -5, -5, 45, -5, -5, -5, 45, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, 46,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, 46, 46,
	46, 46, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5, -5,
	-5, -5, -5, -5, -5, -5, -5, -5, -5, -5, 47, -5, -5, -5, -5, -5, -5, -5, -4,
	-37, -37, -37, -4, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37,
	-37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37,
	-37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -37, -38, -38, -38, -38,
	-38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38,
	-38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38, -38,
	-38, -38, -38, -38, -38, -38, -38, -38, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7,
	-7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, -7, 56, 55, 55, 55, 55,
	55, 55, 55, 55, 55, 55, 55, 53, 55, 55, 55, 55, 51, 55, 55, 55, 55, 55, 55,
	55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, -48,
	-48, -48, -48, -48, -48, -48, -48, -48, -48, -48, -48, 52, -48, -48, -48,
	-48, -48, -48, -48, -48, -48, -48, -48, -48, -48, -48, -48, -48, -48, -48,
	-48, -48, -48, -48, -48, -48, -48, -48, -48, -48, -48, -46, -46, -46, -46,
	-46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46,
	-46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46, -46,
	-46, -46, -46, -46, -46, -46, -46, -46, -48, -48, -48, -48, -48, -48, -48,
	-48, -48, -48, -48, -48, -48, -48, -48, -48, -48, 54, -48, -48, -48, -48,
	-48, -48, -48, -48, -48, -48, -48, -48, -48, -48, -48, -48, -48, -48, -48,
	-48, -48, -48, -48, -48, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47,
	-47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47,
	-47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47, -47,
	-47, -47, -48, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, -48, 55, 55, 55,
	55, -48, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55, 55,
	55, 55, 55, 55, 55, 55, 55, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45,
	-45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45,
	-45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45, -45,
	-45, -45,
}

var tmBacktracking = []int{
	4, 12, // in Identifier
	4, 10, // in Identifier
	21, 34, // in '.'
	32, 45, // in multiline
}
