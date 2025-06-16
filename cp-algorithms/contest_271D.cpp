#include <iostream>
#include <unordered_set>

using namespace std;

long long solve(string s, string ch, int k) {

    int n = s.size();
    const int p = 31;
    const int m = 1e9+9;
    vector<long long> p_pow(n);
    p_pow[0] = 1;
    for (int i = 1; i < n; i++) 
        p_pow[i] = (p_pow[i-1] * p) % m;
        
    vector<long long> h(n+1, 0);
    for (int i = 0; i < n; i++) 
        h[i+1] = (h[i] + (s[i] - 'a' + 1) * p_pow[i]) % m;
    
    long long count = 0;
    for (int i = 
        
    return count;
}

int main() {
    
    string s, ch;
    int k;
    cin >> s >> ch >> k;
    long long res = solve(s, ch, k);
    cout << res << "\n";
 
    return 0;
}
