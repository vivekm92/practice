#include <iostream>

using namespace std;

// https://leetcode.com/problems/pascals-triangle-ii/description/
class Solution {
private:
    // Approach 1
    // using recursion
    // T(n) : O(n) ; S(n) : O(n)
    vector<int> solveGetRow(int rowIndex) {
        if (rowIndex == 0) return {1};
        else if (rowIndex == 1) return {1,1};

        vector<int> temp = solveGetRow(rowIndex - 1);
        vector<int> row {1};
        for (int i = 0; i < temp.size() - 1; i++) {
            row.push_back(temp[i] + temp[i+1]);
        }
        row.push_back(1);

        return row;
    }

    // Approach 2
    // using binomial coefficient
    // T(n) : O(n) ; S(n) : O(n)
    vector<int> solveGetRow1(int rowIndex) {
        if (rowIndex == 0) return {1};
        else if (rowIndex == 1) return {1,1};

        vector<int> row = {1};
        for (int i = 1; i <= rowIndex; i++) {
            long long val = (long long )row[i-1] * (rowIndex - i + 1) / i;
            row.push_back(val);
        }

        return row;
    }
public:
    vector<int> getRow(int rowIndex) {
        return solveGetRow1(rowIndex);
    }
};

// Driver Code for testing
int main() {
    return 0;
}
