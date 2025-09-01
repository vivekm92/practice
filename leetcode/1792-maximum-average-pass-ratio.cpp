#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/maximum-average-pass-ratio/description/
class Solution {
private:
    // Approach 1
    // using priority queue
    // T(n) : O(nlogn) ; S(n) : O(n)
    double solveMaxAverageRatio(vector<vector<int> > &classes, int extraStudents) {

        priority_queue<pair<double, pair<int, int> > > pq;
        for (auto &c : classes) {
            double pass = c[0], total = c[1];
            double delta = ((pass+1)/(total+1)) - (pass/total);
            pq.push({delta, {pass, total}});
        }

        while (extraStudents--) {
            auto [_, c] = pq.top();
            pq.pop();
            double pass = c.first + 1, total = c.second + 1;
            double delta = ((pass+1)/(total+1)) - (pass/total);
            pq.push({delta, {pass, total}});
        }

        double averageSum = 0.0;
        while (pq.size() > 0) {
            auto [_, c] = pq.top();
            pq.pop();
            double pass = c.first, total = c.second;
            averageSum += (pass / total);
        }

        return averageSum / classes.size();
    }
public:
    double maxAverageRatio(vector<vector<int>>& classes, int extraStudents) {
        return solveMaxAverageRatio(classes, extraStudents);
    }
};

// Driver Code for testing.
int main() {

    return 0;
}
