#include<iostream>
#include <vector>
#define ll long long

using namespace std;

/*
	Problem : https://cses.fi/problemset/task/1094
*/ 

ll solve(vector<ll>& arr) {
	
	ll res=0;
	for (int i=1; i<arr.size(); i++) {
		if (arr[i] < arr[i-1]) {
			res += arr[i-1] - arr[i];
			arr[i] += (arr[i-1] - arr[i]);
		}
	}

	return res;
}

int main() {
	
	ll n, a;
	cin >> n;
	vector<ll> arr;
	while (n--) {
		cin >> a;
		arr.push_back(a);
	}
	cout << solve(arr) << "\n";

	return 0;
}