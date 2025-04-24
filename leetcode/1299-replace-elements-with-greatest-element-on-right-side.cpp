#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/replace-elements-with-greatest-element-on-right-side/description/
class Solution {
private:
    // Approach 1
    // using suffix-array
    // T(n) : O(n) ; S(n) : O(n)
    vector<int> solveReplaceElements(vector<int> &arr) {

        int n = arr.size();
        vector<int> brr(n, 0);
        brr[n-1] = -1;
        for (int i = n-2; i >= 0; i--) {
            brr[i] = max(arr[i+1], brr[i+1]);
        }

        return brr;
    }
public:
    vector<int> replaceElements(vector<int>& arr) {
        return solveReplaceElements(arr);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
