#include <iostream>
#include "trees.h"

using namespace std;

/*
    IB : https://www.interviewbit.com/problems/least-common-ancestor/
    LC : https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/description/
*/ 

pair<int, int> solveLeastCommonAncestorHelper(TreeNode* A, int B, int C) {

    if (A == nullptr) {
        return {0, -1};
    }

    pair<int, int> lp = solveLeastCommonAncestorHelper(A->left, B, C);
    if (lp.first == 2) {
        return lp;
    }

    pair<int, int> rp = solveLeastCommonAncestorHelper(A->right, B, C);
    if (rp.first == 2) {
        return rp;
    }

    int num_nodes = lp.first + rp.first;
    if (A->val == B) {
        num_nodes++;
    }

    if (A->val == C) {
        num_nodes++;
    }

    return {num_nodes, num_nodes == 2 ? A->val : -1};
}

// Optimised Approach 1:
// T(n) : O(n) ; S(n) : O(n)
int solveLeastCommonAncestor(TreeNode *A, int B, int C) {

    pair<int, int> p = solveLeastCommonAncestorHelper(A, B, C);
    return p.second;
}
/********************************************************************************/

// Driver Code for testing
int main() {

    return 0;
}