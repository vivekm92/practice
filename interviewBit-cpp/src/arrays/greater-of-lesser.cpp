#include <iostream>
#include <vector>

using namespace std;

// https://www.interviewbit.com/problems/greater-of-lesser/

// Approach 1
// T(n) : O(n) ; S(n) : O(1)
int greaterOfLesser(vector<int>& A, vector<int> &B, int C) {

    int gCount = 0, lCount = 0;
    for (auto &a : A) {
        if (a > C) gCount++;
    }

    for (auto &b : B) {
        if (b < C) lCount++;
    }

    return max(gCount, lCount);
}

// Driver Code for testing
int main() {

    return 0;
}
