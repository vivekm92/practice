#include<iostream>
#define ll long long

using namespace std;

/*
	Problem : https://cses.fi/problemset/task/1083
*/ 

int main() {

	ll n, t;
	cin >> n;
	ll sum1=0, sum2=0;
	for (int i=0; i<n-1; i++) {
		cin >> t;
		sum2 += t;
	}

	sum1 = n * (n + 1) / 2;

	cout << sum1 - sum2 << "\n";

	return 0;
}