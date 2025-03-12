#include <iostream>
#include <vector>

using namespace std;

// https://www.interviewbit.com/problems/greater-than-all/

// T(n) : O(n) ; S(n) : O(1)
int greaterThanAll(vector<int> &A) {

    int maxSoFar = 0, count = 0;
    for (auto &a : A) {
        if (maxSoFar < a) {
            count++;
            maxSoFar = a;
        }
    }

    return count;
}

// Driver Code for testing.
int main() {


}
