#include <iostream>
#include <vector>
#include <unordered_set>

using namespace std;

// https://leetcode.com/problems/contains-duplicate/description/
class Solution {
private:
    // Approach 1
    // using hash-set
    // T(n) : O(n) ; S(n) : O(n)
     bool solvecontainsDuplicate(vector<int>& nums) {
         unordered_set<int> us;
         for (auto &n : nums) {
             if (us.find(n) != us.end()) return true;
             us.insert(n);
         }
         return false;
     }

     // Approach 2
     // using sorting
     // T(n) : O(nlogn) ; S(n) : O(1)
    bool solvecontainsDuplicate1(vector<int> &nums) {

        sort(nums.begin(), nums.end());
        int n = nums.size();
        for (int i = 1; i < n; i++) {
            if (nums[i] == nums[i-1]) return true;
        }
        return false;
    }
public:
    bool containsDuplicate(vector<int>& nums) {
        return solvecontainsDuplicate1(nums);
    }
};

// Driver Code for testing
int main() {
    return 0;
}
