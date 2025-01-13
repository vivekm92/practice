#include <iostream>
#include "trees.h"

using namespace std;

/*
    IB : https://www.interviewbit.com/problems/balanced-binary-tree/
    LC : https://leetcode.com/problems/balanced-binary-tree/description/
*/

/********************************************************************************/
// T(n) : O(n) ; S(n) : O(n) 
int nodeHeight(TreeNode *A) {

    if (A == nullptr) {
        return 0;
    }

    return max(nodeHeight(A->left), nodeHeight(A->right)) + 1;
}

// Brute-Force Solution 1:
// T(n) : O(n2) : S(n) : O(n)
int balancedBinaryTree(TreeNode* A) {

    if (A == nullptr) return 1;
    
    int lDiff = balancedBinaryTree(A->left);
    int rDiff = balancedBinaryTree(A->right);

    int lh = nodeHeight(A->left);
    int rh = nodeHeight(A->right);

    int diff = abs(lh - rh);
    return max(max(lDiff, rDiff), diff);
}
/********************************************************************************/

// Optimised Approach 1:
// T(n) : O(n) ; S(n) : O(n)
pair<int, int> checkBalanacedBinaryTree(TreeNode* A) {

    if (A == nullptr) {
        return {0, 0};
    }

    pair<int, int> lp = checkBalanacedBinaryTree(A->left);
    pair<int, int> rp = checkBalanacedBinaryTree(A->right);

    pair<int, int> p;
    p.first = max(lp.first, rp.first) + 1;
    p.second = max(max(lp.second, rp.second), abs(lp.first - rp.first));
    return p;
}

int balancedBinaryTree(TreeNode* A) {
    
    pair<int, int> p = checkBalanacedBinaryTree(A);   
    return p.second <= 1 ? 1 : 0;
}
/********************************************************************************/

// Optimised Approach 2:
// T(n) : O(n) ; S(n) : O(n)
int checkBalanacedBinaryTree1(TreeNode *A) {
    
    if (A == nullptr) {
        return 0;
    }

    int lh = checkBalanacedBinaryTree1(A->left);
    if (lh == -1) return -1;

    int rh = checkBalanacedBinaryTree1(A->right);
    if (rh == -1) return -1;

    if (abs(lh - rh) > 1) {
        return -1;
    }
    
    return max(lh, rh) + 1;
}

int balancedBinaryTree1(TreeNode *A) {
    return checkBalanacedBinaryTree1(A) != -1 ? 1 : 0;
}
/********************************************************************************/

// Optimised Approach 3:
// T(n) : O(n) ; S(n) : O(n)
pair<int, bool> checkBalanacedBinaryTree2(TreeNode *A) {

    if (A == nullptr) {
        return {0, true};
    }

    pair<int, int> lp = checkBalanacedBinaryTree2(A->left);
    if (!lp.second) {
        return {0, false};
    }

    pair<int, int> rp = checkBalanacedBinaryTree2(A->right);
    if (!rp.second) {
        return {0, false};
    }

    return {max(lp.first, rp.first) + 1, abs(lp.first - rp.first) <= 1};
}

int balancedBinaryTree2(TreeNode *A) {
    return checkBalanacedBinaryTree2(A).second ? 1 : 0;
}
/********************************************************************************/

// Driver Code for testing
int main() {


    return 0;
}