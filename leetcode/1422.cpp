#include <iostream>

using namespace std;

// https://leetcode.com/problems/maximum-score-after-splitting-a-string/description/
class Solution {
    private:
        // using two for loops : brute-force
        // T(n) : O(n2) ; S(n) : O(1)
        int solvemaxScore1(string s) {
            int n = s.size(), maxCount = 0;
            for (int i = 0; i < n; i++) {
                int countZero = 0;
                for (int j = 0; j <= i; j++) {
                    if (s[j] == '0') {
                        countZero++;
                    }
                }

                int countOne = 0;
                for (int j = i + 1; j < n; j++) {
                    if (s[j] == '1') {
                        countOne++;
                    }
                }

                maxCount = max(maxCount, countZero + countOne);
            }

            return maxCount;
        }

        // using suffix arr
        // T(n) : O(n) ; S(n) : O(n)
        int solvemaxScore(string s) {
            int n = s.size(), countOfOne = 0;
            vector<int> arr(n, 0);
            for (int i = n-1; i >= 0; i--) {
                arr[i] = countOfOne;
                if (s[i] == '1') countOfOne++;
            }

            int count = 0, maxCount = INT_MIN, countofZero = 0;
            for (int i = 0; i < n-1; i++) {
                if (s[i] == '0') countofZero++;
                count = countofZero + arr[i];
                maxCount = max(maxCount, count);
            }

            return maxCount;
        }

        // using two-pass approach
        // T(n) : O(n) ; S(n) : O(1)
        int solvemaxScore2(string s){
            int n = s.size(), maxCount = 0, countOfZero = 0;
            int countOfOne = count(s.begin(), s.end(), '1');
            for (int i = 0; i < n-1; i++) {
                if (s[i] == '0') {
                    countOfZero++;
                } else {
                    countOfOne--;
                }
                maxCount = max(maxCount, countOfOne + countOfZero);
            }
            return maxCount;
        }

        // using single-pass approach
        // T(n) : O(n) ; S(n) : O(1)
        int solvemaxScore3(string s) {
            int n = s.size(), count = INT_MIN;
            int countZero = 0, countOne = 0;
            for (int i = 0; i < n - 1; i++) {
                if (s[i] == '0') {
                    countZero++;
                } else {
                    countOne++;
                }
                count = max(count, countZero - countOne);
            }
            if (s[n-1] == '1') countOne++;
            return count + countOne;
        }
    public:
        int maxScore(string s) {
            return solvemaxScore3(s);
        }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    cout << sol->maxScore("010101") << "\n";
    cout << sol->maxScore("00") << "\n";
    return 0;
}
