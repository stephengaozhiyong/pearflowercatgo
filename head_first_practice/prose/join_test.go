package prose

import (
	"fmt"
	"testing"
)

func TestOneElem(t *testing.T) {
	l := []string{"a"}
	got := JoinWithCommas(l)
	want := "a"

	if got != want {
		t.Error(fmt.Errorf(errorString(l, got, want)))
	}
}

func TestTwoElem(t *testing.T) {
	l := []string{"a", "b"}
	got := JoinWithCommas(l)
	want := "a and b"

	if got != want {
		t.Error(fmt.Errorf(errorString(l, got, want)))
	}
}

func TestThreeElem(t *testing.T) {
	l := []string{"a", "b", "c"}
	got := JoinWithCommas(l)
	want := "a, b, and c"

	if got != want {
		t.Error(fmt.Errorf(errorString(l, got, want)))
	}
}

type testData struct {
	l    []string
	want string
}

func TestBatch(t *testing.T) {
	tests := []testData{
		testData{l: []string{}, want: ""},
		testData{l: []string{"a"}, want: "a"},
		testData{l: []string{"a", "b"}, want: "a and b"},
		testData{l: []string{"a", "b", "c"}, want: "a, b, and c"},
	}
	for _, test := range tests {
		got := JoinWithCommas(test.l)
		if got != test.want {
			t.Error(errorString(test.l, got, test.want))
		}
	}
}

func errorString(l []string, got string, want string) string {
	return fmt.Sprintf("JoinWithCommas(%#v) = \"%s\", but want \"%s\".", l, got, want)
}
