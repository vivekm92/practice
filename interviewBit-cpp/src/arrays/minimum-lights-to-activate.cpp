#include <iostream>
#include <vector>

using namespace std;

// Optimised Solution : Approach 1
// T(n) : O(n) ; S(n) : O(1)
int solveMinimumLightsToActivate(vector<int>& A, int B) {

    int count=0, curr=0, n=A.size();
    while (curr < n) {
        int prev = curr-B+1 < 0 ? 0 : curr-B+1;
        int nxt = curr+B-1 >= n ? n-1 : curr+B-1;

        int idx=nxt;
        while (idx >= prev) {
            if (A[idx] == 1) {
                break;
            }
            idx--;
        }

        if (idx < prev) return -1;
        count++;
        curr = idx + B;
    }

    return count;
}

// Driver Code for testing
int main() {

    int t, n, b;
    cin >> n;
    vector <int> arr;
    while (n--) {
        cin >> t;
        arr.push_back(t);
    }
    cin >> b;

    int minLightsToActivate = solveMinimumLightsToActivate(arr, b);
    cout << minLightsToActivate << "\n";

    return 0;
}