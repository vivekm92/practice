#include <iostream>
#include "trees.h"

using namespace std;

/*
    IB : https://www.interviewbit.com/problems/path-sum/
    LC : https://leetcode.com/problems/path-sum/description/
*/ 


// Optimised Approach 1:
// T(n) : O(n) ; S(n) : O(h)
int pathSum(TreeNode *A, int B) {

    if (A == nullptr) return 0;

    B -= A->val;
    if (A->left == nullptr && A->right == nullptr) {
        return B == 0;
    }

    return pathSum(A->left, B) || pathSum(A->right, B);
}
/********************************************************************************/

// Driver Code for testing
int main() {

    return 0;
}