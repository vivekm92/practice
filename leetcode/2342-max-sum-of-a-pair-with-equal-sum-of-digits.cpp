#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

// https://leetcode.com/problems/max-sum-of-a-pair-with-equal-sum-of-digits/description/
class Solution {
    int findDigitsSum(int n) {
        int sum_ = 0;
        while (n>0) {
            sum_ += (n%10);
            n /= 10;
        }
        return sum_;
    }
    
    // T(n) : O(nlogn) ; S(n) : O(n)
    int solveMaximumSum(vector<int>& nums) {
        unordered_map<int, vector<int> > um;
        for (auto &num : nums) {
            int digitSum=findDigitsSum(num);
            if (um.find(digitSum) == um.end()) {
                um[digitSum] = {num};
            } else {
                um[digitSum].push_back(num);
            }
        }
        
        sort(nums.begin(), nums.end(), greater<int>());
        int maxSum=-1;
        for (auto &[digitSum, vec] : um) {
            if (vec.size() < 2) continue;
            maxSum = max(maxSum, vec[0]+vec[1]);
        }
        
        return maxSum;
    }
    
    
public:
    int maximumSum(vector<int>& nums) {
        return solveMaximumSum(nums);
    }
};


// Driver Code for testing
int main() {
    
    return 0;
}
