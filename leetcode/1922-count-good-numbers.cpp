#include <iostream>
#include <functional>

using namespace std;

// https://leetcode.com/problems/count-good-numbers/description/
class Solution {
private:
    long long modPow(long long x, long long n, const int MOD) {

        if (n == 0) return 1;
        else if (n == 1) return x;

        long long v1 = modPow(x, n/2, MOD);
        if (n % 2 == 0) return ((v1 % MOD) * (v1 % MOD)) % MOD;

        return ((((v1 % MOD) * (v1 % MOD)) % MOD) * (x % MOD)) % MOD;
    }

    // Approach 1
    // using mod exponentiation ( recursion )
    // T(n) : O(logn) ; S(n) : O(logn)
    int solveCountGoodNumbers(long long n) {

        const int MOD = 1000000007;
        long long evenPlaces = (n+1)/2 , oddPlaces = n/2;
        long long evenPlacesCount = 5, oddPlacesCount = 4;

        long long v1 = modPow(evenPlacesCount, evenPlaces, MOD);
        long long v2 = modPow(oddPlacesCount, oddPlaces, MOD);

        return ((v1 % MOD) * (v2 % MOD)) % MOD;
    }

    // Approach 2
    // using mod exponentiation ( iterative )
    // T(n) : O(logn) ; S(n) : O(1)
    int solveCountGoodNumbers1(long long n) {

        const int MOD = 1000000007;
        long long evenPlaces = (n+1)/2, oddPlaces = n/2;
        long long evenPlacesCount = 5, oddPlacesCount = 4;

        function<long long(long long, long long)> modPow1 = [&](long long x, long long y) {

            long long res = 1;
            while (y > 0) {

                if (y % 2 != 0) {
                    res = ((res % MOD) * (x % MOD)) % MOD;
                    y -= 1;
                }

                x = ((x % MOD )* (x % MOD)) % MOD;
                y /= 2;
            }

            return res;
        };

        long long v1 = modPow1(evenPlacesCount, evenPlaces);
        long long v2 = modPow1(oddPlacesCount, oddPlaces);

        return (v1 * v2) % MOD;
    }
public:
    int countGoodNumbers(long long n) {
        return solveCountGoodNumbers(n);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
