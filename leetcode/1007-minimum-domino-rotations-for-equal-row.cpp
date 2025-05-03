#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/minimum-domino-rotations-for-equal-row/description/
class Solution {
private:
    // Approach 1
    // using simulation
    // T(n) : O(n) ; S(n) : O(1)
    int solveMinDominoRotations(vector<int>& tops, vector<int> &bottoms) {

        int n = tops.size(), v = -1;
        for (int i = 1; i < 7; i++) {
            bool flag = false;
            for (int j = 0; j < n; j++) {
                if (tops[j] != i && bottoms[j] != i) break;
                if (j == n-1) {
                    flag = true;
                    break;
                }
            }
            if (flag) {
                v = i;
                break;
            }
        }
        if (v == -1) return v;

        int cnt1 = 0, cnt2 = 0;
        for (int i = 0; i < n; i++) {
            if (tops[i] != v) cnt1++;
            if (bottoms[i] != v) cnt2++;
        }

        return min(cnt1, cnt2);
    }

    // Appraoch 2
    // using greedy
    // T(n) : O(n) ; S(n) : O(1)
    int solveMinDominoRotations1(vector<int> &tops, vector<int> &bottoms) {

        int n = tops.size();
        if (n == 0) return -1;

        int a = tops[0], b = bottoms[0];
        int cnt1 = 0, cnt2 = 0, cnt3 = 0, cnt4 = 0;
        for (int i = 0; i < n; i++) {
            if (a != tops[i] && a != bottoms[i]) a = 0;
            if (b != tops[i] && b != bottoms[i]) b = 0;
            if (a == 0 && b == 0) return -1;
            if (a == tops[i]) cnt1++;
            if (a == bottoms[i]) cnt2++;
            if (b == tops[i]) cnt3++;
            if (b == bottoms[i]) cnt4++;
        }

        return min(n - max(cnt1, cnt2), n - max(cnt3, cnt4));
    }
public:
    int minDominoRotations(vector<int>& tops, vector<int>& bottoms) {
        return solveMinDominoRotations(tops, bottoms);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
