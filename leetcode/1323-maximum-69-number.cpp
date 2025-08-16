#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/maximum-69-number/description/
class Solution {
private:
    // Approach 1
    // T(n) : O(logn) ; S(n) : O(1)
    int solveMaximum69Number(int num) {

        int val = num, pos = -1, cnt = 0;
        while (val > 0) {
            int n = val % 10;
            if (n == 6) {
                pos = cnt;
            }
            val /= 10;
            cnt++;
        }

        if (pos == -1) return num;
        return num - 6 * pow(10, pos) + 9 * pow(10, pos);
    }

    // Approach 2
    // T(n) : O(logn) ; S(n) : O(logn)
    int solveMaximum69Number1(int num) {

        string t = to_string(num);
        for (auto &n : t) {
            if (n == '6') {
                n = '9';
                break;
            }
        }

        return stoi(t);
    }
public:
    int maximum69Number (int num) {
        return solveMaximum69Number1(num);
    }
};

// Driver Code for testing.
int main() {

    return 0;
}
