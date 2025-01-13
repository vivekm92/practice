#include <iostream>
#include "trees.h"

using namespace std;

// Optimised Approach 1:
// T(n) : O(n) ; S(n) : O(1)
TreeNode* solveNextGreaterNumberBst(TreeNode* A, int B) {

    TreeNode *nextGreaterBstNode = nullptr;
    while (A != nullptr) {
        if (A->val <= B) {
            A = A->right;
        } else {
            nextGreaterBstNode = A;
            A = A->left;
        }
    }
    
    return nextGreaterBstNode;
}

// Driver Code for testing
int main() {

    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);

    
    return 0;
}