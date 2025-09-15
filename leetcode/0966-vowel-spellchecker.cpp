#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/vowel-spellchecker/description/
class Solution {
private:

    string maskCase(string s) {
        for (auto &ch : s) {
            if (ch >= 'A' && ch <= 'Z') {
                ch += ('a' - 'A');
            }
        }
        return s;
    }

    bool isVowel(char ch) {
        return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
            ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U';
    }

    string maskVowel(string s) {
        for (auto &ch : s) {
            if (isVowel(ch)) {
                ch = '*';
            }
        }
        return s;
    }

    // Approach 1
    // using hash-map
    // T(n) : O(n) ; S(n) : O(n)
    vector<string> solveSpellChecker(vector<string> &wordlist, vector<string> &queries) {

        unordered_set<string> original;
        unordered_map<string, int> maskedCaseLookup, maskedVowelLookup;
        for (int i = 0; i < wordlist.size(); i++) {
            string s = wordlist[i];
            original.insert(s);

            string maskedCaseStr = maskCase(s);
            if (maskedCaseLookup.find(maskedCaseStr) == maskedCaseLookup.end()) {
                maskedCaseLookup[maskedCaseStr] = i;
            }

            string maskedVowelStr = maskVowel(maskedCaseStr);
            if (maskedVowelLookup.find(maskedVowelStr) == maskedVowelLookup.end()) {
                maskedVowelLookup[maskedVowelStr] = i;
            }
        }

        vector<string> res;
        for (auto &query : queries) {
            if (original.find(query) != original.end()) {
                res.push_back(query);
                continue;
            }

            string maskedCaseStr = maskCase(query);
            if (maskedCaseLookup.find(maskedCaseStr) != maskedCaseLookup.end()) {
                res.push_back(wordlist[maskedCaseLookup[maskedCaseStr]]);
                continue;
            }

            string maskedVowelStr = maskVowel(maskedCaseStr);
            if (maskedVowelLookup.find(maskedVowelStr) != maskedVowelLookup.end()) {
                res.push_back(wordlist[maskedVowelLookup[maskedVowelStr]]);
                continue;
            }

            res.push_back("");
        }
        return res;
    }
public:
    vector<string> spellchecker(vector<string>& wordlist, vector<string>& queries) {
        return solveSpellChecker(wordlist, queries);
    }
};

// Driver Code for testing.
int main() {

    return 0;
}
