#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/properties-graph/description/
class Solution {
private:
    bool isConnected(vector<vector<int> > &tempList, int i, int j, int k) {

        int count = 0;
        for (int idx = 0; idx < 101; idx++) {
            count += (tempList[i][idx] && tempList[j][idx]);
        }

        return count >= k;
    }

    void dfs(vector<vector<int> >& adjList, vector<bool> &visited, int idx) {
        visited[idx] = true;
        for (int i = 0; i < adjList[idx].size(); i++) {
            if (!visited[adjList[idx][i]]) {
                dfs(adjList, visited, adjList[idx][i]);
            }
        }
    }

    // using dfs
    // T(n) : O(n + e) ; S(n) : O(n) ; n ---> number of nodes
    int solveNumberOfComponents(vector<vector<int> > &properties, int k) {

        int n = properties.size();
        vector<vector<int> > tempList;
        for (auto &vec : properties) {
            vector<int> t(101, 0);
            for (auto &v : vec) {
                t[v]++;
            }
            tempList.push_back(t);
        }

        vector<vector<int> > adjList(n);
        for (int i = 0; i < n; i++) {
            for (int j = i+1; j < n; j++) {
                if (isConnected(tempList, i, j, k)) {
                    adjList[i].push_back(j);
                    adjList[j].push_back(i);
                }
            }
        }

        int count = 0;
        vector<bool> visited(n, false);
        for (int i = 0; i < n; i++) {
            if (!visited[i]) {
                dfs(adjList, visited, i);
                count++;
            }
        }

        return count;
    }

    void bfs(vector<vector<int> > &adjList, vector<bool> &visited, int idx) {

        visited[idx] = true;
        queue<int> q;
        q.push(idx);
        while(!q.empty()) {
            int sz = q.size();
            for (int i = 0; i < sz; i++) {
                idx = q.front();
                q.pop();
                for (int j = 0; j < adjList[idx].size(); j++) {
                    if (!visited[adjList[idx][j]]) {
                        q.push(adjList[idx][j]);
                        visited[adjList[idx][j]] = true;
                    }
                }
            }
        }
    }

    // using bfs
    // T(n) : O(n + e) ; S(n) : O(n) ; n ---> number of nodes
    int solveNumberOfComponents1(vector<vector<int> > &properties, int k) {

        int n = properties.size();
        vector<vector<int> > tempList(n);
        for (int i = 0; i < n; i++) {
            vector<int> t(101, 0);
            for (int j = 0; j < properties[i].size(); j++) {
                t[properties[i][j]]++;
            }
            tempList[i] = t;
        }

        vector<vector<int> > adjList(n);
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                if (isConnected(tempList, i, j, k)) {
                    adjList[i].push_back(j);
                    adjList[j].push_back(i);
                }
            }
        }

        int count = 0;
        vector<bool> visited(n, false);
        for (int i = 0; i < n; i++) {
            if (!visited[i]) {
                bfs(adjList, visited, i);
                count++;
            }
        }

        return count;
    }

    int find_set(int v, vector<int> &parent) {
        if (v == parent[v]) {
            return v;
        }
        parent[v] = find_set(parent[v], parent);
        return parent[v];
    }

    void union_set(int a, int b, vector<int> &parent, vector<int> &rank){
        a = find_set(a, parent);
        b = find_set(b, parent);
        if (a != b) {
            if (rank[a] < rank[b]) {
                swap(a, b);
            }
            parent[b] = a;
            if (rank[a] == rank[b]) {
                rank[a]++;
            }
        }
    }

    // using DSU
    int solveNumberOfComponents2(vector<vector<int> > &properties, int k) {

        int n = properties.size();
        vector<vector<int> > tempList(n);
        for (int i = 0; i < n; i++) {
            vector<int> t(101, 0);
            for (int j = 0; j < properties[i].size(); j++) {
                t[properties[i][j]]++;
            }
            tempList[i] = t;
        }

        vector<int> parent(n, 0), rank(n, 0);
        for (int i = 0; i < n; i++) {
            parent[i] = i;
        }
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                if (isConnected(tempList, i, j, k)) {
                    union_set(i, j, parent, rank);
                }
            }
        }

        int count = 0;
        for (int i = 0; i < n; i++) {
            if (i == find_set(i, parent)) count++;
        }

        return count;
    }
public:
    int numberOfComponents(vector<vector<int>>& properties, int k) {
        return solveNumberOfComponents2(properties, k);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
