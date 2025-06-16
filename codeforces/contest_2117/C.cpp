#include <bits/stdc++.h>

using namespace std;

void solve(vector<int> &arr, int n) {
    
    int count = 0;
    unordered_set<int> curr, seen;
    for (auto &a : arr) {
        curr.insert(a);
        seen.insert(a);
        if (curr.size() == seen.size()) {
            count++;
            seen.clear();
        }
    }
    
    cout << count << "\n";
    return ;
}

int main() {
    
    int tc, n;
    cin >> tc;
    while (tc--) {
        cin >> n;
        vector<int> arr(n);
        for (auto &a : arr) cin >> a;
        solve(arr, n);
    }

    return 0;
}
