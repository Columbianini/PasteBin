package assert

import (
	"testing"
)

func Equal[T comparable](t *testing.T, actual, expected T) {
	t.Helper() // indicate that this function is a test helper, which will report the report the filename and line number of the code which called the Equal() function in the output

	if actual != expected {
		t.Errorf("got: %v, want: %v", actual, expected)
	}
}
