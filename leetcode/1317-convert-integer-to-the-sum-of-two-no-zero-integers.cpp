#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/convert-integer-to-the-sum-of-two-no-zero-integers/description/
class Solution {
private:
    bool containsZero(int n) {

        while (n > 0) {
            int x = n % 10;
            if (x == 0) return true;
            n /= 10;
        }

        return false;
    }

    // Approach 1
    // T(n) : O(nlogn) ; S(n) : O(1)
    vector<int> solveGetNoZeroIntegers(int n) {

        for (int i = 1; i <= n / 2; i++) {
            int a = i;
            int b = n - i;
            if (!containsZero(a) && !containsZero(b)) return {a, b};
        }

        return {-1, -1};
    }

    // Approach 2
    // T(n) : O(nlogn) ; S(n) : O(logn)
    vector<int> solveGetNoZeroIntegers1(int n) {

        for (int i = 1; i <= n/2; i++) {
            int a = i;
            int b = n - i;
            string t = to_string(a) + to_string(b);
            if (t.find('0') == string::npos) {
                return {a, b};
            }
        }

        return {-1, -1};
    }
public:
    vector<int> getNoZeroIntegers(int n) {
        return solveGetNoZeroIntegers1(n);
    }
};

// Driver Code for testing.
int main() {

    return 0;
}
