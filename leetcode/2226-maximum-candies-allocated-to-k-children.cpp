#include <iostream>

using namespace std;

// https://leetcode.com/problems/maximum-candies-allocated-to-k-children/description/
class Solution {
private:
    bool canAllocate(vector<int> &candies, long long k, int num) {

        long long maxChildren = 0;
        for (auto &candie : candies) {
            maxChildren += candie / num;
        }

        return maxChildren >= k;
    }

    // T(n) : O(nlogn) ; S(n) : O(1)
    int solveMaximumCandies(vector<int> &candies, long long k) {

        int l = 1, r = *max_element(candies.begin(), candies.end()) + 1;
        while (l < r) {
            int mid = l + (r - l) / 2;
            if (canAllocate(candies, k, mid)) {
                l = mid + 1;
            } else {
                r = mid;
            }
        }

        return l - 1;
    }
public:
    int maximumCandies(vector<int>& candies, long long k) {
        return solveMaximumCandies(candies, k);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
