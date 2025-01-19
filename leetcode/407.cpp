#include <iostream>

using namespace std;

// https://leetcode.com/problems/trapping-rain-water-ii/description/
class Solution {
private:
    class Cell {
    public:
        int h, x, y;
        Cell(int h, int x, int y) : h(h), x(x), y(y) {}
        bool operator<(const Cell& other) const {
            return h >= other.h;
        }
    };
    bool isvalid(int x, int y, int nR, int nC) {
        return x >= 0 && x < nR && y >= 0 && y < nC;
    }

    // T(n) : O(logn) ; S(n) : O(n)
    int solvetrapRailWater(vector<vector<int> >& heightMap) {
        int dx[4] = {0, 0, -1, 1};
        int dy[4] = {1, -1, 0, 0};
        int nR = heightMap.size(), nC = heightMap[0].size();
        vector<vector<bool> > visited(nR, vector<bool>(nC, false));
        priority_queue<Cell> boundary;
        for (int i = 0; i < nR; i++) {
            boundary.push(Cell(heightMap[i][0], i, 0));
            boundary.push(Cell(heightMap[i][nC-1], i, nC-1));
            visited[i][0] = true;
            visited[i][nC-1] = true;
        }
        
        for (int i = 0; i < nC; i++) {
            boundary.push(Cell(heightMap[0][i], 0, i));
            boundary.push(Cell(heightMap[nR-1][i], nR-1, i));
            visited[0][i] = true;
            visited[nR-1][i] = true;
        }
        
        int trappedWater = 0;
        while (!boundary.empty()) {
            Cell curr = boundary.top();
            boundary.pop();
            int h = curr.h, x = curr.x, y = curr.y;
            for (int i = 0; i < 4; i++) {
                int nx = x + dx[i];
                int ny = y + dy[i];
                if (isvalid(nx, ny, nR, nC) && !visited[nx][ny]) {
                    int nh = heightMap[nx][ny];
                    if (nh < h) {
                        trappedWater += (h - nh);
                    }
                    boundary.push(Cell(max(nh, h), nx, ny));
                    visited[nx][ny] = true;
                }
            }
        }
        
        return trappedWater;
    }
public:
    int trapRainWater(vector<vector<int>>& heightMap) {
        return solvetrapRailWater(heightMap);
    }
};

int main() {
    
    return 0;
}
