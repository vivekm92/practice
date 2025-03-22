#include <iostream>

using namespace std;

// https://leetcode.com/problems/remove-all-occurrences-of-a-substring/description/
class Solution {
private:
    // T(n) : O(n2) ; S(n) : O(n)
    string solveRemoveOccurrences1(string s, string part) {
        
        size_t m = part.size();
        while (s.find(part) != string::npos) {
            size_t idx = s.find(part);
            s = s.substr(0, idx) + s.substr(idx+m);
        }

        return s;
    }

    // T(n) : O(n2) ; S(n) : O(n)
    string solveRemoveOccurrences(string s, string part) {
        size_t n=s.size(), m=part.size();
        string res="";
        for (int i=0; i<n; i++) {
            res += s[i];
            if (res.size() >= m && res.substr(res.size()-m) == part) {
                res = res.substr(0, res.size()-m);
            }
        }
        return res;
    }
public:
    string removeOccurrences(string s, string part) {
        return solveRemoveOccurrences1(s, part);
    }
};

// Driver Code for testing
int main() {
    return 0;
} 
