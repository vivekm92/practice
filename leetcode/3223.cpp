#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/minimum-length-of-string-after-operations/description/
class Solution {
private:
    // using hashmap
    // T(n) : O(n) ; S(n) : O(n)
    int solveminimumLength(string s) {

        int n = s.size();
        unordered_map<char, int> um;
        for (int i = 0; i < n; i++) {
            if (um.find(s[i]) == um.end()) {
                um[s[i]] = 1;
            } else {
                um[s[i]] += 1;
            }
        }

        int count = 0;
        for (auto it = um.begin(); it != um.end(); it++) {
            if (it->second > 2) {
                if (it->second % 2 == 0) {
                    count += 2;
                } else {
                    count += 1;
                }
            } else {
                count += it->second;
            }
        }

        return count;
    }

    // using arr
    // T(n) : O(n) ; S(n) : O(1)
    int solveminimumLength1(string s) {

        vector<int> arr(26, 0);
        for (auto &c : s) {
            arr[c-'a'] += 1;
        }

        int count = 0;
        for (int i = 0; i < 26; i++) {
            if (arr[i] > 2) {
                if (arr[i] % 2 == 0) {
                    count += 2;
                } else {
                    count += 1;
                }
            } else {
                count += arr[i];
            }
        }

        return count;
    }
public:

    int minimumLength(string s) {
        return solveminimumLength1(s);
    }
};

// Driver Code for testing
int main() {

    Solution *sol = new Solution();
    cout << sol->minimumLength("babcb") << "\n";
    cout << sol->minimumLength("aabbaab") << "\n";
    cout << sol->minimumLength("aaaaaaaa") << "\n";
    cout << sol->minimumLength("a") << "\n";
    cout << sol->minimumLength("aa") << "\n";
    cout << sol->minimumLength("aaa") << "\n";
    cout << sol->minimumLength("aba") << "\n";
    cout << sol->minimumLength("abab") << "\n";

    return 0;
}
