#include <iostream>
#include <vector>
#include <unordered_set>
#include <unordered_map>

using namespace std;

// https://leetcode.com/problems/intersection-of-two-arrays/description/
class Solution {
private:
    // using arrays as map
    // T(n) ; O(n + m) ; S(n) : O(n + m)
    vector<int> solveIntersection(vector<int> &nums1, vector<int> &nums2) {

        vector<int> t1(1001, 0), t2(1001, 0);
        for (auto &num1 : nums1) t1[num1]++;
        for (auto &num2 : nums2) t2[num2]++;

        vector<int> res;
        for (int i = 0; i < 1001; i++) {
            if (t1[i] > 0 && t2[i] > 0) res.push_back(i);
        }

        return res;
    }

    // using hash-set
    // T(n) ; O(n + m) ; S(n) : O(n + m)
    vector<int> solveIntersection1(vector<int> &nums1, vector<int> &nums2) {

        unordered_set<int> us1(nums1.begin(), nums1.end());
        unordered_set<int> us2(nums2.begin(), nums2.end());

        vector<int> res;
        for (auto it = us1.begin(); it != us1.end(); it++) {
            if (us2.find(*it) != us2.end()) res.push_back(*it);
        }

        return res;
    }

    // using hash-map
    // T(n) : O(n + m) ; S(n) : O(n + m)
    vector<int> solveIntersection2(vector<int> &nums1, vector<int> &nums2) {

        unordered_map<int, int> um;
        for (auto &num1 : nums1) um[num1] = 1;

        for (auto &num2 : nums2) {
            if (um.find(num2) != um.end()) {
                um[num2] = 0;
            }
        }

        vector<int> res;
        for (auto it = um.begin(); it != um.end(); it++) {
            if (it->second == 0) res.push_back(it->first);
        }

        return res;
    }
public:
    vector<int> intersection(vector<int>& nums1, vector<int>& nums2) {
        return solveIntersection2(nums1, nums2);
    }
};

// Driver Code for testing
int main() {


    return 0;
}
