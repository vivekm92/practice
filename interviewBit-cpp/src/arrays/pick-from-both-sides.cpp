#include <iostream>
#include <vector> 

using namespace std;

/*
    IB : https://www.interviewbit.com/problems/pick-from-both-sides/
    LC : https://leetcode.com/problems/maximum-points-you-can-obtain-from-cards/description/
*/ 

// Brute-Force Solution 1:
// 
int solveBruteForcePickFromBothSides(vector<int>& A, int B) {

    return -1;
}

// Optimised Approach 1: 
// T(n) : O(n) ; S(n) : O(1) 
int solvePickFromBothSides(vector<int>& A, int B) {

    int tempSum=0, n=A.size();
    for (int i=0; i<B; i++) {
        tempSum += A[i];
    }

    int maxSum=tempSum; 
    for (int i=0; i<B; i++) {
        tempSum = tempSum + A[n-1-i] - A[B-1-i];
        maxSum = max(maxSum, tempSum);
    }

    return maxSum;
}

// Optimised Approach 2: 
// T(n) : O(n) ; S(n) : O(1)
int solvePickFromBothSides1(vector<int>& A, int B) {
    
}


// Driver Code for testing
int main() {

    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);

    int N, T, B;
    cin >> N;
    vector<int> arr;
    while (N--) {
        cin >> T;
        arr.push_back(T);
    }
    cin >> B;

    int maxSum = solvePickFromBothSides(arr, B);
    cout << maxSum << "\n";

    return 0;
}
