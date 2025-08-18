#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/24-game/description/
class Solution {
private:
    vector<double> helper(double a, double b) {
        vector<double> temp = {a + b, a - b, b - a, a * b};
        if (a != 0) {
            temp.push_back(b / a);
        }
        if (b != 0) {
            temp.push_back(a / b);
        }

        return temp;
    }

    bool backtrack(vector<double> &arr) {

        if (arr.size() == 1) {
            return abs(arr[0] - 24) <= 0.1;
        }

        for (int i = 0; i < arr.size(); i++) {
            for (int j = i + 1; j < arr.size(); j++) {

                vector<double> newArr;
                for (int k = 0; k < arr.size(); k++) {
                    if (k != i && k != j) {
                        newArr.push_back(arr[k]);
                    }
                }

                for (double &val : helper(arr[i], arr[j])) {
                    newArr.push_back(val);
                    if (backtrack(newArr)) {
                        return true;
                    }
                    newArr.pop_back();
                }
            }
        }

        return false;
    }

    bool solveJudgePoint24(vector<int> &cards) {
        vector<double> arr (cards.begin(), cards.end());
        return backtrack(arr);
    }
public:
    bool judgePoint24(vector<int>& cards) {
        return solveJudgePoint24(cards);
    }
};

// Driver Code for testing.
int main() {

    return 0;
}
