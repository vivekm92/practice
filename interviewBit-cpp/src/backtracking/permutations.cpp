#include <bits/stdc++.h>

using namespace std;

// https://www.interviewbit.com/problems/permutations/
class Solution {
	vector<vector<int> > permute(vector<int> &A);
};

void backtrackSolvePermute(size_t idx, vector<int> &A, vector<vector<int> > &allPermutations) {
	
	if (idx >= A.size()) {
		allPermutations.push_back(A);
		return ;
	}
	
	for (size_t i = idx; i < A.size(); i++) {
		swap(A[idx], A[i]);
		backtrackSolvePermute(idx + 1, A, allPermutations);
		swap(A[idx], A[i]);
	}
}

// Approach 1
// T(n) : O(n*n!) ; S(n) : O(n)
vector<vector<int> > solvePermute(vector<int> &A) {
	
	size_t n = A.size();
	vector<vector<int> > allPermutations;
	if (n == 0) return allPermutations;
	
	size_t idx = 0;
	backtrackSolvePermute(idx, A, allPermutations);
	return allPermutations; 
}

void backtrackSolvePermute1(vector<int> &curr, vector<int> &A, vector<vector<int> > &allPermutations) {
	
	if (curr.size() == A.size()) {
		allPermutations.push_back(curr);
		return ;
	}
	
	for (auto &a : A) {
		if (find(curr.begin(), curr.end(), a) == curr.end()) {
			curr.push_back(a);
			backtrackSolvePermute1(curr, A, allPermutations);
			curr.pop_back();
		}
	}
}

// Approach 2
// T(n) : O(n*n!) ; S(n) : O(n)
vector<vector<int> > solvePermute1(vector<int> &A) {
	
	size_t n = A.size();
	vector<vector<int > > allPermutations;
	if (n == 0) return allPermutations;
	
	vector<int> temp = {};
	backtrackSolvePermute1(temp, A, allPermutations);
	return allPermutations;
}

vector<vector<int> > Solution::permute(vector<int> &A) {
	return solvePermute1(A);
}

// Driver Code for testing
int main() {

    return 0;
}
