package assert

import "testing"

func Equal[K comparable](t *testing.T, actual, expected K) {
	t.Helper()

	if actual != expected {
		t.Errorf("got: %v; want: %v", actual, expected)
	}
}
