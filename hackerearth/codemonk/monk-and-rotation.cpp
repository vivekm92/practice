#include <bits/stdc++.h>

using namespace std;

// https://www.hackerearth.com/problem/algorithm/monk-and-rotation-3-bcf1aefe/
void solve(vector<int> &arr, int k) {
    
    int n = arr.size();
    k = k % n;
    for (int i = n - k; i < 2*n - k; i++) 
        cout << arr[i % n] << " ";
    
    return ;
}

int main() {
    
    int tc, n, k;
    cin >> tc;
    while (tc--) {
        cin >> n >> k;
        vector<int> arr(n);
        for (auto &a : arr) cin >> a;
        solve(arr, k);
        cout << "\n";
    } 
    
    return 0;
}
