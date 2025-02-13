#include <iostream>
#include <vector>
#include <queue>

#define ll long long

using namespace std;

// https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-ii/description/
class Solution {
private:
    // T(n) : O(nlogn) ; S(n) : O(n)
    int solveMinOperations(vector<int>& nums, int k) {

        priority_queue<ll, vector<ll>, greater<ll> > pq;
        for (auto &num : nums) {
            pq.push(num);
        }

        int count = 0;
        while (pq.size() > 1 && pq.top() < k) {
            ll m1 = pq.top();
            pq.pop();

            ll m2 = pq.top();
            pq.pop();

            pq.push(m1*2 + m2);
            count++;
        }

        return count;
    }
public:
    int minOperations(vector<int>& nums, int k) {
        return solveMinOperations(nums, k);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
