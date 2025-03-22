#include <iostream>

using namespace std;

// https://leetcode.com/problems/number-of-provinces/description/
class DSU {
    private:
        vector<int> root, rank;
    public:
        DSU(int n) : root(n), rank(n) {
            for (int i = 0; i < n; i++) {
                root[i] = i;
                rank[i] = 1;
            }
        }
        
        int find(int x) {
            if (x == root[x]) {
                return x;
            }
            root[x] = find(root[x]);
            return root[x];
        }
        
        void union_set(int x, int y) {
            int rootX = find(x);
            int rootY = find(y);
            if (rootX != rootY) {
                if (rank[rootX] < rank[rootY]) {
                    root[rootX] = rootY;
                    rank[rootY] += rank[rootX];
                } else if (rank[rootX] > rank[rootY]) {
                    root[rootY] = rootX;
                    rank[rootX] += rank[rootY];
                } else {
                    root[rootX] = rootY;
                    rank[rootY] += 1;
                }
            }
        }
        
        bool connected(int x, int y) {
            return find(x) == find(y);
        }
        
        vector<int> get_root() {
            return root;
        }
};

class Solution {
private:
    // using disjoint-set
    int solvefindCircleNum(vector<vector<int> >& isConnected) {
        
        int m = isConnected.size(), n = isConnected[0].size();
        DSU dsu(n);
        for (int i = 0; i < m; i++) {
            for (int j = i+1; j < n; j++) {
                if (isConnected[i][j] == 1) {
                    dsu.union_set(i, j);
                }
            }
        }
        
        int count = 0;
        vector<int> root = dsu.get_root();
        for (int i = 0; i < n; i++) {
            if (root[i] == i) {
                count++;
            }
        }
        return count;
    }
public:
    int findCircleNum(vector<vector<int>>& isConnected) {
        return solvefindCircleNum(isConnected);
    }
};


// Driver Code for testing
int main() {

    return 0;
}
