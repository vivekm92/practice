#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/find-the-original-typed-string-i/description/
class Solution {
private:
    // Approach 1
    // basic simulation
    // T(n) : O(n) ; S(n) : O(1)
    int solvePossibleStringCount(string word) {

        int totalCount = 0;
        int idx = 0, n = word.size();
        while (idx < n) {
            int idx1 = idx + 1, cnt = 1;
            while (idx1 < n && word[idx] == word[idx1]) {
                cnt++;
                idx1++;
            }
            totalCount += (cnt - 1);
            idx = idx1;
        }

        return totalCount + 1;
    }

    // Approach 2
    // basic simulation
    // T(n) : O(n) ; S(n) : O(1)
    int solvePossibleStringCount1(string word) {

        int n = word.size(), totalCount = 1;
        for (int i = 1; i < n; i++) {
            if (word[i-1] == word[i]) totalCount++;
        }

        return totalCount;
    }
public:
    int possibleStringCount(string word) {
        return solvePossibleStringCount1(word);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
