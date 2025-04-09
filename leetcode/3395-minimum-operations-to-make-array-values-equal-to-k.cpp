#include <iostream>
#include <unordered_set>
#include <vector>
#include <map>

using namespace std;

// https://leetcode.com/problems/minimum-operations-to-make-array-values-equal-to-k/description/
class Solution {
private:
    // Approach 1
    // using ordered-map
    // T(n) : O(nlogn) ; S(n) : O(n)
    int solveMinOperations(vector<int> &nums, int k) {

        map<int, int> m;
        for (auto &num : nums) m[num]++;

        auto t = m.begin();
        if (k > t->first) return -1;

        int count = 0;
        for (auto it = m.begin(); it != m.end(); it++) {
            if (k < it->first) count++;
            k = it->first;
        }

        return count;
    }

    // Approach 2
    // using unordered_set
    // T(n) : O(n) ; S(n) : O(n)
    int solveMinOperations1(vector<int> &nums, int k) {

        unordered_set<int> us;
        for (int x : nums) {
            if (x < k) return -1;
            else if (x > k) us.insert(x);
        }

        return us.size();
    }
public:
    int minOperations(vector<int>& nums, int k) {
        return solveMinOperations1(nums, k);
    }
};


// Driver Code for testing
int main() {

    return 0;
}
