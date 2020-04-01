package words

import (
	"reflect"
	"testing"
)

type TestRunner struct {
	t *testing.T
}

func (r *TestRunner) assertLongestConsec(want string, words []string, by int) {
	got := LongestConsec(words, by)
	if got != want {
		r.t.Errorf("FAIL: want: '%s', got: '%s'", want, got)
		return
	}
	r.t.Logf("PASS: '%s'", want)
}

func (r *TestRunner) assertCombine(message string, want []string, words []string, by int) {
	got := combine(words, by)
	if !reflect.DeepEqual(got, want) {
		r.t.Errorf("FAIL: %s => '%s', got: '%s'", message, want, got)
		return
	}
	r.t.Logf("PASS: %s", message)
}

// noinspection SpellCheckingInspection
func TestLongestConsec(t *testing.T) {
	r := &TestRunner{t: t}

	r.assertLongestConsec("abigailtheta",
		[]string{"zone", "abigail", "theta", "form", "libe", "zas"},
		2)

	r.assertLongestConsec("oocccffuucccjjjkkkjyyyeehh",
		[]string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"},
		1)

	r.assertLongestConsec("", []string{}, 3)

	r.assertLongestConsec("wkppqsztdkmvcuwvereiupccauycnjutlvvweqilsfytihvrzlaodfixoyxvyuyvgpck",
		[]string{"itvayloxrp", "wkppqsztdkmvcuwvereiupccauycnjutlv", "vweqilsfytihvrzlaodfixoyxvyuyvgpck"},
		2)

	r.assertLongestConsec("wlwsasphmxxowiaxujylentrklctozmymu",
		[]string{"wlwsasphmxx", "owiaxujylentrklctozmymu", "wpgozvxxiu"}, 2)

	r.assertLongestConsec("",
		[]string{"zone", "abigail", "theta", "form", "libe", "zas"}, -2)

	r.assertLongestConsec("ixoyx3452zzzzzzzzzzzz",
		[]string{"it", "wkppv", "ixoyx", "3452", "zzzzzzzzzzzz"}, 3)

	r.assertLongestConsec("",
		[]string{"it", "wkppv", "ixoyx", "3452", "zzzzzzzzzzzz"}, 15)

	r.assertLongestConsec("",
		[]string{"it", "wkppv", "ixoyx", "3452", "zzzzzzzzzzzz"}, 0)
}

func TestCombine(t *testing.T) {
	r := &TestRunner{t}

	// noinspection SpellCheckingInspection
	r.assertCombine("Even list, even by",
		[]string{"headshoulder", "shoulderknees", "kneestoes"},
		[]string{"head", "shoulder", "knees", "toes"}, 2)

	// noinspection SpellCheckingInspection
	r.assertCombine("Odd list, even by",
		[]string{"headshoulder", "shoulderknees", "kneestoes", "toesfeet"},
		[]string{"head", "shoulder", "knees", "toes", "feet"}, 2)

	// noinspection SpellCheckingInspection
	r.assertCombine("Even list, odd by",
		[]string{"headshoulderknees", "shoulderkneestoes"},
		[]string{"head", "shoulder", "knees", "toes"}, 3)

	// noinspection SpellCheckingInspection
	r.assertCombine("Odd list, odd by",
		[]string{"headshoulderknees", "shoulderkneestoes", "kneestoesfeet"},
		[]string{"head", "shoulder", "knees", "toes", "feet"}, 3)

	r.assertCombine("Special case by 1",
		[]string{"head", "shoulder", "knees", "toes"},
		[]string{"head", "shoulder", "knees", "toes"}, 1)

	r.assertCombine("Special case by 0",
		[]string{},
		[]string{"head", "shoulder", "knees", "toes"}, 0)

	r.assertCombine("Special case empty input",
		[]string{},
		[]string{}, 1)

	r.assertCombine("Special case bad input, prevent crash",
		[]string{},
		[]string{"head", "shoulder", "knees", "toes"}, 10)

	r.assertCombine("Handle negative indices",
		[]string{},
		[]string{"head", "shoulder", "knees"}, -1)
}
