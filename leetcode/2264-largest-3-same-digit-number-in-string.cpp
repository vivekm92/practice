#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/largest-3-same-digit-number-in-string/description/
class Solution {
private:
    // Approach 1
    // T(n) : O(n) ; S(n) : O(1)
    string solveLargestGoodInteger(string num) {

        string res = "";
        int n = num.size(), maxVal = -1;
        for (int i = 0; i < n - 2; i++) {
            if (num[i] == num[i+1] && num[i+1] == num[i+2]) {
                int val = num[i]*100 + num[i]*10 + num[i];
                if (maxVal < val) res = num.substr(i, 3);
                maxVal = max(maxVal, val);
            }
        }

        return res;
    }

    // Approach 2
    // T(n) : O(n) ; S(n) : O(1)
    string solveLargestGoodInteger1(string num) {

        char ch = '0' - 1;
        int n = num.size();
        for (int i = 0; i < n - 2; i++) {
            char c = num[i];
            if (c == num[i+1] && c == num[i+2]) {
                ch = max(ch, c);
            }
        }

        return ch == '0' - 1 ? "" : string(3, ch);
    }
public:
    string largestGoodInteger(string num) {
        return solveLargestGoodInteger1(num);
    }
};

// Driver Code for testing.
int main() {

    return 0;
}
