#include <iostream>

using namespace std;

// https://leetcode.com/problems/count-symmetric-integers/description/
class Solution {
private:
    // Approach 1
    // brute-force
    // T(n) : O(nlogn) ; S(n) : O(1)
    int solveCountSymmetricIntegers(int low, int high) {

       int count = 0;
       for (int i = low; i <= high; i++) {
           int num = i, sum = 0, cntDigits = 0;
           while (num > 0) {
               sum += num % 10;
               num /= 10;
               cntDigits++;
           }

           if (cntDigits % 2 != 0) continue;
           num = i;
           int newSum = 0, tempCount = cntDigits / 2;
           while (tempCount > 0 && num > 0) {
               newSum += num % 10;
               num /= 10;
               tempCount--;
           }

           if (sum == 2 * newSum) count++;
       }

       return count;
    }
public:
    int countSymmetricIntegers(int low, int high) {
        return solveCountSymmetricIntegers(low, high);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
