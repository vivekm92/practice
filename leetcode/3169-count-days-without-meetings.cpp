#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/count-days-without-meetings/description/
class Solution {
private:
    // using sorting
    // T(n) : O(nlogn) ; S(n) : O(1)
    int solveCountDays(int days, vector<vector<int> > &meetings) {
        
        sort(meetings.begin(), meetings.end(), [&](vector<int> &v1, vector<int> &v2){
            if (v1[0] != v2[0]) return v1[0] < v2[0];
            return v1[1] < v2[1];
        });
        
        int start = meetings[0][0], end = meetings[0][1];
        int count = start - 1, n = meetings.size();
        for (int i = 1; i < n; i++) {
            int s = meetings[i][0], e = meetings[i][1];
            if (s > end) {
                count += (s - end - 1);
            }
            start = s;
            end = max(e, end);
        }
        count += (days - end);
        
        return count;
    }
public:
    int countDays(int days, vector<vector<int>>& meetings) {
        return solveCountDays(days, meetings);
    }
};

// Driver Code for testing
int main() {
    
    return 0;
}
