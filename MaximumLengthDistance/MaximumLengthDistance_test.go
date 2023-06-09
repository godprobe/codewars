package kata

import "testing"

/*
DESCRIPTION:
You are given two arrays a1 and a2 of strings. Each string is composed with letters from a to z.
Let x be any string in the first array and y be any string in the second array.

Find max(abs(length(x) − length(y)))

If a1 and/or a2 are empty return -1 in each language except in Haskell (F#) where you will return Nothing (None).

Example:
a1 = ["hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"]
a2 = ["cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"]
mxdiflg(a1, a2) --> 13
Bash note:
input : 2 strings with substrings separated by ,
output: number as a string
*/

var (
	got, want int
	cases     = []struct {
		a1, a2 []string
		num    int
	}{
		{
			[]string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"},
			[]string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"},
			13,
		},
		{
			[]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"},
			[]string{"bbbaaayddqbbrrrv"},
			10,
		},
		{
			[]string{"ccct", "tkkeeeyy", "ggiikffsszzoo", "nnngssddu", "rrllccqqqqwuuurdd", "kkbbddaakkk"},
			[]string{"tttxxxxxxgiiyyy", "ooorcvvj", "yzzzhhhfffaaavvvpp", "jjvvvqqllgaaannn", "tttooo", "qmmzzbhhbb"},
			14,
		},
		{
			[]string{},
			[]string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"},
			-1,
		},
		{
			[]string{},
			[]string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"},
			-1,
		}, {
			[]string{},
			[]string{},
			-1,
		},
	}
)

func TestMxDifLg(t *testing.T) {
	for _, val := range cases {
		got = MxDifLg(val.a1, val.a2)
		want = val.num
		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	}
}
