#include <iostream>
#define ll long long

using namespace std;

/*
	Problem : https://cses.fi/problemset/task/1068
*/

void solveSimulation(ll n) {

	if (n == 1) {
		cout << n;
		return;
	}

	cout << n << " ";
	
	if (n % 2 == 0) {
		n /= 2;

	} else {
		n = n * 3 + 1;
	}

	solveSimulation(n);
}

int main() {

	int n;
	cin >> n;
	solveSimulation(n);

	return 0;
}