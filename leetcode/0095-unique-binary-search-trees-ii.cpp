#include <iostream>
#include <vector>
#include <map>

using namespace std;

// Definition for a binary tree node.
    struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

// https://leetcode.com/problems/unique-binary-search-trees-ii/description/
class Solution {
private:
    vector<TreeNode* > allPossibleBst(int start, int end, map<pair<int, int>, vector<TreeNode* > > &um) {
        if (start > end) return {nullptr};

        pair<int, int> p = {start, end};
        if (um.find(p) != um.end()) return um[p];
        vector<TreeNode* > res;
        for (int i = start; i <= end; i++) {
            auto left = allPossibleBst(start, i - 1, um);
            auto right = allPossibleBst(i + 1, end, um);

            for (auto &l : left) {
                for (auto &r : right) {
                    TreeNode *root = new TreeNode(i, l, r);
                    res.push_back(root);
                }
            }
        }

        return um[p] = res;
    }

    // Approach 1
    // using top-down DP
    // T(n) : O() ; S(n) : O()
    vector<TreeNode* > solveGenerateTrees(int n) {
        map<pair<int, int>, vector<TreeNode* > > um;
        return allPossibleBst(1, n, um);
    }
public:
    vector<TreeNode*> generateTrees(int n) {
        return solveGenerateTrees(n);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
