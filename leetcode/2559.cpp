#include <iostream>
#include <unordered_set>

using namespace std;

// https://leetcode.com/problems/count-vowel-strings-in-ranges/description/
class Solution {
    private:
        bool isVowel(char c) {
            return (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u');
        }

        bool isValid(string s) {
            int n = s.size();
            return isVowel(s[0]) && isVowel(s[n-1]);
        }

        // using prefix arr
        // T(n) : O(n) ; S(n) : O(n)
        vector<int> solvevowelStrings(vector<string>& words, vector<vector<int>>& queries) {
            int n = words.size();
            vector<int> prefixArray(n + 1, 0);
            for (int i = 0; i < n; i++) {
                if (isValid(words[i])) {
                    prefixArray[i + 1] += 1;
                }
                prefixArray[i + 1] += prefixArray[i];
            }

            vector<int> res;
            for (auto &query: queries) {
                int start = query[0], end = query[1];
                res.push_back(prefixArray[end + 1] - prefixArray[start]);
            }
            return res;
        }

        // using prefix arr : alternative implementation
        // T(n) : O(n) ; S(n) : O(n)
        vector<int> solvevowelStrings1(vector<string>& words, vector<vector<int>>& queries) {
            int n = words.size();
            vector<int> prefixArray(n + 1, 0);
            unordered_set<char> vowels{'a', 'e', 'i', 'o', 'u'};
            for (int i = 0; i < n; i++) {
                auto word = words[i];
                if (vowels.count(word[0]) > 0 && vowels.count(word[word.size()-1]) > 0) {
                    prefixArray[i + 1] += 1;
                }
                prefixArray[i + 1] += prefixArray[i];
            }

            vector<int> res;
            for (auto &query: queries) {
                int start = query[0], end = query[1];
                res.push_back(prefixArray[end + 1] - prefixArray[start]);
            }
            return res;
        }
    public:
        vector<int> vowelStrings(vector<string>& words, vector<vector<int>>& queries) {
            return solvevowelStrings1(words, queries);
        }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    vector<string> words = {"aba", "obo", "ebe"};
    vector<vector<int> > queries = {{0, 1}, {0,2}};
    vector<int> res = sol->vowelStrings(words, queries);
    for (auto &r : res) {
        cout << r << " ";
    }
    cout << "\n";

    return 0;
}
