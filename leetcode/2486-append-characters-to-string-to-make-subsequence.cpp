#include <iostream>

using namespace std;

// https://leetcode.com/problems/append-characters-to-string-to-make-subsequence/description/
class Solution {
private:
    // Approach 1
    // using two-pointers
    // T(n) : O(n) ; S(n) : O(1)
    int solveAppendCharacters(string s, string t) {

        int ti = 0, tn = t.size(), sn = s.size();
        for (int si = 0; si < sn; si++) {
            if (s[si] == t[ti]) ti++;
        }

        return tn - ti;
    }
public:
    int appendCharacters(string s, string t) {
        return solveAppendCharacters(s, t);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
