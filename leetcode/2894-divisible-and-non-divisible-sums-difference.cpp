#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/divisible-and-non-divisible-sums-difference/description/
class Solution {
private:
    // Approach 1
    // using math
    // T(n) : O(1) ; S(n) : O(1)
    int solveDifferenceOfSums(int n, int m) {

        int q = n / m ;
        return (n * (n + 1) / 2)  - (2 * m * (q * (q + 1) / 2 ) );
    }
public:
    int differenceOfSums(int n, int m) {
        return solveDifferenceOfSums(n, m);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
