package arrayProblems

import (
	"fmt"
	"reflect"
	"testing"
)

type isDigitTestCase struct {
	A        string
	Expected bool
}

var isDigitTestCases = []isDigitTestCase{
	{
		"dig1-8-1-5-1",
		true,
	},
	{
		"let1-art-can",
		false,
	},
	{
		"dig2-3-6",
		true,
	},
	{
		"let2-own-kit-dig",
		false,
	},
	{
		"let3-art-zero",
		false,
	},
}

func TestIsDigitLogs(t *testing.T) {

	for idx, test := range isDigitTestCases {
		if output := IsDigitLogs(test.A); output != test.Expected {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}

}

type reOrderLogsTestCase struct {
	A        []string
	Expected []string
}

var reOrderLogsTestCases = []reOrderLogsTestCase{
	{
		[]string{"dig1-8-1-5-1", "let1-art-can", "dig2-3-6", "let2-own-kit-dig", "let3-art-zero"},
		[]string{"let1-art-can", "let3-art-zero", "let2-own-kit-dig", "dig1-8-1-5-1", "dig2-3-6"},
	},
	{
		[]string{"a1-9-2-3-1", "g1-act-car", "zo4-4-7", "ab1-off-key-dog", "a8-act-zoo"},
		[]string{"g1-act-car", "a8-act-zoo", "ab1-off-key-dog", "a1-9-2-3-1", "zo4-4-7"},
	},
}

func TestReorderLogs(t *testing.T) {

	for idx, test := range reOrderLogsTestCases {
		if output := ReorderLogs(test.A); !reflect.DeepEqual(output, test.Expected) {
			fmt.Println(test.A, test.Expected, output)
			t.Errorf("Failed %v testCase \nGOT : %v \nEXPECTED : %v\n", idx, output, test.Expected)
		}
	}
}
