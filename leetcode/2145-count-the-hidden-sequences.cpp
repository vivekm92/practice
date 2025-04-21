#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/count-the-hidden-sequences/description/
class Solution {
private:
    // Approach 1
    // T(n) : O(n) ; S(n) : O(1)
    int solveNumberOfArrays(vector<int> &differences, int lower, int upper) {

        int hMin = 0, hMax = 0, curr = 0;
        for (int &diff : differences) {
            curr += diff;
            hMin = min(hMin, curr);
            hMax = max(hMax, curr);
            if (hMax - hMin > upper - lower) return 0;
        }

        return (upper - lower) - (hMax - hMin) + 1;
    }
public:
    int numberOfArrays(vector<int>& differences, int lower, int upper) {
        return solveNumberOfArrays(differences, lower, upper);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
