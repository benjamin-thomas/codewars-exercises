package words

import (
	"reflect"
	"testing"
)

type TestRunner struct {
	t *testing.T
}

func (r *TestRunner) assertHello(wanted string, input string) {
	got := hello(input)
	if got != wanted {
		r.t.Errorf("FAIL: '%s', got: '%s'", wanted, got)
		return
	}
	r.t.Logf("PASS: %s", wanted)
}

func (r *TestRunner) assertCombine(message string, want []string, words []string, by int) {
	got := combine(words, by)
	if !reflect.DeepEqual(got, want) {
		r.t.Errorf("FAIL: %s => '%s', got: '%s'", message, want, got)
		return
	}
	r.t.Logf("PASS: %s", message)
}

func TestHello(t *testing.T) {
	r := &TestRunner{t}

	r.assertHello("Hello Ben", "Ben")
	r.assertHello("Hello Ewan", "Ewan")
	r.assertHello("Hello Dorian", "Dorian")
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
