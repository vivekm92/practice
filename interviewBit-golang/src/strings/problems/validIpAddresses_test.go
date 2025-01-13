package stringProblems

import (
	"reflect"
	"sort"
	"testing"
)

func TestIIsValidIPDigit(t *testing.T) {
	tests := []struct {
		A        string
		B, C     int
		Expected bool
	}{
		{"25525511135", 0, 2, true},
		{"101023", 0, 2, true},
		{"0000", 0, 2, false},
	}

	for _, test := range tests {
		result := IsValidIPDigit(test.A, test.B, test.C)
		if result != test.Expected {
			t.Errorf("IsValidIPDigit(%v, %v, %v) = %v ; want %v", test.A, test.B, test.C, result, test.Expected)
		}
	}
}

func TestRestoreIpAddresses(t *testing.T) {
	tests := []struct {
		A        string
		Expected []string
	}{
		{"0000", []string{"0.0.0.0"}},
		{"25525511135", []string{"255.255.111.35", "255.255.11.135"}},
		{"101023", []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}},
	}

	sortSlice := func(A []string) []string {
		sort.Slice(A, func(i, j int) bool {
			return A[i] < A[j]
		})
		return A
	}

	for _, test := range tests {
		result := RestoreIpAddresses(test.A)
		sortedResult, sortedTestCase := sortSlice(result), sortSlice(test.Expected)
		if !reflect.DeepEqual(sortedResult, sortedTestCase) {
			t.Errorf("RestoreIpAddresses(%v) = %v ; want %v", test.A, result, test.Expected)
		}
	}
}
