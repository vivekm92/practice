#include <iostream>

#define ll  long long int
#define vi  vector<int>
#define vll vector<long long int>

using namespace std;

int reverseSort(vi &A) {
	
	int cost = 0, n=A.size();
	for (int i=0; i<n-1; i++) {
		int m = *min(A.begin() + i, A.end());
		int idx = *find(A.begin()+i, A.end(), m);
		reverse(A.begin() + i, A.begin() + idx + 1);
		cost += (idx - i + 1);
	}

	return cost;
}

int main() {
	
	ios_base::sync_with_stdio(false);
	cin.tie(nullptr);
	
	int T;
	cin >> T;
		
	while (T--) {
		
		int N;
		cin >> N;
		
		vi A;
		for (int i=0; i<N; i++) {
			int num;
			cin >> num;
			A.push_back(num);
		}
		
		
		cout << "Case #1: " << reverseSort(A) << "\n";
		
	}
	
	return 0;
}
