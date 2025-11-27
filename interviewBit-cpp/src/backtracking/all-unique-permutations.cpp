#include <bits/stdc++.h>

using namespace std;

void allUniquePermutationsRecursive(vector<int> &A, int idx, vector<int> &temp, vector<vector<int> > &permutations, unordered_map<int, int> &um) {

    if (idx == A.size()) {
        permutations.push_back(temp);
        return ;
    }

    for (auto &[k, v] : um) {
        if (v == 0) continue;
        um[k]--;
        temp.push_back(k);
        allUniquePermutationsRecursive(A, idx + 1, temp, permutations, um);
        um[k]++;
        temp.pop_back();
    }

    return ;
}

// https://www.interviewbit.com/problems/all-unique-permutations/
// T(n) : O(n*n!) ; S(n) : O(n)
vector<vector<int> > solveAllUniquePermutations(vector<int> &A) {

    int n = A.size();
    vector<vector<int> > allPermutations;
    if (n == 0) return allPermutations;

    unordered_map<int, int> um;
    for (auto &a : A) um[a]++;

    vector<int> temp;
    allUniquePermutationsRecursive(A, 0, temp, allPermutations, um);
    return allPermutations;
}

// Driver Code for testing
int main() {


    return 0;
}
