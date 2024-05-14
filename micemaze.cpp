#include <bits/stdc++.h>
using namespace std;

// https://www.spoj.com/status/MICEMAZE,vivekm__/
int solve(vector<vector<pair<int, int> > >& graph, int start, int time) {
    auto comparator = [](pair<int, int> p1, pair<int, int> p2){
        return p1.second > p2.second;
    };

    priority_queue<pair<int, int>, vector<pair<int, int> >, decltype(comparator) > pq(comparator);
    vector<int> distances(graph.size(), INT_MAX);
    vector<bool> visited(graph.size(), false);
    pq.push(make_pair(start, 0));
    distances[start] = 0;
    while (!pq.empty()) {
        pair<int, int> p = pq.top();
        pq.pop();
        int u = p.first, distance = p.second;
        if (!visited[u]) {
            visited[u] = true;
            for (pair<int, int> p1 : graph[u]) {
                int v = p1.first, weight = p1.second;
                if (distances[v] > distance + weight) {
                    distances[v] = distances[u] + weight;
                    pq.push(make_pair(v, distances[v])); 
                }
            }
        }
    }

    int miceCount = 0;
    for (int i = 0; i < distances.size(); i++) {
        if (distances[i] <= time) {
            miceCount++;
        }
    }

    return miceCount;
}

int main() { 
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
	
    int n, e, t, m, a, b, w;
    cin >> n >> e >> t >> m;
    vector<vector<pair<int, int> > > maze(n);
    for (int i = 0; i < m; i++) {
        cin >> a >> b >> w;
        maze[a-1].push_back(make_pair(b-1, w));
    }

    // Reverse the edges
    vector<vector<pair<int, int> > > graph(n);
    for (int i = 0; i < maze.size(); i++) {
        int u = i;
        for (int j = 0; j < maze[i].size(); j++) {
            pair<int, int> p = maze[i][j];
            graph[p.first].push_back(make_pair(u, p.second));
        }
    }

    int nMice = solve(graph, e-1, t);
    cout << nMice << "\n";

    graph.clear();
    maze.clear();
    return 0;
}