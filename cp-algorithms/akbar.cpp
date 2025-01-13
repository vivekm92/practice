#include <iostream> 
#include <vector>
#include <queue>

using namespace std;

// https://www.spoj.com/status/AKBAR,vivekm__/
void bfs(vector<vector<int>>& graph, vector<bool>& visited, int u, int strength) {

	if (strength == 0) {
		visited[u] = true;
		return;
	}

	queue<int> q;
	q.push(u);
	visited[u] = true;
	while (!q.empty()) {
		int sz = q.size();
		strength--;
		for (int i = 0; i < sz; i++) {	
			int v = q.front();
			q.pop();
			for (int i=0; i<graph[v].size(); i++) {
				int u = graph[v][i];
				if (!visited[u]) {
					visited[u] = true;
					q.push(u);
				}
			}
		}

		if (strength == 0) break;
	}

	return ;
}

bool solve(vector<vector<int>>& graph, vector<int>& strength) {
	
	int n = strength.size();
	vector<bool> visited(n);
	for (int i=1; i<n; i++) {
		if (strength[i] >= 0) {
			if (visited[i]) return false;
			bfs(graph, visited, i, strength[i]);
		}
	}

	// If any of the node is not visited, then return false;
	for (int i=1; i<n; i++) {
		if (!visited[i]) return false;
	}

	return true;
}

int main() {
	ios_base::sync_with_stdio(false);
	cin.tie(nullptr);

	int t,n,r,m,r1,r2,m1,m2;
	vector<vector<int> > graph;
	vector<int> strength;
	cin >> t;
	while (t--) {
		cin >> n >> r >> m;
		graph.resize(n+1);
		strength.resize(n+1, -1);
		while (r--) {
			cin >> r1 >> r2;
			graph[r1].push_back(r2);
			graph[r2].push_back(r1);
		}
		while (m--) {
			cin >> m1 >> m2;
			strength[m1] = m2;
		}

		bool flag = solve(graph, strength);
		if (flag) cout << "Yes\n";
		else cout << "No\n";
		graph.clear();
		strength.clear();
	}

	return 0;
}

