#include <bits/stdc++.h>
#define ll long long

using namespace std;

string solve(ll n, vector<ll> &arr) {
    
    if (n == 1) return "YES";
    ll x = (arr[n-1] * n - arr[0]) / (n*n - 1);
    ll y = (arr[0] * n - arr[n-1]) / (n*n - 1);
    
    if (x < 0 || y < 0) return "NO";
    for (int i = 0; i < n; i++) {
        ll v = arr[i] - x * (i + 1) - y * (n - (i + 1) + 1);
        if (v != 0) return "NO";
    }
    
    return "YES";
}

int main() {

    int tc, n;
    cin >> tc;
    while (tc--) {
        cin >> n;
        vector<ll> arr(n);
        for (auto &a : arr) cin >> a;
        cout << solve(n, arr) << "\n";
    }
    
    return 0;
}
