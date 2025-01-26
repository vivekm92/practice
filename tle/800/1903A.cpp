#include <iostream>
#include <vector>

using namespace std;

string solve(vector<int>& arr, int n, int k) {
    if (k == 1) {
        for (int i = 1; i < n; i++) {
            if (arr[i] < arr[i-1]) {
                return "NO";
            }
        }
    }
    
    return "YES";
}

int main() {
    int t, n, k;
    cin >> t;
    while(t--) {
        cin >> n >> k;
        vector<int> arr(n, 0);
        for (int i = 0; i < n; i++) {
            cin >> arr[i];
        }
        cout << solve(arr, n, k) << "\n";
    }
    return 0;
}
