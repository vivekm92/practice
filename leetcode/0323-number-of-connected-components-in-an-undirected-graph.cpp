#include <iostream>
#include <functional>
#include <vector>
#include <queue>

using namespace std;

// https://leetcode.com/problems/number-of-connected-components-in-an-undirected-graph/description/
class Solution {
private:
    // Approach 1
    // using dfs
    // T(n) ; O(E + V) ; S(n) : O(E + V)
    int solveCountComponents(int n, vector<vector<int> > &edges) {

        vector<vector<int> > adjList(n);
        for (auto &edge : edges) {
            int u = edge[0], v = edge[1];
            adjList[u].push_back(v);
            adjList[v].push_back(u);
        }

        vector<bool> visited(n, false);
        function<void(int)> dfs = [&](int node) {
            visited[node] = true;
            for (int i = 0; i < adjList[node].size(); i++) {
                if (visited[adjList[node][i]] == false) {
                    dfs(adjList[node][i]);
                }
            }
        };

        int componentsCount = 0;
        for (int i = 0; i < n; i++) {
            if (visited[i] == false) {
                dfs(i);
                componentsCount++;
            }
        }

        return componentsCount;
    }

    // Approach 2
    // using bfs
    // T(n) : O(E + V) ; S(n) : O(E + V)
    int solveCountComponents1(int n, vector<vector<int> > &edges) {

        vector<vector<int> > adjList(n);
        for (auto &edge : edges) {
            adjList[edge[0]].push_back(edge[1]);
            adjList[edge[1]].push_back(edge[0]);
        }

        vector<bool> visited(n, false);
        function<void(int)> bfs = [&](int node) {
            queue<int> q;
            q.push(node);
            while (!q.empty()) {
                int sz = q.size();
                for (int i = 0; i < sz; i++) {
                    visited[node] = true;
                    node = q.front();
                    q.pop();
                    for (int i = 0; i < adjList[node].size(); i++) {
                        if (visited[adjList[node][i]] == false) {
                            q.push(adjList[node][i]);
                            visited[adjList[node][i]] = true;
                        }
                    }
                }
            }
        };

        int componentCount = 0;
        for (int i = 0; i < n; i++) {
            if (visited[i] == false) {
                bfs(i);
                componentCount++;
            }
        }

        return componentCount;
    }

    // Approach 3
    // using DSU
    // T(n) : O(E + kV) ; S(n) : O(V)
    int solveCountComponents2(int n, vector<vector<int> > &edges) {

        vector<int> parent(n), rank(n, 0);
        for (int i = 0; i < n; i++) parent[i] = i;

        // find_set DSU with Path Compression
        function<int(int)> find_set = [&](int v) {
            if (v == parent[v]) {
                return v;
            }
            return parent[v] = find_set(parent[v]);
        };

        // union_set with union by rank
        function<void(int, int)> union_set = [&](int u, int v) {
            int a = find_set(u), b = find_set(v);
            if (a == b) return;
            if (rank[a] < rank[b]) {
                swap(a, b);
            }
            parent[b] = a;
            if (rank[a] == rank[b]) {
                rank[a]++;
            }
            return ;
        };

        for (auto &edge : edges) {
            int u = edge[0], v = edge[1];
            union_set(u, v);
        }

        int componentCount = 0;
        for (int i = 0; i < n; i++) {
            if (i == find_set(i)) componentCount++;
        }

        return componentCount;
    }
public:
    int countComponents(int n, vector<vector<int>>& edges) {
        return solveCountComponents2(n, edges);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
