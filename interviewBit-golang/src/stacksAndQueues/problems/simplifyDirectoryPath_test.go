package stacksAndQueues

import "testing"

func TestSimplifyDirectoryPath(t *testing.T) {
	tests := []struct {
		A        string
		Expected string
	}{
		{"/home/", "/home"},
		{"../../b/../s/./d//c/e/..", "/s/d/c"},
		{"/a/./b/../../c/", "/c"},
	}

	for _, test := range tests {
		result := SimplifyDirectoryPath(test.A)
		if result != test.Expected {
			t.Errorf("SimplifyDirectoryPath(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
