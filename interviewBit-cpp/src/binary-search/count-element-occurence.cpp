#include <iostream>

using namespace std;

// https://www.interviewbit.com/problems/count-element-occurence/

// Brute-Force Approach
// T(n) : O(n) ; S(n) : O(1)
int countElementOccurence(const vector<int>& A, int B) {
    int count = 0;
    for (auto &element : A) {
        if (element == B) count++;
    }
    return count;
}
/********************************************************************************/

// Optimised-Solution
// T(n) : O(logN) ; S(n) : O(1)
int countElementOccurence1(const vector<int>& A, int B) {
    
    int n = A.size();
    int l1 = 0, r1 = n;
    while (l1 < r1) {
        int mid = l1 + (r1 - l1) / 2;
        if (A[mid] <= B) {
            l1 = mid + 1;
        } else {
            r1 = mid;
        }
    }
    
    int l2 = 0, r2 = n;
    while (l2 < r2) {
        int mid = l2 + (r2 - l2) / 2;
        if (A[mid] < B) {
            l2 = mid + 1;
        } else {
            r2 = mid;
        }
    }
    
    return l1 - l2;
}
/********************************************************************************/

// Driver Code for testing
int main() {
    vector<int> arr {1, 2, 2, 7, 7, 7, 7, 8};
    cout << countElementOccurence1(arr, 1) << "\n";
    cout << countElementOccurence1(arr, 2) << "\n";
    cout << countElementOccurence1(arr, 3) << "\n";
    cout << countElementOccurence1(arr, 7) << "\n";
    cout << countElementOccurence1(arr, 8) << "\n";
    cout << countElementOccurence1(arr, 9) << "\n";
    return 0;
}
