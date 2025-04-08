#include <iostream>
#include <vector>
#include <unordered_set>

using namespace std;

// https://leetcode.com/problems/minimum-number-of-operations-to-make-elements-in-array-distinct/description/
class Solution {
private:
    // Approach 1
    // Brute-Force Approach ( Simulation )
    // T(n) ; O(n2) ; S(n) : O(n)
    int solveMinimumOperations(vector<int> &nums) {

        int  n = nums.size();
        auto checkUnique = [&](int start) {

            unordered_set<int> seen;
            for (int i = start; i < n; i++) {
                if (seen.count(nums[i]) > 0) return false;
                seen.insert(nums[i]);
            }
            return true;
        };

        int count = 0;
        for (int i = 0; i < n; i += 3) {
            if (checkUnique(i)) return count;
            count++;
        }

        return count;
    }

    // Approach 2
    // using hash-set
    // T(n) : O(n) ; S(n) : O(n)
    int solveMinimumOperations1(vector<int> &nums) {

        int n = nums.size();
        unordered_set<int> seen;
        for (int i = n - 1; i >= 0; i--) {
            if (seen.count(nums[i]) > 0) {
                return (i/3) + 1;
            }
            seen.insert(nums[i]);
        }
        return 0;
    }

    // Approach 3
    // using hash-map
    // T(n) : O(n) ; S(n)
    int solveMinimumOperations2(vector<int> &nums) {

        unordered_map<int, int> um;
        for (auto &num : nums) um[num]++;

        int idx = 0, count = 0, n = nums.size();
        while (idx < n) {
            if (n - idx == um.size()) return count;
            int i = idx;
            while (i < idx + 3) {
                if (i >= n) break;
                um[nums[i]]--;
                if (um[nums[i]] == 0) {
                    um.erase(nums[i]);
                }
                i++;
            }
            idx = i;
            count++;
        }

        return count;
    }
public:
    int minimumOperations(vector<int>& nums) {
        return solveMinimumOperations1(nums);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
