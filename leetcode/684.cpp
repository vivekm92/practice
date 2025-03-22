#include <iostream>

using namespace std;

// https://leetcode.com/problems/redundant-connection/description/
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
};

class Solution {
private:
    // using disjoint-set
    // T(n) : O(a*logn) ; S(n) : O(n); where a is acekerman constant
    vector<int> solvefindRedundantConnection(vector<vector<int>>& edges) {
        
        int n = edges.size();
        DSU dsu(n);
        for (auto &edge : edges) {
            int u = edge[0]-1, v = edge[1]-1;
            if (dsu.connected(u, v)) {
                return edge;
            }
            dsu.union_set(u, v);
        }
        return {};
    }
/********************************************************************************/
    // dfs helper
    bool solvefindRedundantConnectionHelper(int start, int end, vector<bool>& visited, vector<vector<int> >& adj) {
        if (start == end) return true;
        visited[start] = true;
        bool isFound = false;
        for (int i = 0; i < adj[start].size(); i++) {
            int curr = adj[start][i];
            if (!isFound && !visited[curr]) {
                isFound = isFound || solvefindRedundantConnectionHelper(curr, end, visited, adj);
            }
        }
        return isFound;
    }
    
    // using dfs
    // T(n) : O(n2) ; S(n) : O(n)
    vector<int> solvefindRedundantConnection1(vector<vector<int>>& edges) {
        
        int n = edges.size();
        vector<vector<int> > adj(n, vector<int>());
        for (auto &edge : edges) {
            vector<bool> visited(n, false);
            int u = edge[0]-1, v = edge[1]-1;
            if (solvefindRedundantConnectionHelper(u, v, visited, adj)) {
                return edge;
            }
            adj[u].push_back(v);
            adj[v].push_back(u);
        }
        return {};
    }
/********************************************************************************/
    
public:
    vector<int> findRedundantConnection(vector<vector<int>>& edges) {
        return solvefindRedundantConnection(edges);
    }
};

// Driver Code for testing
int main() {
    return 0;
}
