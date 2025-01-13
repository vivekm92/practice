#include <iostream>
#include "trees.h"

using namespace std;

/*
    IB : https://www.interviewbit.com/problems/path-to-given-node
*/ 

// Optimisd Approach 1:
// T(n) : O(n) ; S(n) : O(n)
void solvePathToGivenNode(TreeNode *A, int B, vector<int> temp, vector<int>& res) {
        
    if (A == nullptr) {
        return ;
    }

    temp.push_back(A->val);
    if (A->val == B) {
        res = temp;
    }

    solvePathToGivenNode(A->left, B, temp, res);
    solvePathToGivenNode(A->right, B, temp, res);
}

vector<int> pathToGivenNode(TreeNode *A, int B) {
    
    vector<int> temp, res;
    solvePathToGivenNode(A, B, temp, res);
    return res;
}
/********************************************************************************/

// Driver Code for testing
int main() {


    return 0;
}