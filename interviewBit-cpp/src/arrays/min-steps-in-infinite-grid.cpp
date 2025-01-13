#include <iostream>
#include <vector>

using namespace std;

// T(n) : O(n) ; S(n) : O(1)
int solveMinStepsInInfiniteGrid(vector<int>& A, vector<int>& B) {

    int n=A.size(), m=B.size();
    if (n != m) return -1;
    else if (n==0) return 0;
    int x=A[0], y=B[0], steps=0;
    for (int i=1; i<n; i++) {
        int xi= x>A[i] ? x-A[i] : A[i]-x;
        int yi= y>B[i] ? y-B[i] : B[i]-y;
        steps += max(xi, yi);
        x=A[i];
        y=B[i];
    }

    return steps;
}

// Driver Code for testing
int main() {

    int n, t;
    cin >> n;
    vector<int> arr;
    while (n--) {
        cin >> t;
        arr.push_back(t);
    }
    cin >> n;
    vector<int> brr;
    while (n--) {
        cin >> t;
        brr.push_back(t);
    }

    int minStepsCount = solveMinStepsInInfiniteGrid(arr, brr);
    cout << minStepsCount << "\n";

    return 0;
}