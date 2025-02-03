#include <iostream>

using namespace std;

// https://leetcode.com/problems/two-sum/description/
class Solution {
private:
    // brute-force
    // T(n) : O(n2) ; S(n) : O(1)
    vector<int> solvetwoSum(vector<int>& nums, int target) {
        int n=nums.size();
        for (int i=0; i<n; i++) {
            for (int j=i+1; j<n; j++) {
                if (nums[i]+nums[j] == target) {
                    return {i, j};
                }
            }
        }
        return {-1, -1};
    }
    
    // using hash-table
    // T(n) : O(n) ; S(n) : O(n)
    vector<int> solvetwoSum1(vector<int>& nums, int target) {
        int n=nums.size();
        unordered_map<int, int> um;
        for (int i=0; i<n; i++) {
            if (um.find(target-nums[i]) == um.end()) {
                um[nums[i]] = i;
            } else {
                return {um[target-nums[i]], i};
            }
        }
        return {-1, -1};
    }
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        return solvetwoSum1(nums, target);
    }
};

// Driver Code for testing
int main() {
    vector<int> arr {2, 7, 12, 11};
    Solution *sol = new Solution();
    vector<int> res = sol->twoSum(arr, 14);
    for (auto &r : res) cout << r << " ";
    cout << "\n";
    
    vector<int> res1 = sol->twoSum(arr, 9);
    for (auto &r : res1) cout << r << " ";
    cout << "\n";
    
    vector<int> res2 = sol->twoSum(arr, 13);
    for (auto &r : res2) cout << r << " ";
    cout << "\n";
    
    return 0;
}
