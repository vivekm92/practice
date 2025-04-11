#include <iostream>

using namespace std;

// https://leetcode.com/problems/powx-n/description/
class Solution {
private:
    // Approach 1
    // recursive 
    // T(n) : O(logn) ; S(n) : O(logn)
    double solveMyPow(double x, int n) {
        if (n == 0) return 1;
        else if (n == 1) return x;
        
        double v1 = solveMyPow(x, n/2);
        if (n%2 == 0) return v1 * v1;
        
        return v1 * v1 * x;
    }
    
    
    // Approach 2
    // iterative
    // T(n) : O(logn) ; S(n) : O(1)
    double solveMyPow1(double x, long n) {
        if (n == 0) return 1;
        else if (n == 1) return x;
    
        double res = 1;
        while (n > 0) {
            if (n % 2 == 1) {
                res = res * x;
                n -= 1;
            }
            
            x = x * x;
            n /= 2;
        }
        
        return res;
    }
public:
    double myPow(double x, int n) {
        if (n < 0) return 1 / solveMyPow1(x, abs((long )n));
        return solveMyPow1(x, abs(n));
    }
};

// Driver Code for testing
int main() {
    
    return 0;
}
