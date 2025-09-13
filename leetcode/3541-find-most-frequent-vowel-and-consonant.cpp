#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/find-most-frequent-vowel-and-consonant/description/
class Solution {
private:
    // Approach 1
    // T(n) : O(n) ; S(n) : O(1)
    int solveMaxFreqSum(string s) {

        vector<int> arr(26, 0);
        for (auto &ch : s) {
            arr[ch - 'a']++;
        }

        int mcf = 0, mvf = 0;
        for (int i = 0; i < 26; i++) {
            if ((i+'a'=='a') || (i+'a'=='e') || (i+'a'=='i') || (i+'a'=='o') || (i+'a'=='u')) {
                mvf = max(mvf, arr[i]);
            } else {
                mcf = max(mcf, arr[i]);
            }
        }

        return mcf + mvf;
    }
public:
    int maxFreqSum(string s) {
        return solveMaxFreqSum(s);
    }
};

// Driver Code for testing.
int main() {

    return 0;
}
