#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

// https://leetcode.com/problems/count-number-of-bad-pairs/description/
class Solution {
private:
    // T(n) : O(n) ; S(n) : O(n)
    long long solvecountBadPairs(vector<int>& nums) {
        
        size_t n = nums.size();
        for (int i=0; i<n; i++) {
            nums[i] -= i;
        }
        
        unordered_map<int, int> um;
        for (auto &num : nums) um[num]++;
        
        long long goodPairs = 0, totalPairs = (n*(n-1))/2;
        for (auto &[diff, f] : um) {
            goodPairs += ((long long )f*((long long )f-1))/2;
        }
        
        return totalPairs - goodPairs;
    }
    
    // T(n) : O(n) ; S(n) : O(n)
    long long solvecountBadPairs1(vector<int>& nums) {
        
        size_t n = nums.size();
        long long badPairs = 0;
        unordered_map<int, int> um;
        for (int i=0; i<n; i++) {
            int diff = i - nums[i];
            int goodPairs = um[diff];
            badPairs += i - goodPairs;
            um[diff] = goodPairs + 1;
        }
        
        return badPairs;
    }
public:
    long long countBadPairs(vector<int>& nums) {
        return solvecountBadPairs1(nums);
    }
};

// Driver Code for testing
int main() {
        
    return 0;
}
