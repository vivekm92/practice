#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/maximum-number-of-words-you-can-type/description/
class Solution {
private:
    // Approach 1
    // using lookup
    // T(n) : O(n) ; S(n) : O(1)
    int solveCanBeTypedWords(string text, string brokenLetters) {

        vector<bool> isBroken(26, false);
        for (auto &ch : brokenLetters) {
            isBroken[ch-'a'] = true;
        }

        int totalWords = 0, brokenWordsCount = 0;
        int n = text.size();
        bool brokenWord = false;
        for (int i = 0; i < n; i++) {
            if (text[i] != ' ' && isBroken[text[i]-'a']) brokenWord = true;
            if (text[i] == ' ' || i == n-1) {
                totalWords++;
                if (brokenWord) {
                    brokenWordsCount++;
                }
                brokenWord = false;
                continue;
            }
        }

        return totalWords - brokenWordsCount;
    }
public:
    int canBeTypedWords(string text, string brokenLetters) {
        return solveCanBeTypedWords(text, brokenLetters);
    }
};

// Driver Code for testing.
int main() {

    return 0;
}
