#include <iostream>
#include "trees.h"

using namespace std;

/*
    IB : https://www.interviewbit.com/problems/symmetric-binary-tree/
    LC : https://leetcode.com/problems/symmetric-tree/description/
*/ 

// Optimised Solution 1:
// T(n) : O(n) ; S(n) : O(n)
bool isSymmetric(TreeNode *A, TreeNode *B) {
    if (A == nullptr && B == nullptr) return true;
    else if (A == nullptr || B == nullptr || A->val != B->val ) return false;
    return isSymmetric(A->left, B->right) && isSymmetric(A->right, B->left);
}

// T(n) : O(n) ; S(n) : O(1)
int solveSymmetricBinaryTree(TreeNode* A) {
    return isSymmetric(A, A) ? 1 : 0;
}
/********************************************************************************/

// Driver Code for testing
int main() {

    return 0;
}