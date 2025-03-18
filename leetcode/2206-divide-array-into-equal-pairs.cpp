#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/divide-array-into-equal-pairs/description/
class Solution {
private:
    // Approach 1
    // using hash-map
    // T(n) : O(n) ; S(n) : O(n)
    bool solveDivideArray(vector<int>& nums) {

        unordered_map<int, int> um;
        for (auto &num : nums) um[num]++;

        for (auto it = um.begin(); it != um.end(); it++) {
            if (it->second % 2 != 0) return false;
        }

        return true;
    }

    // Approach 2
    // using sorting
    // T(n) : O(nlogn) ; S(n) : O(1)
    bool solveDivideArray1(vector<int>& nums) {


        sort(nums.begin(), nums.end());
        int n = nums.size(), start = 0, end = 0;
        while (end < n) {
            while (start < n && nums[start] == nums[end]) {
                start++;
            }
            int cnt = start - end;
            if (cnt % 2 != 0) return false;
            end = start;
        }

        return true;
    }
public:
    bool divideArray(vector<int>& nums) {
        return solveDivideArray(nums);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
