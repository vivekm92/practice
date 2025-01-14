#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/shifting-letters-ii/description/
class Solution {
private:
    // using difference arr
    // T(n) : O(n) ; S(n) : O(n)
    string solveshiftingLetters1(string s, vector<vector<int>>& shifts) {
        int n = s.size();
        vector<int> net_shift(n+1);
        for (auto &shift : shifts) {
            int left = shift[0], right = shift[1], direction = shift[2];
            net_shift[left] = direction == 0 ? net_shift[left] - 1 : net_shift[left] + 1;
            net_shift[right+1] = direction == 0 ? net_shift[right+1] + 1 : net_shift[right+1] - 1;
        }

        for (int i = 0; i < n; i++) {
            net_shift[i + 1] += net_shift[i];
        }

        for (int i = 0; i < n; i++) {
            int shift = net_shift[i] % 26;
            if (shift + s[i] >= 'a' && shift + s[i] <= 'z') {
                s[i] = s[i] + shift;
            } else if (shift + s[i] < 'a') {
                s[i] = 'z' + (shift + s[i] - 'a' + 1);
            } else {
                s[i] = 'a' + (shift + s[i] - 'z' - 1);
            }
        }

        return s;
    }

    // using difference arr : alternate implementation
    // T(n) : O(n) ; S(n) : O(n)
    string solveshiftingLetters(string s, vector<vector<int>>& shifts) {
        int n = s.size();
        vector<int> net_shift(n+1, 0);
        for (auto &shift : shifts) {
            int left = shift[0], right = shift[1], direction = shift[2];
            net_shift[left] = direction == 0 ? net_shift[left] - 1 : net_shift[left] + 1;
            net_shift[right+1] = direction == 0 ? net_shift[right+1] + 1 : net_shift[right+1] - 1;
        }

        for (int i = 0; i < n; i++) {
            net_shift[i + 1] += net_shift[i];
        }

        for (int i = 0; i < n; i++) {
            int shift = (s[i] - 'a' + net_shift[i]) % 26;
            s[i] = 'a' + (shift < 0 ? shift + 26 : shift);
        }

        return s;
    }
public:
    string shiftingLetters(string s, vector<vector<int>>& shifts) {
        return solveshiftingLetters1(s, shifts);
    }
};

// Driver Code for testing
int main() {

    Solution *sol = new Solution();
    vector<vector<int> > shifts = {{0,1,0}, {1,2,1}, {0,2,1}};
    cout << sol->shiftingLetters("abc", shifts) << "\n";

    shifts = {{0,0,0}, {1,1,1}};
    cout << sol->shiftingLetters("dztz", shifts) << "\n";
}

/*
 * s = "abc" shift = [[0, 1, 1], [0, 2, 0]]
 * bcc
 * o/p : abb
 *
 * */
