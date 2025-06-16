#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/finding-3-digit-even-numbers/description/
class Solution {
private:
    // Approach 1
    // using basic simulation
    // T(n) : O(nlogn) ; S(n) : O(1)
    vector<int> solveFindEvenNumbers(vector<int> &digits) {

        vector<int> digitMap(10, 0);
        for (auto &d : digits) digitMap[d]++;

        vector<int> res;
        for (int i = 100; i < 999; i += 2) {
            int k = i;
            vector<int> temp(10, 0);
            while (k > 0) {
                temp[k%10]++;
                k /= 10;
            }

            bool flag = true;
            for (int j = 0; j < 10; j++) {
                if (temp[j] > digitMap[j]) {
                    flag = false;
                    break;
                }
            }
            temp.clear();

            if (flag) res.push_back(i);
        }

        return res;
    }
public:
    vector<int> findEvenNumbers(vector<int>& digits) {
        return solveFindEvenNumbers(digits);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
