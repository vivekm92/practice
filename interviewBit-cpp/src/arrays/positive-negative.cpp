#include <iostream>

using namespace std;

// https://www.interviewbit.com/problems/positive-negative/

// T(n) : O(n) ; S(n) : O(1)
vector<int> positiveNegative(vector<int> &A) {

    int pCount = 0, nCount = 0;
    for (auto &a : A) {
        if (a > 0) pCount++;
        else if (a < 0) nCount++;
    }

    return {pCount, nCount};
}

// Driver Code for testing
int main() {

    return 0;
}
