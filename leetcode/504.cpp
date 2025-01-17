#include <iostream>

using namespace std;


// https://leetcode.com/problems/base-7/description/
class Solution {
private:
    // T(n) : O(logn) : S(n) : O(1)
    string solveconvertToBase7(int num) {
        if (num == 0) return "0";
        string res = "";
        int n = abs(num);
        while (n > 0) {
            res = to_string(n%7) + res;
            n /= 7;
        }

        return num < 0 ? "-"+res : res;
    }
    
    // T(n) : O(logn) : S(n) : O(1)
    string _solveconvertToBase7(int num) {
        if (num == 0) return "0";
        string res = "";
        int n = abs(num);
        while (n > 0) {
            res += to_string(n%7);
            n /= 7;
        }
        reverse(res.begin(), res.end());
        return num < 0 ? "-"+res : res;
    }
public:
    string convertToBase7(int num) {
        return _solveconvertToBase7(num);
    }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    cout << sol->convertToBase7(7) << "\n";
    cout << sol->convertToBase7(-7) << "\n";
    return 0;
}
