#include <iostream>
#include <vector>
#include <functional>

using namespace std;

// https://leetcode.com/problems/find-numbers-with-even-number-of-digits/description/
class Solution {
private:
    // Approach 1
    // using basic simulation
    // T(n) : O(nlogm) ; S(n) : O(1)
    int solveFindNumbers(vector<int> &nums) {
        
        function<bool(int)> isEvenCount = [&] (int n) {
            int c = 0;
            while (n > 0) {
                c++;
                n /= 10;
            }
            return c % 2 == 0;
        };
        
        int count = 0;
        for (auto &num : nums ) {
            if (isEvenCount(num)) count++;
        }
        
        return count;
    }
    
    // Approach 2
    // using math
    // T(n) : O(nlogn) ; S(n) : O(1)
    int solveFindNumbers1(vector<int> &nums) {
        
        int count = 0;
        for (auto &num : nums) {
            int digitCount = (int )floor(log10(num)) + 1;
            if (digitCount % 2 == 0) count++;
        }
        
        return count;
    }
    
    // Approach 3
    // using string conversion
    // T(n) : O(nlogn) ; S(n) ; O(logn)
    int solveFindNumbers2(vector<int> &nums) {
        
        int count = 0;
        for (auto &num : nums) {
            string s = to_string(num);
            if (s.size() % 2 == 0) count++;
        }
        
        return count;
    }
public:
    int findNumbers(vector<int>& nums) {
        return solveFindNumbers2(nums);
    }
};

// Driver Code for testing
int main() {
    
    return 0;
}
