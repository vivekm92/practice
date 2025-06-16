#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/longest-unequal-adjacent-groups-subsequence-i/description/
class Solution {
private:
    // Appraoch 1
    // using DP
    // T(n) : O(n2) ; S(n) ; O(1)
    vector<string> solveGetLongestSubsequence(vector<string> &words, vector<int> &groups) {
        
        int n = words.size();
        
        return words;
    }
    
    // Approach 2
    // using greedy
    // T(n) ; O(n) ; S(n) : O(n)
    vector<string> solveGetLongestSubsequence1(vector<string> &words, vector<int> &groups) {
        
        int n = words.size();
        vector<string> res;
        if (n == 0) return res;

        res.push_back(words[0]);
        for (int i = 1; i < n; i++) {
            if (groups[i] != groups[i-1]) {
                res.push_back(words[i]);
            }
        }
        
        return res;
    }
    
    // Approach 3
    // using greedy
    // T(n) : O(n) ; S(n) : O(1)
    vector<string> solveGetLongestSubsequence2(vector<string> &words, vector<int> &groups) {
        
        int n = words.size(), idx = 0;
        for (int i = 1; i < n; i++) {
            if (groups[i] != groups[i-1]) {
                swap(words[idx+1], words[i]);
                idx++;
            }
        }
        
        auto first = words.begin();
        auto last = words.begin() + idx + 1;
        vector<string> res(first, last);
        return res;
    }
public:
    vector<string> getLongestSubsequence(vector<string>& words, vector<int>& groups) {
        return solveGetLongestSubsequence1(words, groups);
    }
};

// Driver Code for testing
int main() {
    
    
    return 0;
}
