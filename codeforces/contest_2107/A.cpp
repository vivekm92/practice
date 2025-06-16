#include <bits/stdc++.h>

using namespace std;

void solve(vector<int> &arr) {

    int mx = *max_element(arr.begin(), arr.end());
    int mi = *min_element(arr.begin(), arr.end());

    if (mx != mi) {
        cout << "Yes" << "\n";
        for (auto &a : arr) {
            cout << ((a == mx) ? 1 : 2) << " ";
        }
        cout << "\n";
        return ;
    }
    cout << "No" << "\n";
    return ;
}

int main() {

    int tc, n;
    cin >> tc;
    while (tc--) {
        cin >> n;
        vector<int> arr(n);
        for (int i = 0; i < n; i++) {
            cin >> arr[i];
        }
        solve(arr);
        arr.clear();
    }
    return 0;
}
