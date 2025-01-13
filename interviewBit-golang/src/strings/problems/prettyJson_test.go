package stringProblems

import (
	"reflect"
	"testing"
)

func TestPrettyJson(t *testing.T) {
	tests := []struct {
		A        string
		Expected []string
	}{
		{"{A:\"B\",C:{D:\"E\",F:{G:\"H\",I:\"J\"}}}", []string{"{", "\tA:\"B\",", "\tC:", "\t{", "\t\tD:\"E\",", "\t\tF:", "\t\t{", "\t\t\tG:\"H\",", "\t\t\tI:\"J\"", "\t\t}", "\t}", "}"}},
	}

	for _, test := range tests {
		result := PrettyJson(test.A)
		if !reflect.DeepEqual(result, test.Expected) {
			t.Errorf("PrettyJson(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}

}
