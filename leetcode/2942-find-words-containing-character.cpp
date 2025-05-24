#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/find-words-containing-character/description/
class Solution {
private:
    // Approach 1
    // using basic simulation
    // T(n) : O(n*m) ; S(n) : O(n)
    vector<int> solveFndWordsContaining(vector<string> &words, char x) {

        vector<int> res;
        for (int i = 0; i < words.size(); i++) {
            for (auto &ch : words[i]) {
                if (ch == x) {
                    res.push_back(i);
                    break;
                }
            }
        }

        return res;
    }
public:
    vector<int> findWordsContaining(vector<string>& words, char x) {
        return solveFndWordsContaining(words, x);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
