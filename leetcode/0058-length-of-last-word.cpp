#include <iostream>

using namespace std;

// https://leetcode.com/problems/length-of-last-word/description/
class Solution {
private:
    // Approach 1
    // using basic simulation
    // T(n) : O(n) ; S(n) : O(1)
    int solveLengthOfLastWord(string s) {

        int i = s.size() - 1, lastWordSize = 0;
        while(i >= 0) {
            while (i >= 0 && s[i] == ' ') {
                i--;
            }

            while (i >= 0 && s[i] != ' ') {
                lastWordSize++;
                i--;
            }
            break;
        }

        return lastWordSize;
    }

    // Approach 2
    // using basic simulation
    // T(n) : O(n) ; S(n) : O(1)
    int solveLengthOfLastWord1(string s) {

        int i = 0, n = s.size();
        int lastWordSize = 0;
        while (i < n) {
            if (s[i] !=  ' ') lastWordSize = 0;
            while (i < n && s[i] != ' ') {
                lastWordSize++;
                i++;
            }
            i++;
        }

        return lastWordSize;
    }
public:
    int lengthOfLastWord(string s) {
        return solveLengthOfLastWord1(s);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
