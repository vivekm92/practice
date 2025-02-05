#include <iostream>

using namespace std;

/*
    Problem : https://cses.fi/problemset/task/1071
*/

void solve(int n) {
    if (n <= 3) {
        cout << "NO SOLUTION" << "\n";
    }

    vector<int> arr;
    for (int i = 1; i <= n; i+=2) {
        cout << i << " ";
    }
    for (int i =2; i <= n; i+=2) {
        cout << i << " ";
    }
    cout << "\n";
}

int main() {

    int n;
    cin >> n;

    solve(n);

    return 0;
}
