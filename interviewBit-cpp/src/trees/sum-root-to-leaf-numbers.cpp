#include <iostream>
#include "trees.h"

using namespace std;

/*
    IB : https://www.interviewbit.com/problems/sum-root-to-leaf-numbers/
    LC : https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/description/
*/ 

// Optimised Approach 1:
// T(n) : O(n) ; S(n) : O(n)
int solveSumRootToLeafNumbers(TreeNode* A, int sum) {
    
    if (A == nullptr) {
        return 0;
    }

    sum = ((sum % 1003) * 10) % 1003;
    sum = (sum % 1003 + A->val) % 1003;
    if (A->left == nullptr && A->right == nullptr) {
        return sum;
    }
    return (solveSumRootToLeafNumbers(A->left, sum) % 1003 + solveSumRootToLeafNumbers(A->right, sum) % 1003) % 1003; 
}

int sumRootToLeafNumbers(TreeNode* A) {
    return solveSumRootToLeafNumbers(A, 0);
}
/********************************************************************************/

// Driver Code for testing
int main() {


    return 0;
}