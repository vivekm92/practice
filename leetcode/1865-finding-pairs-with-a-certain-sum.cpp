#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/finding-pairs-with-a-certain-sum/description/
class FindSumPairs {
private:
    vector<int> arr1, arr2;
    unordered_map<int, int> um1, um2;
public:
    // T(n) : O(n) ; S(n) : O(n)
    FindSumPairs(vector<int>& nums1, vector<int>& nums2) {
        this->arr1 = nums1;
        this->arr2 = nums2;

        for (auto &num : this->arr1) this->um1[num]++;
        for (auto &num : this->arr2) this->um2[num]++;
    }

    // T(n) : O(1) ; S(n) : O(1)
    void add(int index, int val) {
        int curr_val = this->arr2[index];
        this->um2[curr_val]--;
        if (this->um2[curr_val] == 0) {
            this->um2.erase(curr_val);
        }
        this->um2[curr_val + val]++;
        this->arr2[index] = curr_val + val;
    }

    // T(n) : O(n) ; S(n) : O(1)
    int count(int tot) {

        int cnt = 0;
        for (auto it = um1.begin(); it != um2.end(); it++) {
            int target = tot - it->first;
            if (um2.find(target) != um2.end()) {
                cnt += um2[target] * it->second;
            }
        }

        return cnt;
    }
};

// Approach 2
class FindSumPairs1 {
private:
    vector<int> arr1, arr2;
    unordered_map<int, int> um2;
public:
    FindSumPairs1(vector<int> &nums1, vector<int> &nums2) {
        this->arr1 = nums1;
        this->arr2 = nums2;
        for (auto &num : this->arr2) this->um2[num]++;
    }

    void add(int index, int val) {
        int curr_val = this->arr2[index];
        this->um2[curr_val]--;
        if (this->um2[curr_val] == 0) {
            this->um2.erase(curr_val);
        }
        this->um2[curr_val + val]++;
        this->arr2[index] = curr_val + val;
    }

    int count(int tot) {

        int cnt = 0;
        for (auto &num : this->arr1) {
            int target = tot - num;
            if (um2.find(target) != um2.end()) {
                cnt += um2[target];
            }
        }
        return cnt;
    }
};

/**
 * Your FindSumPairs object will be instantiated and called as such:
 * FindSumPairs* obj = new FindSumPairs(nums1, nums2);
 * obj->add(index,val);
 * int param_2 = obj->count(tot);
 */

// Driver Code for testing
int main() {

    return 0;
}
