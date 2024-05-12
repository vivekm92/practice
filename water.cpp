#include <bits/stdc++.h>
#define T tuple<int, int, int> 

using namespace std;

const int dx[4] = {1, -1, 0, 0};
const int dy[4] = {0, 0, 1, -1};

// https://www.spoj.com/status/WATER,vivekm__/
int solve(vector<vector<int> >& graph) {

    int n = graph.size(), m = graph[0].size();
    vector<vector<bool> > floodedCols(n, vector<bool>(m, false));
    priority_queue<T, vector<T>, greater<T>> pq;
    for (int i = 0; i < m; i++) {
        pq.push(make_tuple(graph[0][i], 0, i));
        pq.push(make_tuple(graph[n-1][i], n-1, i));
        floodedCols[0][i] = true;
        floodedCols[n-1][i] = true;
    }

    for (int i = 0; i < n; i++) {
        pq.push(make_tuple(graph[i][0], i, 0));
        pq.push(make_tuple(graph[i][m-1], i, m-1));
        floodedCols[i][0] = true;
        floodedCols[i][m-1] = true;
    }

    int waterTrapped = 0, waterLevel = 0;
    while (!pq.empty()) {
        T p = pq.top();
        int h = get<0>(p), r = get<1>(p), c = get<2>(p);
        pq.pop();
        waterLevel = max(waterLevel, h);
        for (int i = 0; i < 4; i++) {
            int x = r + dx[i], y = c + dy[i];
            if (x >= 0 && x < n && y >= 0 && y < m && !floodedCols[x][y]) {
                int hN = graph[x][y];
                pq.push(make_tuple(hN, x, y));
                if (hN < waterLevel) {
                    waterTrapped += (waterLevel - hN);
                }
                floodedCols[x][y] = true;
            }
        }
    }

    return waterTrapped;
}

int main () {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);

    int t, n, m, v;
    cin >> t ;
    vector<vector<int > > graph;
    while (t--) {
        cin >> n >> m ;
        graph.resize(n, vector<int>(m));
        for (int i=0; i < n; i++) {
            for (int j=0; j <m; j++) {
                cin >> v;
                graph[i][j] = v;
            }
        }

        int trappedWater = solve(graph);
        cout << trappedWater << "\n";
        
        graph.clear();
    }
    return 0;
}