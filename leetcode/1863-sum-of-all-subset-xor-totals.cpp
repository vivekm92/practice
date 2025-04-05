#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/sum-of-all-subset-xor-totals/description/
class Solution {
private:
    void generateAllSubsets(vector<int> &nums, int idx, vector<vector<int> >& allSubsets, vector<int> &subset) {
        if (idx == nums.size()) {
            allSubsets.push_back(subset);
            return;
        }

        subset.push_back(nums[idx]);
        generateAllSubsets(nums, idx + 1, allSubsets, subset);
        subset.pop_back();
        generateAllSubsets(nums, idx + 1, allSubsets, subset);
    }
    // Approach 1
    // using backtracking to find all subsets
    // T(n) : O(n*2^n) ; S(n) : O(n*2^n)
    int solveSubsetXORSum(vector<int> &nums) {

        vector<int> subset;
        vector<vector<int> > allSubsets;
        generateAllSubsets(nums, 0, allSubsets, subset);

        int allXorTotalSum = 0;
        for (auto &vec : allSubsets) {
            int xorTotalSum = 0;
            for (auto &v : vec) {
                xorTotalSum ^= v;
            }
            allXorTotalSum += xorTotalSum;
        }
        return allXorTotalSum;
    }

    void findXorTotalSum(vector<int> &nums, int idx, int xorSum, vector<int> &subset) {

        if (idx == nums.size()) {
            subset.push_back(xorSum);
            return;
        }

        xorSum ^= nums[idx];
        findXorTotalSum(nums, idx + 1, xorSum, subset);
        xorSum ^= nums[idx];
        findXorTotalSum(nums, idx + 1, xorSum, subset);
    }

    // Approach 2
    // using optimised backtracking
    // T(n) : O(2^n) ; S(n) : O(n)
    int solveSubsetXORSum1(vector<int> &nums) {

        vector<int> subset;
        findXorTotalSum(nums, 0, 0, subset);

        int xorTotalSum = 0;
        for (auto &xorSum : subset) xorTotalSum += xorSum;

        return xorTotalSum;
    }

    // Approach 3
    // using optimised backtracking
    // T(n) : O(2^n) ; S(n) : O(n)
    int solveSubsetXORSum2(vector<int> &nums, int idx = 0, int currXor = 0) {
        if (idx == nums.size()) {
            return currXor;
        }

        int include = solveSubsetXORSum2(nums, idx + 1, currXor ^ nums[idx]);
        int exclude = solveSubsetXORSum2(nums, idx + 1, currXor);

        return include + exclude;
    }

    void solveSubsetXORSumHelper3(vector<int> &nums, int idx, int sum, int &total) {

        total += sum;
        for (int i = idx; i < nums.size(); i++) {
            solveSubsetXORSumHelper3(nums, i + 1, sum ^ nums[i], total);
        }
    }
    // Approach 4
    // using optimised backtracking
    // T(n) : O(2^n) ; S(n) : O(n)
    int solveSubsetXORSum3(vector<int> &nums) {
        int total = 0;
        solveSubsetXORSumHelper3(nums, 0, 0, total);
        return total;
    }

    // Approach 5
    // using bit-wise
    // T(n) : O(n) ; S(n) : O(1)
    int solveSubsetXORSum4(vector<int> &nums) {

        int total = 0;
        for (auto &num : nums) total |= num;

        return total << ( nums.size() - 1);

    }
public:
    int subsetXORSum(vector<int>& nums) {
        return solveSubsetXORSum2(nums);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
