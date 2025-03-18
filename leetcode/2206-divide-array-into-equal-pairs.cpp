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

    // Approach 3
    // using constant arr
    // T(n) : O(n) ; S(n) : O(n)
    bool solveDivideArray2(vector<int>& nums) {

        int maxNum = *max_element(nums.begin(), nums.end());
        vector<bool> arr(maxNum+1, false);
        for (auto &num : nums) {
            arr[num] = arr[num] ^ true;
        }

        for (auto &num : nums) {
            if (arr[num] == true) return false;
        }

        return true;
    }

    // Approach 4
    // using hash-set
    // T(n) : O(n) ; S(n) : O(n)
    bool solveDivideArray3(vector<int> &nums) {

        unordered_set<int> us;
        for (auto &num : nums) {
            if (us.count(num) > 0) us.erase(num);
            else us.insert(num);
        }

        return us.size() == 0;
    }
public:
    bool divideArray(vector<int>& nums) {
        return solveDivideArray2(nums);
    }
};

// Driver Code for testing
int main() {

    vector<int> arr = {3,2,3,2,2,2};
    Solution *sol = new Solution();
    cout << sol->divideArray(arr) << "\n";

    return 0;
}
