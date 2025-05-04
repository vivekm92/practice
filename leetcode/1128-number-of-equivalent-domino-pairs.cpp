#include <iostream>
#inclide <vector>

using namespace std;

// https://leetcode.com/problems/number-of-equivalent-domino-pairs/description/
class Solution {
private:
    // Approach 1
    // using math
    // T(n) : O(n) ; S(n) : O(1)
    int solveNumEquivDominoPairs(vector<vector<int> >& dominoes) {

        int count = 0;
        vector<int> arr(100, 0);
        for (auto &dominoe : dominoes) {
            int f = dominoe[0], s = dominoe[1], v = 0;
            v = 10 * min(f, s) + max(f, s);
            count += arr[v];
            arr[v]++;
        }

        return count;
    }
public:
    int numEquivDominoPairs(vector<vector<int>>& dominoes) {
        return solveNumEquivDominoPairs(dominoes);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
