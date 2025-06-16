#include <iostream>
#include <vector>

using namespace std;

void permutationsRecursive(vector<int> &A, int idx, vector<vector<int> > &allPermutations) {
    if (idx == A.size()) {
        allPermutations.push_back(A);
        return ;
    }

    for (int i = idx; i < A.size(); i++) {
        swap(A[idx], A[i]);
        permutationsRecursive(A, idx + 1, allPermutations);
        swap(A[idx], A[i]);
    }

}

// https://www.interviewbit.com/problems/permutations/
// T(n) : O(n*n!) ; S(n) : O(n)
vector<vector<int> > solvePermutations(vector<int> &A) {

    int n = A.size();
    vector<vector<int> > allPermutations;
    if (n == 0) return allPermutations;

    permutationsRecursive(A, 0, allPermutations);

    return allPermutations;
}

// Driver Code for testing
int main() {

    return 0;
}
