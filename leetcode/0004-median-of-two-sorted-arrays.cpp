#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/median-of-two-sorted-arrays/description/
class Solution {
private:
    // using two-pointers
    // T(n) : O(n+m) ; S(n) : O(1)
    double solvefindMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
       int i1=0, i2=0, n1=nums1.size(), n2=nums2.size();
       
    }
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        return solvefindMedianSortedArrays(nums1, nums2);
    }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    vector<int> arr1 {1, 2, 4};
    vector<int> arr2 {1, 3, 5};
    cout << sol->findMedianSortedArrays(arr1, arr2) << "\n";
    
    return 0;
}
