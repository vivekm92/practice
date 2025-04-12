#include <iostream>
#include <queue>
#include <vector>

using namespace std;

// https://leetcode.com/problems/shortest-distance-from-all-buildings/description/
class Solution {
private:

    int bfs(vector<vector<int> >& grid, int row, int col, int totalHouses) {

        vector<int> dx {0, 0, 1, -1};
        vector<int> dy {1, -1, 0, 0};
        int nR = grid.size(), nC = grid[0].size();
        queue<pair<int, int> > q;
        q.push({row, col});


        int housesVisited = 0, totalDistance = 0, houseDistance = 0;
        vector<vector<bool> > visited(nR, vector<bool>(nC, false));
        visited[row][col] = true;
        while (!q.empty() && housesVisited < totalHouses) {
            int sz = q.size();
            for (int i = 0; i < sz; i++) {
                auto curr = q.front();
                q.pop();

                row = curr.first;
                col = curr.second;
                if (grid[row][col] == 1) {
                    totalDistance += houseDistance;
                    housesVisited++;
                    continue;
                }

                for (int j = 0; j < 4; j++) {
                    int nRow = row + dx[j];
                    int nCol = col + dy[j];

                    if (nRow >= 0 && nRow < nR && nCol >= 0 && nCol < nC) {
                        if (!visited[nRow][nCol] && grid[nRow][nCol] != 2) {
                            visited[nRow][nCol] = true;
                            q.push({nRow, nCol});
                        }
                    }
                }
            }
            houseDistance++;
        }

        if (housesVisited < totalHouses) {
            for (int i = 0; i < nR; i++) {
                for (int j = 0; j < nC; j++) {
                    if (grid[i][j] == 0 && visited[i][j]) {
                        grid[i][j] = 2;
                    }
                }
            }
            return INT_MAX;
        }

        return totalDistance;
    }

    // Approach 1
    // bfs from empty land to houses
    // T(n) : O(n2m2) ; S(n) : O(nm)
    int solveShortestDistance(vector<vector<int> > &grid) {

        int totalHouses = 0;
        int nR = grid.size(), nC = grid[0].size();
        for (int i = 0; i < nR; i++) {
            for(int j = 0; j < nC; j++) {
                if (grid[i][j] == 1) totalHouses++;
            }
        }

        int minDistance = INT_MAX;
        for (int i = 0; i < nR; i++) {
            for (int j = 0; j < nC; j++) {
                if (grid[i][j] == 0) {
                    minDistance = min(minDistance, bfs(grid, i, j, totalHouses));
                }
            }
        }

        if (minDistance == INT_MAX) return -1;
        return minDistance;
    }

    void bfs1(vector<vector<int> > &grid, int row, int col, vector<vector<int> > &distances, vector<vector<int> > &housesVisited) {

        vector<int> dx {1, -1, 0, 0};
        vector<int> dy {0, 0, 1, -1};

        int nR = grid.size(), nC = grid[0].size();
        vector<vector<bool> > visited(nR, vector<bool>(nC, false));
        queue<pair<int, int> > q;
        q.push({row, col});
        visited[row][col] = true;
        int curr_distance = 0;
        while (!q.empty()) {
            int sz = q.size();
            for (int i = 0; i < sz; i++) {
                auto curr = q.front();
                q.pop();
                row = curr.first;
                col = curr.second;
                if (grid[row][col] == 0) {
                    distances[row][col] += curr_distance;
                    housesVisited[row][col] += 1;
                }

                for (int i = 0; i < 4; i++) {
                    int nRow = row + dx[i];
                    int nCol = col + dy[i];

                    if (nRow >= 0 && nRow < nR && nCol >= 0 && nCol < nC) {
                        if (!visited[nRow][nCol] && grid[nRow][nCol] == 0) {
                            q.push({nRow, nCol});
                            visited[nRow][nCol] = true;
                        }
                    }
                }
            }
            curr_distance++;
        }

    }

    // Approach 2
    // bfs from houses to empty land
    // T(n) : O(n2m2) ; S(n) : O(nm)
    int solveShortestDistance1(vector<vector<int> > &grid) {

        int nR = grid.size(), nC = grid[0].size();
        vector<vector<int> > partialDistances(nR, vector<int>(nC, 0));
        vector<vector<int> > housesVisited(nR, vector<int>(nC, 0));
        int totalHouses = 0;
        for (int i = 0; i < nR; i++) {
            for (int j = 0; j < nC; j++) {
                if (grid[i][j] == 1) {
                    bfs1(grid, i, j, partialDistances, housesVisited);
                    totalHouses++;
                }
            }
        }

        int minDistance = INT_MAX;
        for (int i = 0; i < nR; i++) {
            for (int j = 0; j < nC; j++) {
                if (grid[i][j] == 0 && housesVisited[i][j] == totalHouses) {
                    minDistance = min(minDistance, partialDistances[i][j]);
                }
            }
        }

        if (minDistance == INT_MAX) return -1;
        return minDistance;
    }

    // Approach 3
    // using optimised bfs
    // T(n) : O(n2m2) ; S(n) : O(nm)
    int solveShortestDistance2(vector<vector<int> > &grid) {

        vector<int> dx {0, 0, 1, -1};
        vector<int> dy {1, -1, 0, 0};

        int nR = grid.size(), nC = grid[0].size();
        vector<vector<int> > distances(nR, vector<int>(nC, 0));
        int minDistance = INT_MAX, emptyLandVal = 0;
        for (int i = 0; i < nR; i++) {
            for (int j = 0; j < nC; j++) {
                if (grid[i][j] == 1) {
                    queue<pair<int, int> > q;
                    q.push({i, j});
                    minDistance = INT_MAX;
                    int curr_distance = 0;
                    while (!q.empty()) {
                        curr_distance++;
                        int sz = q.size();
                        for (int k = 0; k < sz; k++) {
                            auto curr = q.front();
                            q.pop();
                            int row = curr.first, col = curr.second;
                            for (int x = 0; x < 4; x++) {
                                int nRow = row + dx[x];
                                int nCol = col + dy[x];
                                if (nRow >= 0 && nRow < nR && nCol >= 0 && nCol < nC && grid[nRow][nCol] == emptyLandVal) {
                                    grid[nRow][nCol]--;
                                    distances[nRow][nCol] += curr_distance;
                                    q.push({nRow, nCol});
                                    minDistance = min(minDistance, distances[nRow][nCol]);
                                }
                            }
                        }
                    }
                    emptyLandVal--;
                }
            }
        }

        return minDistance == INT_MAX ? -1 : minDistance;
    }
public:
    int shortestDistance(vector<vector<int>>& grid) {
        return solveShortestDistance2(grid);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
