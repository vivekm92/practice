#include <bits/stdc++.h>

using namespace std;

class Solution {
private:
    // Approach 1
    // using disjoint-set
    // T(n) : O(n) ; S(n) : O(1)
    string solveSmallestEquivalentString(string s1, string s2, string baseStr) {
        
        array<int, 26> parent;
        
        // find sets
        function<int(int)> find_sets = [&](int x) {
            if (x == parent[x]) return x;
            
            return parent[x] = find_sets(parent[x]);
        };
        
        // union sets
        function<void(int, int)> union_sets = [&](int x, int y) {
            x = find_sets(x);
            y = find_sets(y);
            
            if (x == y) return ;
            parent[max(x, y)] = min(x, y);
        };
        
        // make parent sets
        for (int i = 0; i < 26; i++) parent[i] = i;
        
        // perform union operation on sets
        int n = s1.size();
        for (int i = 0; i < n; i++) {
            union_sets(s1[i]-'a', s2[i]-'a');
        }
        
        for (auto &ch : baseStr) {
            ch = (char )(find_sets(ch-'a') + 'a');
        }
        
        return baseStr;
    }
    
    // Approach 2
    // using dfs
    // T(n) : O() ; S(n) : O()
    string solveSmallestEquivalentString1(string s1, string s2, string baseStr) {
        return "";
    }
public:
    string smallestEquivalentString(string s1, string s2, string baseStr) {
        return solveSmallestEquivalentString(s1, s2, baseStr);
    }
};

// Driver Code for testing
int main() {
    
    return 0;
}
