#include <iostream>
#include <vector>

using namespace std;

// https://www.interviewbit.com/problems/product-of-all/

// Approach 1
// T(n) : O(n) ; S(n) : O(n)
vector<int> productOfAll(vector<int> &A) {

    int n = A.size();
    if (n == 0) return A;

    const int MOD = 1000000007;
    vector<int> productArr(n, 0);
    long long prod = 1;
    for (int i = n-1; i >= 0; i--) {
        prod = ((prod % MOD) * (A[i] % MOD)) % MOD;
        productArr[i] = prod;
    }

    prod = 1;
    for (int i = 0; i < n; i++) {
        if (i < n - 1) {
            productArr[i] = ((prod % MOD) * (productArr[i+1] % MOD)) % MOD;
        } else {
            productArr[i] = prod;
        }
        prod = ((prod % MOD) * (A[i] % MOD)) % MOD;
    }

    return productArr;
}

// using little fermat's theorem, since m is prime
long long invModulo(long long a, long long m, long long k) {
    a = a % m;
    long long res = 1;
    while (k > 0) {
        if (k & 1) {
            res = (res * a) % m;
        }
        a = a * a % m;
        k >>= 1;
    }
    return res;
}

// Approach 2
// using inverse modulo exp
// T(n) : O(nlogn) ; S(n) : O(1)
vector<int> productOfAll1(vector<int> &A) {

    const int MOD = 1000000007;
    long long product = 1;
    for (auto &a : A) {
        product = ((product % MOD) * (a % MOD) % MOD);
    }

    int n = A.size();
    vector<int> productArr(n, 0);
    for (int i = 0; i < n; i++) {
        productArr[i] = ((product % MOD) * (invModulo(A[i], MOD, MOD-2) % MOD)) % MOD;
    }

    return productArr;
}

// Driver Code for testing
int main() {

    vector<int> arr = {1,2,3,4};
    vector<int> res = productOfAll(arr);
    for (auto &r : res) cout << r << " ";
    cout << "\n";

    vector<int> arr1 = {9,9,9};
    vector<int> res1 = productOfAll(arr1);
    for (auto &r : res1) cout << r << " ";
    cout << "\n";

    vector<int> arr2 = {1,2,3,4};
    vector<int> res2 = productOfAll1(arr2);
    for (auto &r : res2) cout << r << " ";
    cout << "\n";

    vector<int> arr3 = {9,9,9};
    vector<int> res3 = productOfAll1(arr3);
    for (auto &r : res3) cout << r << " ";
    cout << "\n";

    return 0;
}
