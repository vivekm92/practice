#include <iostream>

using namespace std;

/*
	Problem : https://cses.fi/problemset/task/1069
*/ 

int main() {

	string s;
	cin >> s;

	int res=1, temp=1;
	for (int i=1; i<s.size(); i++) {
		if (s[i] == s[i-1]) {
			temp++;
		} else {
			temp = 1;
		}

		res = max(res, temp);
	}

	cout << res << "\n";
	return 0;
}