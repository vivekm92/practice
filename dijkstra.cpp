#include <bits/stdc++.h>
using namespace std;


typedef pair<int, int> NodeAndDistance;
typedef vector<vector<NodeAndDistance> > AdjacencyList;

class NoPathExistsException: public runtime_error {
public:
    NoPathExistsException() : runtime_error("No path exists between the nodes.") {}
};

class Dijkstra {
private:
    AdjacencyList adjacenctList;
    vector<int> distances;
    vector<int> parents;
    int UNKNOWN = -1;

    void runDijkstra(int start) {
        auto comparator = [](NodeAndDistance p1, NodeAndDistance p2) {
            return p1.second > p2.second;
        };
        
        priority_queue<NodeAndDistance, vector<NodeAndDistance>, decltype(comparator)> distanceQueue(comparator);
        vector<bool> visited(this->adjacenctList.size(), false);
        distanceQueue.push(make_pair(start, 0));
        parents[start] = start;
        distances[start] = 0;

        while (!distanceQueue.empty()) {
            // auto [u, distance] = distanceQueue.top();
            pair<int, int> p = distanceQueue.top();
            int u = p.first, distance = p.second;
            distanceQueue.pop();

            if (visited[u]) continue;
            visited[u] = true;

            for (pair<int, int> p1 : this->adjacenctList[u]) {
                int v = p1.first, weight = p1.second;
                if (distances[v] == UNKNOWN || distance + weight < distances[v]) {
                    distances[v] = distances[u] + weight;
                    parents[v] = u;
                    distanceQueue.push(make_pair(v, distances[v]));
                }
            }
        }

    }

public:
    Dijkstra(AdjacencyList &_adjacenctList)
    : adjacenctList {_adjacenctList}
    , distances {vector<int>(_adjacenctList.size(), UNKNOWN)}
    , parents {vector<int>(_adjacenctList.size(), UNKNOWN)}
    {}

    pair <int, vector<int> > computeShortestPath(int start, int end) {
        this->runDijkstra(start);
        if (this->parents[end] == UNKNOWN) {
            throw NoPathExistsException();
        }

        int distance = this->distances[end];
        vector<int> path;
        path.push_back(end);
        int current = end;
        while (current != start) {
            current = parents[current];
            path.push_back(current);
        }

        reverse(path.begin(), path.end());
        return make_pair(distance, path);
    }
};

int main() {
    int n = 5;
    AdjacencyList adjacencyList(n);
    vector<tuple<int, int, int>> edges {
        make_tuple(0, 1, 6),
        make_tuple(0, 2, 10),
        make_tuple(0, 3, 4),
        make_tuple(1, 0, 2),
        make_tuple(1, 2, 3),
        make_tuple(1, 4, 5),
        make_tuple(3, 1, 1),
        make_tuple(3, 4, 2),
        make_tuple(4, 2, 1)
    };
    for (tuple<int, int, int> t : edges) {
        int u = get<0>(t), v = get<1>(t), w = get<2>(t);
        adjacencyList[u].push_back(make_pair(v, w));
    }

    Dijkstra dijkstra(adjacencyList);
    pair<int, vector<int> > p = dijkstra.computeShortestPath(0, 2);
    int distance = p.first;
    vector<int> path02 = p.second;
    cout << "Shortest path from 0 to 2: ";
    for (int u : path02) {
        cout << u << " ";
    }
    cout << endl << "Length of the path: " << distance;
}