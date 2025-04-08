#include <iostream>
#include <functional>
#include <unordered_map>

using namespace std;

// https://leetcode.com/problems/fibonacci-number/description/
class Solution {
private:
    // Approach 1
    // using recursion
    // T(n) : O(2^n) ; S(n) : O(n)
    int solveFib(int n) {
        if (n <= 1) return n;
        return solveFib(n-1) + solveFib(n-2);
    }

    int solveFibHelper1(int n, unordered_map<int, int> &um) {
        if (n <= 1) return um[n];
        if (um.find(n) != um.end()) return um[n];
        return um[n] = solveFibHelper1(n-1, um) + solveFibHelper1(n-2, um);
    }

    // Approach 2
    // using top-down DP
    // T(n) : O(n) ; S(n) : O(n)
    int solveFib1(int n) {
        unordered_map<int, int> um;
        um[0] = 0;
        um[1] = 1;
        return solveFibHelper1(n, um);
    }

    // Approach 3
    // iterative approach
    // T(n) : O(n) ; S(n) ; O(1)
    int solveFib2(int n) {

        int a = 0, b = 1;
        for (int i = 0; i < n-1; i++) {
            int c = a + b;
            a = b;
            b = c;
        }

        return b;
    }

    // Approach 4
    // using bottom-up dp
    // T(n) : O(n) ; S(n) : O(n)
    int solveFib3(int n) {

        if (n <= 1) return n;
        vector<int> dp(n + 1);
        dp[0] = 0;
        dp[1] = 1;
        for (int i = 2; i <= n; i++) {
            dp[i] = dp[i-1] + dp[i-2];
        }

        return dp[n];
    }

    // Approach 5
    // using matrix exponentiation
    // T(n) : O(logn) ; S(n) : O(logn)
    int solveFib4(int n) {

        if (n <= 1) return n;

        function<vector<vector<int> >(vector<vector<int> >&, vector<vector<int> >&)>  mul = [&](vector<vector<int> > &A, vector<vector<int> > &B) {
            int a = A[0][0] * B[0][0] + A[0][1] * B[1][0];
            int b = A[0][0] * B[1][0] + A[0][1] * B[1][1];
            int c = A[1][0] * B[0][0] + A[1][1] * B[1][0];
            int d = A[1][0] * B[1][0] + A[1][1] * B[1][1];
            A[0][0] = a;
            A[0][1] = b;
            A[1][0] = c;
            A[1][1] = d;

            return A;
        };

        function<void(vector<vector<int> >&, int)> matrixPower = [&](vector<vector<int> > &A, int n) {

            if (n <= 1) return ;
            matrixPower(A, n/2);
            mul(A, A);

            vector<vector<int> > B = {{1,1}, {1,0}};
            if (n % 2 != 0) {
                mul(A, B);
            }

            return ;
        };

        vector<vector<int> > A {{1,1}, {1,0}};
        matrixPower(A, n - 1);
        return A[0][0];
    }

    // Approach 5
    // using math (golden ratio)
    // T(n) : O(logn) ; S(n) ; O(1)
    int solveFib5(int n) {
        double goldenRatio = (1 + sqrt(5)) / 2;
        return round(pow(goldenRatio, n) / sqrt(5));
    }
public:
    int fib(int n) {
        return solveFib3(n);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
