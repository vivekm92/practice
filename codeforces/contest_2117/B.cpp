#include <bits/stdc++.h>

using namespace std;

void solve(int n) {
    
    cout << 1 << " ";
    for (int i = 3; i <= n; i++) {
        cout << i << " ";
    }
    cout << 2 << "\n";
}

int main() {
    
    int tc, n;
    cin >> tc;
    while (tc--) {
        cin >> n;
        solve(n);
    }
    
    return 0;
}
