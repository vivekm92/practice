#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/rearranging-fruits/description/
class Solution {
private:
    // Approach 1
    // using greedy approach
    // T(n) : O(nlogn) ; S(n) : O(n)
    long long solveMinCost(vector<int> &basket1, vector<int> &basket2) {

        int n = basket1.size(), min_val = INT_MAX;
        map<int, int> m;
        for (int i = 0; i < n; i++) {
            m[basket1[i]]++;
            m[basket2[i]]--;
            min_val = min(min_val, basket1[i]);
            min_val = min(min_val, basket2[i]);
        }

        vector<int> temp;
        for (auto it = m.begin(); it != m.end(); it++) {
            int k = abs(it->second);
            if (k % 2 != 0) return -1;
            k = k / 2;
            while (k--) {
                temp.push_back(it->first);
            }
        }

        n = temp.size();
        long long cost = 0;
        for (int i = 0; i < n / 2; i++) {
            cost += min(temp[i], 2*min_val);
        }

        return cost;
    }
public:
    long long minCost(vector<int>& basket1, vector<int>& basket2) {
        return solveMinCost(basket1, basket2);
    }
};

// Driver Code for testing.
int main() {

    return 0;
}
