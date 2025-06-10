#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/maximum-difference-between-even-and-odd-frequency-i/description/
class Solution {
private:
    // Approach 1
    // basic simulation
    // T(n) : O(n) ; S(n) ; O(26)
    int solveMaxDifference(string s) {

        array<int, 26> arr;
        for (auto &ch : s) arr[ch-'a']++;

        int maxOdd = 0, minEven = s.size();
        for (int i = 0; i < 26; i++) {
            if (arr[i] > 0) {
                if (arr[i] % 2 == 0) minEven = min(minEven, arr[i]);
                else maxOdd = max(maxOdd, arr[i]);
            }
        }

        return maxOdd - minEven;
    }
public:
    int maxDifference(string s) {
        return solveMaxDifference(s);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
