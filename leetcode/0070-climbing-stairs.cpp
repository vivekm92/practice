#include <iostream>
#include <vector>
#include <functional>

using namespace std;

// https://leetcode.com/problems/climbing-stairs/description/
class Solution {
private:
    // Approach 1
    // brute-force ( using recursion )
    // T(n) ; O(2^n) ; S(n) : O(n)
    int solveClimbStairs(int n) {
        if (n <= 2) return n;
        return solveClimbStairs(n-1) + solveClimbStairs(n-2);
    }

    // Approach 2
    // bottom-up dp
    // T(n) : O(n) ; S(n) ; O(n)
    int solveClimbStairs1(int n) {

        if (n <= 2) return n;
        vector<int> dp(n+1, 0);
        dp[1] = 1;
        dp[2] = 2;
        for (int i = 3; i <= n; i++) {
            dp[i] = dp[i-1] + dp[i-2];
        }

        return dp[n];
    }

    // Approach 3
    // top-down dp
    // T(n) ; O(n) ; S(n) : O(n)
    int solveClimbStairs2(int n) {

        if (n <= 2) return n;

        vector<int> dp(n+1, -1);
        dp[1] = 0;
        dp[1] = 1;
        dp[2] = 2;

        function<int(int)> dfs = [&](int i) {
          if (i <= 2) return i;
          if (dp[i] != -1) return dp[i];
          return dp[i] = dfs(i-1) + dfs(i-2);
        };

        return dfs(n);
    }

    // Approach 4
    // iterative method
    // T(n) : O(n) ; S(n) ; O(1)
    int solveClimbStairs3(int n) {
        if (n <= 2) return n;
        int a = 1, b = 2;
        for (int i = 3; i <= n; i++) {
            int c = a + b;
            a = b;
            b = c;
        }

        return b;
    }

    // Approach 4
    // matrix exponentiation (binet's method)
    // T(n) : O(logn) ; S(n) : O(1)
    int solveClimbStairs4(int n) {

        if (n <= 2) return n;
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
        matrixPower(A, n);
        return A[0][0];
    }

    // Approach 5
    // using math ( golden's ratio )
    // T(n) : O(logn) ; S(n) : O(1)
    int solveClimbStairs5(int n) {
        double goldenRatio = (1 + sqrt(5)) / 2;
        return round(pow(goldenRatio, n+1) / sqrt(5));
    }
public:
    int climbStairs(int n) {
        return solveClimbStairs(n);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
