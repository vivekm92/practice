#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/campus-bikes-ii/description/
class Solution {
private:
    int findDistance(vector<int> &workerPos, vector<int> &bikePos) {
        return abs(workerPos[0] - bikePos[0]) + abs(workerPos[1] - bikePos[1]);
    }

    void solveBackTrack(vector<vector<int> > &workers, int wIdx, vector<vector<int> > &bikes, int currentDistanceSum, int &smallestDistanceSum, vector<bool> &visited) {

        if (wIdx >= workers.size()) {
            smallestDistanceSum = min(smallestDistanceSum, currentDistanceSum);
            return ;
        }

        if (currentDistanceSum >= smallestDistanceSum) {
            return ;
        }

        for (int bIdx = 0; bIdx < bikes.size(); bIdx++) {
            if (!visited[bIdx]) {
                visited[bIdx] = true;
                int tempDistance = findDistance(workers[wIdx], bikes[bIdx]);
                solveBackTrack(workers, wIdx + 1, bikes, currentDistanceSum + tempDistance, smallestDistanceSum, visited);
                visited[bIdx] = false;
            }
        }
    }

	// Approach 1
	// using backtracking
	// T(n) : O(n!) ; S(n) : O(n)
    int solveAssignBikes(vector<vector<int> > &workers, vector<vector<int> > &bikes) {

        int smallestDistanceSum = INT_MAX;
        vector<bool> visited(10, false);
        solveBackTrack(workers, 0, bikes, 0, smallestDistanceSum, visited);
        return smallestDistanceSum;
    }
public:
    int assignBikes(vector<vector<int>>& workers, vector<vector<int>>& bikes) {
        return solveAssignBikes(workers, bikes);
    }
};

// Driver Code for testing.
int main() {

    return 0;
}
