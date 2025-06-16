#include <iostream>
#include <set>
#include <vector>

using namespace std;

// https://leetcode.com/problems/maximum-number-of-tasks-you-can-assign/description/
class Solution {
private:
    bool check(vector<int> &tasks, vector<int> &workers, int pills, int strength, int t) {
        
        int m = workers.size();
        multiset<int> ws;
        for (int i = m - t; i < m); i++)
            ws.insert(workers[i]);
            
        for (int i = t-1; i >= 0; i--) {
            if (auto it = prev(ws.end()); *it >= tasks[i]) {
                ws.erase(it);
            } else {
                if (pills == 0) return false;
                auto rep = ws.lower_bound(tasks[i] - strength);
                if (rep == ws.end()) return false;
                pills--;
                ws.erase(rep);
            }
        }
        
        return true;
    }

    int solveMaxTaskAssign(vector<int> &tasks, vector<int> &workers, int pills, int strength) {
        
        sort(tasks.begin(), tasks.end());
        sort(workers.begin(), workers.end());
        
        int n = tasks.size(), m = workers.size();
        int l = 0, r = min(n, m);
        while (l < r) {
            int mid = l + (r - l) / 2;
            if (check(tasks, workers, pills, strength, mid)) {
                l = mid + 1;
            } else {
                r = mid;
            }
        }
        
        return l - 1;
    }
    
public:
    int maxTaskAssign(vector<int>& tasks, vector<int>& workers, int pills, int strength) {
        return solveMaxTaskAssign(tasks, workers, pills, strength);
    }
};

// Driver Code for testing
int main() {
    
    return 0;
}
