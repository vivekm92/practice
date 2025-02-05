#include <iostream>
#define ll unsigned long long

using namespace std;

/*
    Problem : https://cses.fi/problemset/task/1071
*/

ll solve(ll y, ll x) {
    if (y%2 == 0 && y >= x) return y*y - x + 1;
    else if (x%2 != 0 && x >= y) return x*x - y + 1;
    else if (y > x) return (y-1)*(y-1) + x;
    else return (x-1)*(x-1) + y;
}

int main() {
    ll t, y, x;
    cin >> t;
    while (t--) {
       cin >> y >> x;
        cout << solve(y, x) << "\n";
    }

    return 0;
}
