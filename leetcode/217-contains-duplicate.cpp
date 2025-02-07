#include <iostream>
#include <vector>
#include <unordered_set>

using namespace std;

// https://leetcode.com/problems/contains-duplicate/description/
class Solution {
private:
    // using hash-set
    // T(n) : O(n) ; S(n) : O(n)
     bool solvecontainsDuplicate(vector<int>& nums) {
         unordered_set<int> us;
         for (auto n : nums) {
             if (us.find(n) != us.end()) return true;
             us.insert(n);
         }
         return false;
     }
public:
    bool containsDuplicate(vector<int>& nums) {
        return solvecontainsDuplicate(nums);
    }
};

// Driver Code for testing
int main() {
    return 0;
}
