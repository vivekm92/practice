#include <iostream>
#include <vector>

using namespace std;

// https://www.interviewbit.com/problems/chips-factory/

// T(n) : O(n) ; S(n) : O(1)
vector<int> chipsFactory(vector<int> &A) {

    int n = A.size(), start = 0;
    for (int end = 0; end < n; end++) {
        if (A[end] > 0) {
            swap(A[start], A[end]);
            start++;
        }
    }

    return A;
}

// Driver Code for testing
int main() {

    vector<int> arr = {0, 0, 1, 2, 3};
    vector<int> res = chipsFactory(arr);

    for (auto &r : res) cout << r << " ";
    cout << "\n";


    return 0;
}
