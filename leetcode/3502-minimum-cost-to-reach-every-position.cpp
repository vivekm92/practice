#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/minimum-cost-to-reach-every-position/description/
class Solution {
private:
    // Approach 1
    // using min array
    // T(n) : O(n) ; S(n) : O(n)
    vector<int> solveMinCosts(vector<int> &cost) {

        vector<int> res;
        int tempCost = INT_MAX;
        for (auto &c : cost) {
            tempCost = min(tempCost, c);
            res.push_back(tempCost);
        }

        return res;
    }
public:
    vector<int> minCosts(vector<int>& cost) {
        return solveMinCosts(cost);
    }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    vector<int> arr = {1,2,3,4,5};
    vector<int> res = sol->minCosts(arr);
    for (auto &r : res) cout << r << " ";
    cout << "\n";

    return 0;
}
