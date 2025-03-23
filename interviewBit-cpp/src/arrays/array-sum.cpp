#include <iostream>

using namespace std;

// https://www.interviewbit.com/problems/array-sum/

// Approach 1
// using two-pointes
// T(n) : O(n) ; S(n) ; O(n)
vector<int> arraySum(vector<int> &A, vector<int> &B) {

    int nA = A.size(), nB = B.size();
    if (nA < nB) return arraySum(B, A);

    int i = nA-1, j = nB-1, carry = 0;
    while (i >= 0 || j >= 0) {
        int v1 = i >= 0 ? A[i] : 0;
        int v2 = j >= 0 ? B[j] : 0;

        int sum = v1 + v2 + carry;
        A[i] = sum % 10;
        carry = sum / 10;

        i--;
        j--;
    }

    if (carry > 0) {
        vector<int> t;
        t.push_back(carry);
        for (auto &a : A) t.push_back(a);
        return t;
    }

    return A;
}

// Driver Code for testing
int main() {

    vector<int> arr = {1, 2, 3};
    vector<int> brr = {1};

    vector<int> crr = arraySum(arr, brr);
    for (auto &c : crr) cout << c << " ";
    cout << "\n";

    return 0;
}
