#include <iostream>
#include <unordered_map>

using namespace std;

// https://leetcode.com/problems/count-of-substrings-containing-every-vowel-and-k-consonants-ii/description/
class Solution {
private:
    bool isVowel(char ch) {
        return (ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u');
    }

    // using hash-map with sliding window
    // T(n) : O(n) ; S(n) : O(n)
    long long solveCountOfSubstrings(string word, int k) {

        int n = word.size();
        int prev = n;
        vector<int> arr(n, 0);
        for (int i = n-1; i >= 0; i--) {
            arr[i] = prev;
            if (!isVowel(word[i])) {
                prev = i;
            }
        }

        long long validSubstringCount = 0;
        int consonantCount = 0, start = 0, end = 0;
        unordered_map<char, int> um;
        while (end < n) {
            if (isVowel(word[end])) {
                um[word[end]]++;
            } else {
                consonantCount++;
            }

            while (consonantCount > k) {
                if (isVowel(word[start])) {
                    um[word[start]]--;
                    if (um[word[start]] == 0) {
                        um.erase(word[start]);
                    }
                } else {
                    consonantCount--;
                }
                start++;
            }

            while (um.size() == 5 && consonantCount == k) {
                validSubstringCount += (arr[end] - end);
                if (isVowel(word[start])) {
                    um[word[start]]--;
                    if (um[word[start]] == 0) {
                        um.erase(word[start]);
                    }
                } else {
                    consonantCount--;
                }
                start++;
            }
            end++;
        }

        return validSubstringCount;
    }

    long long solveCountOfSubstringsHelper(string word, int k) {

        int n = word.size();
        int start = 0, end = 0, consonantCount = 0;
        long long validSubstringCount = 0;
        unordered_map<char, int> um;
        while (end < n) {
            if (isVowel(word[end])) {
                um[word[end]]++;
            } else {
                consonantCount++;
            }

            while (um.size() == 5 && consonantCount >= k) {
                validSubstringCount += (n - end);
                if (isVowel(word[start])) {
                    um[word[start]]--;
                    if (um[word[start]] == 0) {
                        um.erase(word[start]);
                    }
                } else {
                    consonantCount--;
                }
                start++;
            }
            end++;
        }
        return validSubstringCount;
    }

    // T(n) : O(n) ; S(n) : O(1)
    long long solveCountOfSubstrings1(string word, int k) {
        return solveCountOfSubstringsHelper(word, k) - solveCountOfSubstringsHelper(word, k + 1);
    }
public:
    long long countOfSubstrings(string word, int k) {
        return solveCountOfSubstrings1(word, k);
    }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    cout << sol->countOfSubstrings("ieaouqqieaouqq", 1) << "\n";

    return 0;
}
