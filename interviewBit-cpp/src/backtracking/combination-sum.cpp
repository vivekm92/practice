#include <iostream>
#include <vector>

using namespace std;

void combinationSumHelper(const vector<int> &A, int B, int idx, vector<int> &temp, vector<vector<int> > &res) {

    if (B == 0) {
        res.push_back(temp);
        return ;
    }

    while (idx < A.size() && A[idx] <= B) {
        temp.push_back(A[idx]);
        combinationSumHelper(A, B - A[idx], idx, temp, res);
        idx++;
        temp.pop_back();
    }
}

// https://www.interviewbit.com/problems/combination-sum/
// T(n) : O(
vector<vector<int> > solveCombinationSum(vector<int> &A, int B) {

    int n = A.size();
    vector<vector<int> > res;
    if (n == 0) return res;

    sort(A.begin(), A.end());
    vector<int> temp;
    temp.push_back(A[0]);
    for (int i = 1; i < n; i++) {
        if (A[i] != A[i-1]) {
            temp.push_back(A[i]);
        }
    }
    vector<int> temp1;
    combinationSumHelper(temp, B, 0, temp1, res);

    return res;
}

// Driver Code for testing
int main() {

    vector<int> A {1, 2, 3};
    vector<vector<int> > res = solveCombinationSum(A, 4);

    for (auto &r : res) {
        for (auto &a : r) {
            cout << a << " ";
        }
        cout << "\n";
    }

    return 0;
}
