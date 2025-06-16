#include <bits/stdc++.h>

using namespace std;

// T(n) : O(logn) ; S(n) : O(1)
long long binpow(long long a, long long b) {
    
    long long res = 1;
    while (b > 0) {
        if (b & 1) res = res * a;
        a = a * a;
        b >>= 1;
    }
    return res;
}

// T(n) : O(logn) ; S(n) : O(1) recursive call stack utilises O(logn) 
long long binpow_recursive(long long a, long long b) {
    
    if (b == 0) return 1;
    long long res = binpow_recursive(a, b/2);
    if (b & 1) return res * res * a;
    return res * res; 
}

// T(n) : O(logn) ; S(n) : O(1)
long long binpow_mod(long long a, long long b, long long m) {
    
    long long res = 1;
    while (b > 0) {
        if (b & 1) res = res * a % m;
        a = a * a % m;
        b >>= 1;
    }
    
    return res;
}

// Driver Code for testing
int main() {
    
    cout << binpow(2, 16) << "\n";
    cout << binpow_recursive(2, 16) << "\n";   
    cout << binpow_mod(2, 16, 111) << "\n";
    
    return 0;
}
