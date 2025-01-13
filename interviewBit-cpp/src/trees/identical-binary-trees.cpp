#include <iostream>
#include "trees.h"

using namespace std;

/*
    IB : https://www.interviewbit.com/problems/identical-binary-trees/
    LC : https://leetcode.com/problems/same-tree/description/
*/ 

/********************************************************************************/
// Optimised Aproach 1:
// T(n) : O(n) ; S(n) : O(n)
int identicalBinaryTrees(TreeNode *A, TreeNode *B) {

    if (A == nullptr && B == nullptr) {
        return 1;
    } else if (A == nullptr || B == nullptr || A->val != B->val) {
        return 0;
    }

    return identicalBinaryTrees(A->left, B->left) && identicalBinaryTrees(A->right, B->right);
}
/********************************************************************************/

// Driver Code for testing
int main() {


    return 0;
}
