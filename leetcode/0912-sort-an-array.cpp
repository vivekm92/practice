#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/sort-an-array/description/
class Solution {
private:
    vector<int> merge(vector<int> &left, vector<int> &right) {
        
        int l2 = left.size(), r2 = right.size();
        if (l2 == 0) return right;
        else if (r2 == 0) return left;
        
        vector<int> temp;
        int l1 = 0, r1 = 0;
        while (l1 < l2 && r1 < r2) {
            if (left[l1] < right[r1]) {
                temp.push_back(left[l1]);
                l1++;
            } else {
                temp.push_back(right[r1]);
                r1++;
            }
        }
        
        while (l1 < l2) {
            temp.push_back(left[l1]);
            l1++;
        }
        
        while (r1 < r2) {
            temp.push_back(right[r1]);
            r1++;
        }
        
        return temp;
    }
    
    vector<int> solveSortArrayHelper(vector<int> &nums, int start, int end) {
        if (start + 1 == end) return {nums[start]};
        
        int mid = start + (end - start) / 2;
        vector<int> lnums = solveSortArrayHelper(nums, start, mid);
        vector<int> rnums = solveSortArrayHelper(nums, mid, end);
        
        vector<int> res = merge(lnums, rnums);
        return res;
    }
    
    // Approach 1
    // using merge-sort
    // T(n) : O(nlogn) ; S(n) : O(n)
    vector<int> solveSortArray(vector<int> &nums) {
        vector<int> sortedNums = solveSortArrayHelper(nums, 0, nums.size());
        return sortedNums;
    }
public:
    vector<int> sortArray(vector<int>& nums) {
        return solveSortArray(nums);
    }
};

// Driver Code for testing
int main() {
    vector<int> arr {5,4,3,2,1};
    Solution *sol = new Solution();
    vector<int> brr = sol->sortArray(arr);
    for (int &b : brr) cout << b << " ";
    cout << "\n";
    return 0;
}

/*
 * Merge Sort
 * Quick Sort
 * Heap Sort
 * Counting Sort
 * Radix Sort
 * Bubble Sort
 * Insertion Sort
 * Selection Sort
 */
