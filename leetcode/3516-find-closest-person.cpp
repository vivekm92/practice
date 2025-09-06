#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/find-closest-person/description/
class Solution {
private:
    // Approach 1
    // T(n) : O(1) ; S(n) : O(1)
    int solveFindClosest(int x, int y, int z) {
        int d1 = abs(x-z), d2 = abs(y-z);
        if (d1 == d2) return 0;
        return d1 < d2 ? 1 : 2;
    }

    // Appraoch 2
    // T(n) : O(1) ; S(n) : O(1)
    int solveFindClosest1(int x, int y, int z) {
        int d = (x - y) * ( x + y - 2*z);
        // d < 0 then x is closer, d > 0 then y is closer
        return (2 * (d > 0 )) | (d < 0);
    }
public:
    int findClosest(int x, int y, int z) {
        return solveFindClosest1(x, y, z);
    }
};

// Driver Code for testing.
int main() {

    return 0;
}
