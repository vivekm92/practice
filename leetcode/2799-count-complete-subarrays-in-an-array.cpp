#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

// https://leetcode.com/problems/count-complete-subarrays-in-an-array/description/
class Solution {
private:
    // Approach 1
    // using hash-map
    // T(n) : O(n) ; S(n) : O(n)
    int solveCountCompleteSubarrays(vector<int> &nums) {

        unordered_map<int, int> um1, um2;
        for (auto &num : nums) um1[num]++;

        int n = nums.size(), distinctCount = um1.size();
        int count = 0, idx = 0;
        for (int i = 0; i < n; i++) {
            while (idx < n && um2.size() < distinctCount) {
                um2[nums[idx]]++;
                idx++;
            }
            if (um2.size() < distinctCount) break;
            count += (n - idx + 1);
            um2[nums[i]]--;
            if (um2[nums[i]] == 0) um2.erase(nums[i]);
        }

        return count;
    }

    // Approach 2
    // using hash-map
    // T(n) : O(n) ; S(n) : O(n)
    int solveCountCompleteSubarrays1(vector<int> &nums) {

        unordered_map<int, int> um;
        for (auto &num : nums) um[num]++;

        int n = nums.size();
        int countDistinct = um.size();
        um.clear();
        int idx = 0, count = 0;
        for (int i = 0; i < n; i++) {
            while (idx < n && um.size() < countDistinct) {
                um[nums[idx]]++;
                idx++;
            }
            if (um.size() < countDistinct) break;
            count += (n - idx + 1);
            um[nums[i]]--;
            if (um[nums[i]] == 0) um.erase(nums[i]);
        }

        return count;
    }
public:
    int countCompleteSubarrays(vector<int>& nums) {
        return solveCountCompleteSubarrays1(nums);
    }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    vector<int> arr {1,2,3,4,5};
    cout << sol->countCompleteSubarrays(arr) << "\n";

    vector<int> brr {5,5,5,5,5};
    cout << sol->countCompleteSubarrays(brr) << "\n";

    vector<int> crr {1,3,1,2,2};
    cout << sol->countCompleteSubarrays(crr) << "\n";

    vector<int> drr {5,5,5,5};
    cout << sol->countCompleteSubarrays(drr) << "\n";

    return 0;
}
