#include <iostream>

using namespace std;

// https://leetcode.com/problems/convert-a-number-to-hexadecimal/description/
class Solution {
private:
    // T(n) : O(logn) ; S(n) : O(1)
    string solvetoHex(int num) {
        vector<char> lookup {'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'};
        if (num == 0) return "0";
        string res = "";
        for (int i = 0; i < 8 && num != 0; i++) {
            int t = num & 0xf;
            res = lookup[t] + res;
            num = num >> 4;
        }
        return res;
    }
public:
    string toHex(int num) {
        return solvetoHex(num);
    }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    cout << sol->toHex(-1) << "\n";
    cout << sol->toHex(26) << "\n";
    return 0;
}
