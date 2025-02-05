#include <iostream>

using namespace std;

// https://leetcode.com/problems/check-if-one-string-swap-can-make-strings-equal/description/
class Solution {
private:
    // using freq-map
    // T(n) : O(n) ; S(n) : O(1)
    bool solveareAlmostEqual1(string s1, string s2) {
        int n1=s1.size(), n2=s2.size(), nDiff=0;
        if (n1!=n2) return false;
        if (n1==0) return true;
        vector<int> s1Map(26, 0), s2Map(26, 0);
        for (int i=0; i<n1; i++) {
            if (s1[i] != s2[i]) {
                nDiff++;
            }
            s1Map[s1[i]-'a']++;
            s2Map[s2[i]-'a']++;
            if (nDiff>2) return false;
        }
        for (int i=0; i<26; i++) {
            if (s1Map[i] != s2Map[i]) return false;
        }
        return true;
    }

    // T(n) : O(n) ; S(n) : O(1)
    bool solveareAlmostEqual(string s1, string s2) {
        int n1=s1.size(), n2=s2.size();
        if (n1!=n2) return false;
        if (n1==0) return true;
        int idx1=-1, idx2=-1, charCount=0;
        for (int i=0; i<n1; i++) {
            if (s1[i] != s2[i]) {
                charCount++;
                if (idx1==-1) idx1 = i;
                else idx2 = i;
            }
            if (charCount > 2) return false;
        }

        if (charCount == 0) return true;
        if (idx1 == -1 || idx2 == -1) return false; // i.e. charCount = 1;
        return s1[idx1] == s2[idx2] && s1[idx2] == s2[idx1];
    }
public:
    bool areAlmostEqual(string s1, string s2) {
        return solveareAlmostEqual(s1, s2);
    }
};

int main() {

    return 0;
}
