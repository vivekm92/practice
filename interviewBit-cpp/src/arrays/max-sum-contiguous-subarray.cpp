#include <iostream>
#include <vector>
#include <cassert>

using namespace std;

// using kadane's Algo
// T(n) : O(n) ; S(n) : O(1)
int maxSumContiguousSubarray(const vector<int>& A) {

    assert(A.size() > 0);
    int currSum = A[0], maxSum = A[0];
    for (auto &num : A) {
        currSum = currSum + num > num ? currSum + num : num;
        maxSum = maxSum > currSum ? maxSum : currSum;
    }
    return maxSum;
}

// Driver Code for testing
int main() {

    vector<int> arr = {1,2,3,4,5,-10};
    cout << maxSumContiguousSubarray(arr) << "\n";

    vector<int> arr1 =  {-1, -2, -3, -4, -5, -6};
    cout << maxSumContiguousSubarray(arr1) << "\n";

    return 0;
}
