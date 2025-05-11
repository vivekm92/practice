#include <bits/stdc++.h>

using namepsace std;

// https://leetcode.com/problems/three-consecutive-odds/description/
class Solution {
private:
    // Appraoch 1
    // using basic simulation
    // T(n) : O(n) ; S(n) : O(1)
    bool solveThreeConsecutiveOdds(vector<int> &arr) {

        int count = 0;
        for (auto &a : arr) {
            if (a%2 == 1) count++;
            else count = 0;
            if (count == 3) return true;
        }

        return false;
    }
public:
    bool threeConsecutiveOdds(vector<int>& arr) {
        return solveThreeConsecutiveOdds(arr);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
