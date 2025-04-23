#include <iostream>

using namespace std;

// https://leetcode.com/problems/valid-anagram/description/
class Solution {
private:
    // Approach 1
    // using hash-map
    // T(n) : O(n) ; S(n) : O(1)
    bool solveIsAnagram(string s, string t) {

        vector<int> arr(26, 0);
        for (auto &c : s) arr[c-'a']++;
        for (auto &c : t) arr[c-'a']--;
        for (int i = 0; i < 26; i++) {
            if (arr[i] != 0) return false;
        }
        return true;
    }

    // Approach 2
    // using sorting
    // T(n) : O(nlogn) ; S(n) : O(1)
    bool solveIsAnagram1(string s, string t) {

        if (s.size() != t.size()) return false;
        sort(s.begin(), s.end());
        sort(t.begin(), t.end());
        for (int i = 0; i < s.size(); i++) {
            if (s[i] != t[i]) return false;
        }

        return true;
    }
public:
    bool isAnagram(string s, string t) {
        return solveIsAnagram1(s, t);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
