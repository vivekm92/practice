#include <iostream>

using namespace std;

// https://leetcode.com/problems/gray-code/description/
class Solution {
private:
    // T(n) : O(2^n) ; S(n) : O(n)
    vector<int> solvegrayCode(int n) {
        if (n == 0) return {0};
        if (n == 1) return {0, 1};
        vector<int> grayCode = solvegrayCode(n-1);
        int k = grayCode.size(), temp = grayCode[k-1] & -grayCode[k-1];
        for (int i = k-1; i >= 0; i--) {
            grayCode.push_back(grayCode[i] + temp * 2);
        }
        return grayCode;
    }
public:
    vector<int> grayCode(int n) {
        return solvegrayCode(n);
    }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    vector<int> r = sol->grayCode(5);
    for (auto &v : r) cout << v << " ";
    cout << "\n";
    return 0;
}
