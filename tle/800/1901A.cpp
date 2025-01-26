#include <iostream>
#include <vector>

using namespace std;

int solve(vector<int>& arr, int n, int x) {
    
    int temp = 0, min_cap = INT_MIN;
    for (int i = 0; i < n; i++) {
        min_cap = max(min_cap, arr[i] - temp);
        temp = arr[i];
    }
    min_cap = max(min_cap, 2 * (x - arr[n-1]));
    
    return min_cap;
}

int main() {
    ios_base::sync_with_stdio(false);
    cin.tie(nullptr);
    
    int t, n, x;
    cin >> t;
    while(t--) {
        cin >> n >> x;
        vector<int> arr(n, 0);
        for (int i = 0; i<n; i++) {
            cin >> arr[i];
        }
        cout << solve(arr, n, x) << "\n";
    }
    
    return 0;
}
