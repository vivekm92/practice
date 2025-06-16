#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/build-array-from-permutation/description/
class Solution {
private:
    // Approach 1
    // using basic-simulation
    // T(n) : O(n) ; S(n) : O(n)
    vector<int> solveBuildArray(vector<int> &nums) {
        
        int n = nums.size();
        vector<int> res(n);
        for (int i = 0; i < n; i++) {
            res[i] = nums[nums[i]];
        }
        
        return res;
    }
    
    // Approach 2
    // in-place solution
    // T(n) : O(n) ; S(n) : O(1)
    vector<int> solveBuildArray1(vector<int> &nums) {
        
        
    }
public:
    vector<int> buildArray(vector<int>& nums) {
        return solveBuildArray(nums);
    }
};

// Driver Code for testing
int main() {
    
    return 0;
}
