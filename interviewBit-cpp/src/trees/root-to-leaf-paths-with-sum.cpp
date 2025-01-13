#include <iostream>
#include "trees.h"

using namespace std;

/*
    IB : https://www.interviewbit.com/problems/root-to-leaf-paths-with-sum/
    LC : https://leetcode.com/problems/path-sum-ii/description/
*/ 

void solveRootToLeafPathsWithSum(TreeNode *A, int B, vector<int>& pathList, vector<vector<int> >& allPathList) {

    if (A == nullptr) {
        return;
    }

    pathList.push_back(A->val);
    B -= A->val;
    if (A->left == nullptr && A->right == nullptr) {
        if (B == 0) {
            allPathList.push_back(pathList);
        }
        pathList.pop_back();
        return ;
    } 

    solveRootToLeafPathsWithSum(A->left, B, pathList, allPathList);
    solveRootToLeafPathsWithSum(A->right, B, pathList, allPathList);

    if (pathList.size() > 0) pathList.pop_back();
}

// Optimised Approach 1:
// T(n) : O(n) ; S(n) : O(h)
vector<vector<int > > rootToLeafPathsWithSum(TreeNode *A, int B) {
    vector<int> pathList;
    vector<vector<int> > allPathList;
    solveRootToLeafPathsWithSum(A, B, pathList, allPathList);
    return allPathList;
}
/********************************************************************************/

// Driver Code for testing.
int main() {

    return 0;
}