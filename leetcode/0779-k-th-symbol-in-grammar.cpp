#include <iostream>

using namespace std;

// https://leetcode.com/problems/k-th-symbol-in-grammar/description/
class Solution {
private:
    // Approach 1
    // using recursion
    // T(n) : O(n) ; S(n) : O(n)
    int solveKthGrammar(int n, int k) {
        if (k == 2) {
            return 1;
        } else if (n == 1 || k == 1) {
            return 0;
        }

        int x = pow(2, n-1);
        if (k > x / 2) {
            if (n%2 != 0) return solveKthGrammar(n-1, x-k+1);
            else return solveKthGrammar(n-1, x-k+1) ^ 1;
        }

        return solveKthGrammar(n-1, k);
    }

    // Approach 2
    // using math
    // T(n) : O(logn) ; S(n) : O(1)
    int solveKthGrammar1(int n, int k) {
        int count = __builtin_popcount(k-1);
        return count % 2 == 0 ? 0 : 1;
    }
public:
    int kthGrammar(int n, int k) {
        return solveKthGrammar1(n, k);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
