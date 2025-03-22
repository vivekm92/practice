#include <iostream>

using namespace std;

// https://leetcode.com/problems/divide-nodes-into-the-maximum-number-of-groups/description/
class Solution {
private:
    
    // using bfs + dsu
    // T(n) : O(n(n+m)) ; S(n) : O(n)
    int solvemagnificentSets1(int n, vector<vector<int> >& edges) {
        return -1;
    }

/********************************************************************************/
    bool isBipartite(vector<vector<int> >& adjList, int node, vector<int>& colors) {
        
        for (auto neighbour : adjList[node]) {
            if (colors[neighbour] == colors[node]) return false;
            if (colors[neighbour] == -1) {
                colors[neighbour] = colors[node] ^ 1;
                if (!isBipartite(adjList, neighbour, colors)) return false;
            }
        }
        
        return true;
    }
    
    int bfs(vector<vector<int> >& adjList, int node, int n) {
        
        queue<int> q;
        q.push(node);
        int distance = 0;
        vector<bool> visited(n, false);
        while (!q.empty()) {
            int sz = q.size();
            for (int i = 0; i < sz; i++) {
                int curr = q.front();
                q.pop();
                visited[curr] = true;
                for (auto neighbour : adjList[curr]) {
                    if (!visited[neighbour]) {
                        q.push(neighbour);
                        visited[neighbour] = true;
                    }
                }
            }
            distance++;
        }
        
        return distance;
    }
    
    int dfs(vector<vector<int> >& adjList, int node, vector<int>& distances, vector<bool>& visited) {
        
        visited[node] = true;
        int maxGroup = distances[node];
        for (auto neighbour : adjList[node]) {
            if (!visited[neighbour]) {
                maxGroup = max(maxGroup, dfs(adjList, neighbour, distances, visited));
            }
        }
        
        return maxGroup;
    }

    // using graph coloring
    // T(n) : O(n(n + m)) ; S(n) : O(n)
    int solvemagnificentSets(int n, vector<vector<int>>& edges) {
        
        vector<vector<int> > adjList(n);
        for (auto &edge : edges) {
            int u = edge[0]-1, v = edge[1]-1;
            adjList[u].push_back(v);
            adjList[v].push_back(u);
        }
        
        vector<int> colors(n, -1);
        for (int i = 0; i < n; i++) {
            if (colors[i] == -1) {
                colors[i] = 0;
                if (!isBipartite(adjList, i, colors)) return -1;
            }
        }
        
        vector<int> distances(n, 0);
        for (int i = 0; i < n; i++) {
            distances[i] += bfs(adjList, i, n);
        }
        
        int maxGroups = 0;
        vector<bool> visited(n, false);
        for (int i = 0; i < n; i++) {
            if (!visited[i]) {
                maxGroups += dfs(adjList, i, distances, visited);
            }
        }
        
        return maxGroups;
    }
public:
    int magnificentSets(int n, vector<vector<int>>& edges) {
        return solvemagnificentSets(n, edges);
    }
};

// Driver Code for testing
int main() {
    
    Solution *sol = new Solution();
    vector<vector<int> > edges {{1,2}, {1,4}, {1,5}, {2,6}, {2,3}, {4,6}};
    cout << sol->magnificentSets(6, edges) << "\n";
    
    return 0;
}
