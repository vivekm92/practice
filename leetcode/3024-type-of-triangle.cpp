#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/type-of-triangle/description/
class Solution {
private:
    // Approach 1
    // using basic math
    // T(n) : O(1) ; S(n) : O(1)
    string solveTriangleType(vector<int> &nums) {

        int a = nums[0], b = nums[1], c = nums[2];
        if (a + b <= c || b + c <= a || c + a <= b) return "none";
        else if (a == b && b == c) return "equilateral";
        else if (a == b || b == c || c == a) return "isosceles";
        else return "scalene";
    }
public:
    string triangleType(vector<int>& nums) {
        return solveTriangleType(nums);
    }
};


// Driver Code for testing
int main() {

    return 0;
}
