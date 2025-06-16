#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

// https://leetcode.com/problems/is-subsequence/description/
class Solution {
private:
    // Approach 1
    // using basic simulation
    // T(n) : O(n)  ; S(n) : O(1)
    bool solveIsSubsequence(string s, string t) {

        int idx = 0, ns = s.size(), nt = t.size();
        for (int i = 0; i < nt; i++) {
            if (idx == ns) return true;
            if (idx < ns && s[idx] == t[i]) {
                idx++;
            }
        }

        return idx == ns;
    }

    bool solveIsSubsequenceHelper1(string s, int si, int sn, string t, int ti, int tn) {
        if (si == sn) return true;
        else if (ti == tn) return false;

        if (s[si] == t[ti]) {
            return solveIsSubsequenceHelper1(s, si+1, sn, t, ti+1, tn);
        }

        return solveIsSubsequenceHelper1(s, si, sn, t, ti+1, tn);
    }

    // Approach 2
    // using recursion ( DnC )
    // T(n) : O(n) ; S(n) : O(n)
    bool solveIsSubsequence1(string s, string t) {
        return solveIsSubsequenceHelper1(s, 0, s.size(), t, 0, t.size());
    }
    
    bool solveIsSubsequenceHelper2(string s, int si, int sn, string t, int ti, int tn) {
        
        if (si == sn) return true;
        else if (ti == tn) return false;
        
        if (s[si] == t[ti]) {
            si++;
        }
        ti++;
        return solveIsSubsequenceHelper2(s, si, sn, t, ti, tn);
    }
    
    // Approach 3
    // using recursion ( DnC ) -- tail recursion
    // T(n) : O(n) ; S(n) : O(n)
    bool solveIsSubsequence2(string s, string t) {
        return solveIsSubsequenceHelper2(s, 0, s.size(), t, 0, t.size());
    }

    // Approach 4
    // using hash-map
    // T(n) : O(n) ; S(n) : O(n)
    bool solveIsSubsequence3(string s, string t) {
        
        int tn = t.size();
        unordered_map<char, vector<int> > um;
        for (int i = 0; i < tn; i++) {
            if (um.find(t[i]) == um.end()) {
                um[t[i]] = {i};
            } else {
                um[t[i]].push_back(i);
            }
        }
        
        int sn = s.size();
        int currIdx = -1;
        for (int i = 0; i < sn; i++) {
            if (um.find(s[i]) == um.end()) return false;
            bool flag = false;
            for (auto &idx : um[s[i]]) {
                if (currIdx < idx) {
                    flag = true;
                    currIdx = idx;
                    break;
                }
            }
            if (!flag) return false;
        }
        
        return true;
    }
    
    // Approach 5
    // using hash-map
    // T(n) : 
    
    // Approach 6
    // using DP
    // T(n) ; O(n) ; S(n) : O(n2)
    bool solveIsSubsequence5(string s, string t) {
        
        int sn = s.size(), tn = t.size();
        vector<vector<int> > dp(sn+1, vector<int>(tn+1, 0));
        for (int i = 1; i <= sn; i++) {
            for (int j = 1; j <= tn; j++) {
                if (s[i-1] == t[j-1]) {
                    dp[i][j] = dp[i-1][j-1] + 1;
                } else {
                    dp[i][j] = max(dp[i-1][j], dp[i][j-1]);
                }
            }
        }
        
        return dp[sn][tn] == sn;
    }
public:
    bool isSubsequence(string s, string t) {
        return solveIsSubsequence5(s, t);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
