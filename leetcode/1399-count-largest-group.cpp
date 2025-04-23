#include <iostream>
#include <unordered_map>

using namespace std;

// https://leetcode.com/problems/count-largest-group/description/
class Solution {
private:
    // Approach
    // using hash-map
    // T(n) ; O(nlogn) ; S(n) : O(logn)
    int solveCountLargestGroup(int n) {

        auto digit_sum = [&](int n) {
            int s = 0;
            while (n > 0) {
                s += n % 10;
                n /= 10;
            }
            return s;
        };

        int maxSumCount = 0;
        unordered_map<int, int> um;
        for (int i = 1; i <= n; i++) {
            int s = digit_sum(i);
            um[s]++;
            maxSumCount = max(maxSumCount, um[s]);
        }

        int count = 0;
        for (auto &[_, val] : um) {
            if (val == maxSumCount) count++;
        }

        return count;
    }
public:
    int countLargestGroup(int n) {
        return solveCountLargestGroup(n);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
