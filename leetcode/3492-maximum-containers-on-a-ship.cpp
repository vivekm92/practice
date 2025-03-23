#include <iostream>

using namespace std;

// https://leetcode.com/problems/maximum-containers-on-a-ship/description/
class Solution {
private:
    int solveMaxContainers(int n, int w, int maxWeights) {
        return min(n*n, maxWeights / w);
    }
public:
    int maxContainers(int n, int w, int maxWeight) {
        return solveMaxContainers(n, w, maxWeight);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
