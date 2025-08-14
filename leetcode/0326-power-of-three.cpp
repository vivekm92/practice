#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/power-of-three/description/
class Solution {
private:
    // Approach 1
    // T(n) : O(logn) ; S(n) : O(1)
    bool solveIsPowerOfThree(int n) {

        if (n < 1) return false;
        while (n % 3 == 0) {
            n /= 3;
        }

        return n == 1;
    }

    // Approach 2
    // T(n) : O(logn) ; S(n) : O(1)
    bool solveIsPowerOfThree1(int n) {

        if (n < 1) return false;
        float k = log(n) / log(3);
        return n == pow(3, k);
    }
public:
    bool isPowerOfThree(int n) {
        return solveIsPowerOfThree(n);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
