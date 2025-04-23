#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/concatenation-of-array/description/
class Solution {
private:
    // Approach 1
    // using basic simulation
    // T(n) : O(n) ; S(n) : O(n)
    vector<int> solveGetConcatenation(vector<int> &nums) {

        int n = nums.size();
        for (int i = 0; i < n; i++) nums.push_back(nums[i]);
        return nums;
    }
public:
    vector<int> getConcatenation(vector<int>& nums) {
        return solveGetConcatenation(nums);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
