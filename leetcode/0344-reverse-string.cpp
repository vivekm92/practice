#include <iostream>

using namespace std;

// https://leetcode.com/problems/reverse-string/description/
class Solution {
private:
    // Approach 1
    // using recursion
    // T(n) : O(n) ; S(n) : O(n)
    void solveReverseString(vector<char> &s, int idx) {

        if (idx >= s.size() / 2) {
            return ;
        }

        swap(s[idx], s[s.size() - 1 - idx]);
        solveReverseString(s, idx + 1);
    }

    // Approach 2
    // using iteration
    // T(n) : O(n) ; S(n) : O(1)
    void solveReverseString1(vector<char> &s) {

        int n = s.size();
        for (int i = 0; i < n / 2; i++) {
            swap(s[i], s[n-1-i]);
        }

        return ;
    }
public:
    void reverseString(vector<char>& s) {
        // solveReverseString(s, 0);
        solveReverseString1(s);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
