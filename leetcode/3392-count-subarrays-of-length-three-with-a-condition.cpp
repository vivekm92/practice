#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/count-subarrays-of-length-three-with-a-condition/description/
class Solution {
private:
    // Approach 1
    // using basic-simulation
    // T(n) : O(n) ; S(n) : O(1)
    int solveCountSubArrays(vector<int> &nums) {

        int n = nums.size(), count = 0;
        for (int i = 1; i < n-1; i++) {
            if ((nums[i-1] + nums[i+1]) * 2 == nums[i]) count++;
        }

        return count;
    }
public:
    int countSubarrays(vector<int>& nums) {
        return solveCountSubArrays(nums);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
