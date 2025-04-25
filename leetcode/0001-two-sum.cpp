#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;


// https://leetcode.com/problems/two-sum/description/
class Solution {
private:
    // Approach 1
    // using brute-force
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

    // Approach 2
    // using hash-table
    // T(n) : O(n) ; S(n) : O(n)
    vector<int> solvetwoSum1(vector<int>& nums, int target) {

        int n = nums.size();
        unordered_map<int, int> um;
        for (int i = 0; i < n; i++) {
            if (um.find(target - nums[i]) != um.end()) {
                return {um[target - nums[i]], i};
            }
            um[nums[i]] = i;
        }

        return {-1, -1};
    }

    int bsearch(vector<pair<int, int> > &nums, int t, int l, int r) {

        while (l < r) {
            int mid = l + (r - l) / 2;
            if (nums[mid].first < t) l = mid + 1;
            else r = mid;
        }

        return l == nums.size() ? -1 : (nums[l].first == t ? nums[l].second : -1);
    }

    // Approach 3
    // using binary-search
    // T(n) : O(nlogn) ; S(n) : O(n)
    vector<int> solvetwoSum2(vector<int> &nums, int target) {

        int n = nums.size();
        vector<pair<int, int> > temp;
        for (int i = 0; i < n; i++) temp.push_back({nums[i], i});
        sort(temp.begin(), temp.end());

        for (int i = 0; i < n; i++) {
            int t = target - (temp[i]).first;
            int idx = bsearch(temp, t, i+1, n);
            if (idx != -1) return {temp[i].second, idx};
        }

        return {-1, -1};
    }
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        return solvetwoSum2(nums, target);
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
