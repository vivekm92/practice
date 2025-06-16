#include <bits/stdc++.h>

using namespace std;

string solve(vector<int> &arr, int n, int x) {

    int x1 = -1, x2 = -1;
    for (int i = 0; i < n; i++) {
        if (arr[i] == 1 && x1 == -1) {
            x1 = i;
        }

        if (arr[i] == 1) x2 = i;
    }

    if (x2 - x1 + 1 <= x) return "YES";

    return "NO";
}

int main() {

    int tc, n, x;
    cin >> tc;
    while (tc--) {
        cin >> n >> x;
        vector<int> arr(n);
        for (auto &a : arr) cin >> a;
        cout << solve(arr, n, x) << "\n";
    }
    return 0;
}
