#include <iostream>

using namespace std;

// https://leetcode.com/problems/clear-digits/description/
class Solution {
private:
    // T(n) : O(n) ; S(n) : O(n)
    string solveclearDigits(string s) {
        size_t n = s.size();
        string res = "";
        for (int i=0; i<n; i++) {
            if (s[i] >= '0' && s[i] <= '9') {
                res.pop_back();
            } else {
                res += s[i];
            }
        }
        return res;
    }
public:
    string clearDigits(string s) {
        return solveclearDigits(s);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
