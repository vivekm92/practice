#include <iostream> 
#include <vector>
#include <queue>

using namespace std;

int x[] = {2, 2, -2, -2, 1, 1, -1, -1};
int y[] = {1, -1, 1, -1, 2, -2, 2, -2};

// https://www.spoj.com/status/NAKANJ,vivekm__/
vector<pair<int, int> > getAdjacent(vector<vector<int> >& graph, int r, int c) {
    
    vector<pair<int, int> > arr;
    for (int i = 0; i < 8; i++) {
        int nx = r + x[i];
        int ny = c + y[i];
        if (nx >= 0 && nx < 8 && ny >= 0 && ny < 8 && graph[nx][ny] != -1) {
            arr.push_back(make_pair(nx, ny));
        }
    }

    return arr;
} 

int countMoves(vector<vector<int> > &graph, int rs, int re) {

    if (graph[rs][re] == 1) return 0;
    queue<pair<int, int> > q;
    q.push(make_pair(rs, re));
    int moves = 0;
    while (!q.empty()) {
        int sz = q.size();
        for (int i = 0; i < sz; i++) {
            pair<int, int> p = q.front();
            q.pop();
            int currRow = p.first, currCol = p.second;
            graph[currRow][currCol] = -1;
            vector<pair<int, int> > nextMoves = getAdjacent(graph, currRow, currCol);
            for (int i = 0; i < nextMoves.size(); i++) {
                if (graph[nextMoves[i].first][nextMoves[i].second] == 1)
                    return moves + 1;
                q.push(nextMoves[i]);
            }
        }
        moves++;
    }

    return moves;
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    int t, rs, cs, re, ce;
    cin >> t;
    string st, ed;
    vector<vector<int> > graph;
    while (t--) {
        graph.resize(8, vector<int>(8));
        cin >> st >> ed;
        rs = st[0] - 'a';
        re = ed[0] - 'a';
        cs = st[1] - '1';
        ce = ed[1] - '1';
        graph[re][ce] = 1;

        int minMoves = countMoves(graph, rs, cs);
        cout << minMoves << "\n";
        graph.clear();
    }
    return 0;
}