#include <iostream>
#include <vector>
#include <queue>

using namespace std;

// https://leetcode.com/problems/find-time-required-to-eliminate-bacterial-strains/description/
class Solution {
private:
    // Approach 1
    // using greedy
    // T(n) : O(nlogn) ; S(n) : O(n)
    long long solveMinEliminationTime(vector<int> &timeReq, int splitTime) {

        priority_queue<long, vector<long>, greater<long> > pq {timeReq.begin(), timeReq.end()};
        while (pq.size() > 1) {
            long first = pq.top();
            pq.pop();
            long second = pq.top();
            pq.pop();
            pq.push(max(first, second) + splitTime);
        }

        return pq.top();
    }
public:
    long long minEliminationTime(vector<int>& timeReq, int splitTime) {
        return solveMinEliminationTime(timeReq, splitTime);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
