package lexier

import "testing"

func TestLexier(t *testing.T) {
	do := func(orig string, begin, length int, expect string) {
		got := substr(orig, begin, length)
		if got != expect {
			t.Errorf("(%s, %d, %d) got: %s, expect: %s", orig, begin, length, got, expect)
		}
	}

	orig := "01234"
	do(orig, 0, 5, "01234")
	do(orig, 0, 10, "01234")
	do(orig, 1, 2, "12")
	do(orig, 4, 5, "4")
	do(orig, 0, 0, "")
	do(orig, 2, 10, "234")
	do(orig, 10, 10, "")
}
