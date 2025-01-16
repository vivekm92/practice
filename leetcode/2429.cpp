#include <iostream>

using namespace std;

// https://leetcode.com/problems/minimize-xor/description/
class Solution {
private:
    // using bits manipulation : cleaner implementation
    // T(n) : O(logn) ; S(n) : O(1)
    int solveminimizeXor2(int num1, int num2) {
        int target = __builtin_popcount(num2);
        int current = __builtin_popcount(num1);
        int currentBit = 0;
        // set Bits at LSB
        while (current < target) {
            if (num1 & (1<<currentBit)) {
                currentBit++;
                continue;
            }
            num1 = num1 | (1<<currentBit);
            current++;
            currentBit++;
        }

        // unset Bits at LSB
        while (current > target) {
            if (num1 & (1<<currentBit)) {
                num1 = num1 & ~(1<<currentBit);
                current--;
            }
            currentBit++;
        }

        return num1;
    }

    int countSetBits(int num) {
        int count = 0;
        while (num > 0) {
            if (num & 1) {
                count++;
            }
            num = num >> 1;
        }
        return count;
    }
    // using bits manipulation
    // T(n) : O(logn) ; S(n) : O(1)
    int solveminimizeXor1(int num1, int num2) {
        int set_bits_num1 = countSetBits(num1);
        int set_bits_num2 = countSetBits(num2);
        if (set_bits_num1 == set_bits_num2) {
            return num1;
        } else if (set_bits_num1 > set_bits_num2) {
            int diff = set_bits_num1 - set_bits_num2, i = 0;
            while (diff > 0) {
                if (num1 & (1<<i)) {
                    num1 = num1 & ~(1<<i);
                    diff--;
                }
                i++;
            }
            return num1;
        }
        int diff = set_bits_num2 - set_bits_num1, i = 0;
        while (diff > 0) {
            if (num1 & (1<<i)) {
                i++;
                continue;
            }
            num1 = num1 | (1<<i);
            i++;
            diff--;
        }

        return num1;
    }

    // using bits manipulation and builtin methods
    // T(n) : O(logn) ; S(n) : O(1)
    int solveminimizeXor(int num1, int num2) {
        int set_bits_num1 = __builtin_popcount(num1);
        int set_bits_num2 = __builtin_popcount(num2);
        if (set_bits_num1 == set_bits_num2) {
            return num1;
        } else if (set_bits_num1 > set_bits_num2) {
            int diff = set_bits_num1 - set_bits_num2, i = 0;
            while (diff > 0) {
                if (num1>>i & 1) {
                    num1 = num1 & ~(1 << i);
                    diff--;
                }
                i++;
            }
            return num1;
        }
        int diff = set_bits_num2 - set_bits_num1, i = 0;
        while (diff > 0) {
            if (num1>>i & 1) {
                i++;
                continue;
            }
            num1 = num1 | (1<<i);
            diff--;
            i++;
        }

        return num1;
    }
public:
    int minimizeXor(int num1, int num2) {
        return solveminimizeXor2(num1, num2);
    }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    cout << sol->minimizeXor(3, 5) << "\n";
    cout << sol->minimizeXor(72, 25) << "\n";
    cout << sol->minimizeXor(25, 72) << "\n";
}
