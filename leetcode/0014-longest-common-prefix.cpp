#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/longest-common-prefix/description/
class Solution {
private:
    // Approach 1
    // using two-pointers
    // T(n) : O(m*n) ; S(n) : O(1)
    string solveLongestCommonPrefix(vector<string> &strs) {

        string res = "";
        int n = strs.size();
        if (n == 0) return res;
        else if (n == 1) return strs[0];
        for (int i = 0; i < strs[0].size(); i++) {
            char ch = strs[0][i];
            bool flag = false;
            for (int j = 1; j < n; j++) {
                if (i >= strs[j].size()) break;
                if (ch != strs[j][i]) break;
                if (j == n-1) flag = true;
            }
            if (flag) res += ch;
            else break;
        }

        return res;
    }

    // Approach 2
    // using horizontal scanning
    // T(n) : O(mn) ; S(n) : O(1)
    string solveLongestCommonPrefix1(vector<string> &strs) {

        int n = strs.size();
        if (n == 0) return "";
        string prefix = strs[0];
        for (int i = 1; i < n; i++) {
            while (strs[i].find(prefix) != 0) {
                prefix = prefix.substr(0, prefix.size()-1);
                if (prefix.empty()) return "";
            }
        }

        return prefix;
    }

    string commonPrefix(string s, string t) {

        int sn = s.size(), tn = t.size();
        if (sn == 0 || tn == 0) return "";

        int si = 0;
        while (si < sn && si < tn && s[si] == t[si]) {
            si++;
        }

        return s.substr(0, si);
    }

    string solveLongestCommonPrefixHelper2(vector<string> &strs, int start, int end) {
        if (start > end) return "";
        else if (start == end) return (string )strs[start];

        int mid = start + (end - start) / 2;
        string left = solveLongestCommonPrefixHelper2(strs, start, mid);
        string right = solveLongestCommonPrefixHelper2(strs, mid + 1, end);

        return commonPrefix(left, right);
    }

    // Approach 3
    // using DnC
    // T(n) : O(mlogn) ; S(mlogn)
    string solveLongestCommonPrefix2(vector<string> &strs) {

        int n = strs.size();
        if (n == 0) return "";
        return solveLongestCommonPrefixHelper2(strs, 0, n-1);
    }

    bool isCommonPrefix(vector<string> &strs, int mid) {

        int n = strs.size();
        string s = strs[0].substr(0, mid);
        for (int i = 1; i < n; i++) {
            if (s != strs[i].substr(0, mid)) return false;
        }

        return true;
    }

    // Approach 4
    // using binary-search
    // T(n) ; O(nlogm) ; S(n) : O(1)
    string solveLongestCommonPrefix3(vector<string> &strs) {

        int n = strs.size();
        if (n == 0) return "";

        int minLen = INT_MAX;
        for (auto &s : strs) minLen = min(minLen, (int )s.size());

        int l = 1, r = minLen;
        while (l <= r) {
            int mid = l + (r - l) / 2;
            if (isCommonPrefix(strs, mid)) {
                l = mid + 1;
            } else {
                r = mid - 1;
            }
        }

        return strs[0].substr(0, l-1);
    }
public:
    string longestCommonPrefix(vector<string>& strs) {
        return solveLongestCommonPrefix2(strs);
    }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    vector<string> srr {"a", "aba", "abc"};
    cout << sol->longestCommonPrefix(srr) << "\n";
    return 0;
}
