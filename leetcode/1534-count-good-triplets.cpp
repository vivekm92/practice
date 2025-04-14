#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/count-good-triplets/description/
class Solution {
private:
    // Approach 1
    // using brute-force
    // T(n) : O(n3) ; S(n) : O(1)
    int solveCountGoodTriplets(vector<int>& arr, int a, int b, int c) {

        int n = arr.size(), count = 0;
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                if (abs(arr[i] - arr[j]) <= a) {
                    for (int k = j + 1; k < n; k++) {
                        if (abs(arr[j] - arr[k]) <= b && abs(arr[i] - arr[k] <= c) count++;
                    }
                }
            }
        }

        return count;
    }
public:
    int countGoodTriplets(vector<int>& arr, int a, int b, int c) {
        return solveCountGoodTriplets(arr, a, b, c);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
