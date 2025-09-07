#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/description/
class Solution {
private:
    // Appraoch 1
    // T(n) : O(n) ; S(n) : O(1)
    vector<int> solveSumZero(int n) {

        int k = -n/2;
        vector<int> res;
        for (int i = 0; i < n; i++) {
            if (n%2 == 0 && k == 0) {
                k++;
            }
            res.push_back(k++);
        }
        a
        return res;
    }
public:
    vector<int> sumZero(int n) {
        return solveSumZero(n);
    }
};

// Driver Code for testing.
int main() {

    return 0;
}
