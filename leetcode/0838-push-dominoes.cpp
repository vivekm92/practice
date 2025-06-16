#include <iostream>
#include <vector>
#include <stack>

using namespace std;

// https://leetcode.com/problems/push-dominoes/description/
class Solution {
private:
    // Approach 1
    // by calculating distance
    // T(n) ; O(n) ; S(n) ; O(n)
    string solvePushDominoes(string dominoes) {
        
        int n = dominoes.size();
        stack<int> st;
        vector<int> nextL(n, -1);
        for (int i = 0; i < n; i++) {
            if (dominoes[i] == 'R') {
                st = stack<int>();
            } else if (dominoes[i] == 'L') {
                while (!st.empty()) {
                    nextL[st.top()] = i;
                    st.pop();
                }
                nextL[i] = 0;
            } else {
                st.push(i);
            }
        }
        
        int prevR = -1;
        for (int i = 0; i < n; i++) {
            if (dominoes[i] == 'R') {
                prevR = i;
            } else if (dominoes[i] == 'L') {
                prevR = -1;
            } else {
                if (prevR == -1) {
                    if (nextL[i] >= 0) dominoes[i] = 'L';
                } else if (nextL[i] == -1) {
                    if (prevR >= 0) dominoes[i] = 'R';
                } else {
                    int rdist = i - prevR;
                    int ldist = nextL[i] - i;
                    if (ldist < rdist) dominoes[i] = 'L';
                    else if (ldist > rdist) dominoes[i] = 'R';
                }
            }
        }
        
        return dominoes;
    }
    
    // Approach 2
    // using stack
    // T(n) : O(n) ; S(n) : O(n)
    string solvePushDominoes1(string dominoes) {
        
        int n = dominoes.size();
        int prevR = -1, nextL = -1;
        stack<pair<int, int> > st;
        for (int i = 0; i < n; i++) {
            if (dominoes[i] == 'R') {
                 while (!st.empty()) {
                    auto [f, s] = st.top();
                    if (prevR != -1) {
                        dominoes[f] = 'R';
                    }
                    st.pop();
                }
                prevR = i;
            } else if (dominoes[i] == 'L') {
                while (!st.empty()) {
                    auto t = st.top();
                    if (t.second == -1) {
                        dominoes[t.first] = 'L';
                    } else {
                        int ldist = i - t.first;
                        int rdist = t.first - t.second;
                        if (ldist < rdist) dominoes[t.first] = 'L';
                        else if (ldist > rdist) dominoes[t.first] = 'R';
                    }
                    st.pop();
                }
                prevR = -1;
            } else {
                st.push({i, prevR});
            }   
        }
        
        while (!st.empty()) {
            auto [f, s] = st.top();
            if (prevR != -1) {
                dominoes[f] = 'R';
            }
            st.pop();
        }
        return dominoes;
    }
    
    // Approach 3
    // using math
    // T(n) : O(n) ; S(n) : O(n)
    int solvePushDominoes2(string dominoes) {
        
        int n = dominoes.size();
        vector<int> 
    }
public:
    string pushDominoes(string dominoes) {
        return solvePushDominoes(dominoes);
    }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    cout << sol->pushDominoes(".L.R...LR..L..") << "\n";
    
    return 0;
}
