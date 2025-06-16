#include <bits/stdc++.h>

#define ll long long

using namespace std;

void solve(vector<ll> &arr, ll k) {

    ll minVal = LLONG_MAX, maxVal = LLONG_MIN;
    for (auto &a : arr) {
        minVal = min(minVal, a);
        maxVal = max(maxVal, a);
    }

    int maxCount = 0;
    for (auto &a : arr) maxCount += (maxVal == a);

    if (maxCount == 1) maxVal--;
    if (maxVal - minVal > k) {
        cout << "Jerry\n";
        return ;
    }

    ll total = accumulate(arr.begin(), arr.end(), 0LL);
    if (total % 2 == 1) {
        cout << "Tom" << "\n";
        return ;
    }
    cout << "Jerry" << "\n";
    return ;
}

int main() {

    int tc, n;
    ll k;
    cin >> tc;
    while (tc--) {
        cin >> n >> k;
        vector<ll> arr(n);
        for (int i = 0; i < n; i++) cin >> arr[i];
        solve(arr, k);
        arr.clear();
    }
    return 0;
}
