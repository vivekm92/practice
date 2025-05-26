#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/largest-color-value-in-a-directed-graph/description/
class Solution {
private:
    // Approach 1
    // using topological sort
    // T(n) : O(m+n) ; S(n) : O(m+n)
    int solveLargestPathValue(string colors, vector<vector<int> > &edges) {

        int n = colors.size();
        vector<vector<int> > adj(n);
        vector<int> indegree(n, 0);

        for (auto &edge : edges) {
            int u = edge[0], v = edge[1];
            adj[u].push_back(v);
            indegree[v]++;
        }

        queue<int> q;
        for (int i = 0; i < n; i++) {
            if (indegree[i] == 0) q.push(i);
        }

        vector<vector<int> > count(n, vector<int>(26, 0));
        int res = 0, seen = 0;
        while (!q.empty()) {
            int node = q.front();
            q.pop();
            count[node][colors[node] - 'a']++;
            res = max(res, count[node][colors[node] - 'a']);
            indegree[node]--;
            seen++;

            for (auto &curr : adj[node]) {
                for (int i = 0; i < 26; i++) {
                    count[curr][i] = max(count[curr][i], count[node][i]);
                }

                indegree[curr]--;
                if (indegree[curr] == 0) {
                    q.push(curr);
                }
            }
        }

        return seen < n ? -1 : res;
    }

    int dfs(int node, string &colors, vector<vector<int> > &adj, vector<vector<int> > &count, vector<bool> &visit, vector<bool> &inStack) {

        if (inStack[node]) return INT_MAX;
        if (visit[node]) return count[node][colors[node]-'a'];

        visit[node] = true;
        inStack[node] = true;
        for (auto &neighbour : adj[node]) {

            if (dfs(neighbour, colors, adj, count, visit, inStack) == INT_MAX) {
                return INT_MAX;
            }

            for (int i = 0; i < 26; i++) {
                count[node][i] = max(count[neighbour][i], count[node][i]);
            }
        }

        count[node][colors[node]-'a']++;
        inStack[node] = false;
        return count[node][colors[node]-'a'];

    }

    // Approach 2
    // using dfs
    // T(n) : O(m+n) ; S(n) : O(m+n)
    int solveLargestPathValue1(string colors, vector<vector<int> > &edges) {

        int n = colors.size();
        vector<vector<int> > adj(n);
        for (auto edge : edges) {
            int u = edge[0], v = edge[1];
            adj[u].push_back(v);
        }

        vector<vector<int> > count(n, vector<int>(26));
        vector<bool> visit(n, false), inStack(n, false);
        int res = 0;
        for (int i = 0; i < n; i++) {
            res = max(res, dfs(i, colors, adj, count, visit, inStack));
        }

        return res == INT_MAX ? -1 : res;
    }
public:
    int largestPathValue(string colors, vector<vector<int>>& edges) {
        return solveLargestPathValue(colors, edges);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
