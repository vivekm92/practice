package stringProblems

import (
	"reflect"
	"testing"
)

func TestJustifiedText(t *testing.T) {
	tests := []struct {
		A        []string
		B        int
		Expected []string
	}{
		{[]string{"This", "is", "an", "example", "of", "text", "justification."}, 16, []string{"This    is    an", "example  of text", "justification.  "}},
	}

	for _, test := range tests {
		result := JustifiedText(test.A, test.B)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("JustifiedText(%v, %v) = %v; want %v", test.A, test.B, result, test.Expected)
		}
	}
}
