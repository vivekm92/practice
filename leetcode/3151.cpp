#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/special-array-i/description/
class Solution {
private:
    // T(n) : O(n) ; S(n) : O(1)
    bool solveisArraySpecial(vector<int>& nums) {
        int n = nums.size();
        for (int i = 0; i < n-1; i++) {
            if (nums[i]%2 == nums[i+1]%2) return false;
        }
        return true;
    }
    
    bool solveisArraySpecial1(vector<int>& nums) {
        int n = nums.size();
        for (int i = 0; i < n-1; i++) {
            if (((nums[i]&1) ^ (nums[i+1]&1)) == 0) return false; 
        }
        return true;
    }
public:
    // T(n) : O(n) ; S(n) : O(1)
    bool isArraySpecial(vector<int>& nums) {
        return solveisArraySpecial1(nums);
    }
};

// Driver Code for testing
int main() {
    vector<int> arr {1, 2, 3, 4};
    vector<int> arr1 {1, 2, 3, 5};
    cout << boolalpha;
    Solution *sol = new Solution();
    cout << sol->isArraySpecial(arr) << "\n";
    cout << sol->isArraySpecial(arr1) << "\n";
    
    return 0;
}
