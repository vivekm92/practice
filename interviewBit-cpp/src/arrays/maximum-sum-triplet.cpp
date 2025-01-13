#include <iostream>
#include <vector> 

using namespace std;

// Brute-Force Approach:
// T(n) : O(n3) ; S(n) : O(1)
int bruteForcesolveMaximumSumTriplet(vector<int>& A) {

    int n=A.size(), maxSum=0, currSum=0;
    for (int i=0; i<n; i++) {
        for (int j=i+1; j<n; j++) {
            if (A[i]>=A[j]) continue;
            for (int k=j+1; k<n; k++) {
                if (A[j]>=A[k]) continue;
                currSum = A[i] + A[j] + A[k];
                maxSum = max(maxSum, currSum);
            }
        }
    }

    return maxSum;
}

// 
int solveMaximumSumTriplet(vector<int>& A) {

    return -1;
}

// Driver Code for testing
int main() {

    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);

    int t, n;
    cin >> n;
    vector<int> arr;
    while (n--) {
        cin >> t;
        arr.push_back(t);
    }

    int maxSumTriplet = solveMaximumSumTriplet(arr);
    cout << maxSumTriplet << "\n";

    return 0;
}