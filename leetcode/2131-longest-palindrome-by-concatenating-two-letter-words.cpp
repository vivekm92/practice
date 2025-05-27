#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/longest-palindrome-by-concatenating-two-letter-words/description/
class Solution {
private:
    // Approach 1
    // using greedy approach with hash-map
    // T(n) : O(n) ; S(n) : O(n)
    int solveLongestPalindrome(vector<string> &words) {

        unordered_map<string, int> um;
        for (auto &word : words) um[word]++;

        bool oddLengthPresent = false;
        int maxLength = 0;
        for (auto it = um.begin(); it != um.end(); it++) {
            auto s = it->first;
            if (s[0] == s[1]) {
                if (it->second % 2 != 0) oddLengthPresent = true;
                maxLength += (it->second / 2) * 2 * 2;
            } else {
                string t = s;
                reverse(t.begin(), t.end());
                if (um.count(t) > 0) {
                    maxLength += min(it->second, um[t]) * 2;
                }
            }
        }

        if (oddLengthPresent) maxLength += 2;

        return maxLength;
    }

    // Approach 2
    // using greedy approach with 2-D matrux
    // T(n) : O(n) ; S(n) : O(n)
    int solveLongestPalindrome1(vector<string> &words) {

        vector<vector<int> > matrix(26, vector<int>(26, 0));
        for (auto &word : words) {
            char ch1 = word[0], ch2 = word[1];
            matrix[ch1 - 'a'][ch2 - 'a']++;
        }

        int count = 0;
        bool isOddPresent = false;
        for (int i = 0; i < 26; i++) {
            for (int j = i; j < 26; j++) {
                int tempCount = matrix[i][j];
                if (tempCount == 0) continue;
                else if (i == j) {
                    if (tempCount % 2 != 0) isOddPresent = true;
                    count += (tempCount / 2) * 2;
                } else {
                    count += min(tempCount, matrix[j][i]) * 2;
                }
            }
        }

        if (isOddPresent) count++;

        return 2*count;
    }
public:
    int longestPalindrome(vector<string>& words) {
        return solveLongestPalindrome1(words);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
