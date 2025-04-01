#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/put-marbles-in-bags/description/
class Solution {
private:
    // Approach 1
    // using sort
    // T(n) : O(nlogn) ; S(n) : O(n)
    long long solvePutMarbles(vector<int> &weights, int k) {

        int n = weights.size();
        vector<int> pairCosts;
        for (int i = 0; i < n - 1; i++) {
            pairCosts.push_back(weights[i] + weights[i+1]);
        }

        long long minScore = 0l, maxScore = 0l;
        int m = pairCosts.size();
        sort(pairCosts.begin(), pairCosts.end());
        for (int i = 0; i < k-1; i++) {
            minScore += pairCosts[i];
            maxScore += pairCosts[m-1-i];
        }

        return maxScore - minScore;
    }

    // Approach 2
    // faster Implementation
    // T(n) : O(n) ; S(n) : O(1)
    long long solvePutMarbles1(vector<int> &weights, int k) {

        int n = weights.size() - 1;
        for (int i = 0; i < n; i++) {
            weights[i] += weights[i+1];
        }
        weights.pop_back();

        long long res = 0;
        nth_element(weights.begin(), weights.begin() + k - 1, weights.end());
        for (int i = 0; i < k-1; i++) {
            res -= weights[i];
        }
        nth_element(weights.begin(), weights.end() - k + 1, weights.end());
        for (int i = 0; i < k-1; i++) {
            res += weights[n-i-1];
        }

        return res;
    }
public:
    long long putMarbles(vector<int>& weights, int k) {
        return solvePutMarbles1(weights, k);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
