#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/sort-vowels-in-a-string/description/
class Solution {
private:
	bool isVowel(char ch) {
		return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' ||
			ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U';
	}
	
	// Approach 1
	// using sorting
	// T(n) : O(nlogn) ; S(n) : O(n)
	string solveSortVowels(string s) {
		
		vector<char> temp;
		for (auto &ch : s) {
			if (isVowel(ch)) {
				temp.push_back(ch);
			}
		}
		
		sort(temp.begin(), temp.end());
		int n = s.size(), k = 0;
		for (auto &ch : s) {
			if (isVowel(ch)) {
				ch = temp[k++];
			}
		}
		
		return s;
	}

	// Approach 2
	// using counting sort
	// T(n) : O(n) ; S(n) : O(n)
	string solveSortVowels1(string s) {
		
		vector<int> vowelCount(10, 0);
		unordered_map<char, int> lookup;
		lookup['A'] = 0;
		lookup['E'] = 1;
		lookup['I'] = 2;
		lookup['O'] = 3;
		lookup['U'] = 4;
		lookup['a'] = 5;
		lookup['e'] = 6;
		lookup['i'] = 7;
		lookup['o'] = 8;
		lookup['u'] = 9;
		
		unordered_map<int, char> revLookup;
		revLookup[0] = 'A';
		revLookup[1] = 'E';
		revLookup[2] = 'I';
		revLookup[3] = 'O';
		revLookup[4] = 'U';
		revLookup[5] = 'a';
		revLookup[6] = 'e';
		revLookup[7] = 'i';
		revLookup[8] = 'o';
		revLookup[9] = 'u';
		for (auto &ch : s) {
			if (lookup.find(ch) != lookup.end()) {
				vowelCount[lookup[ch]]++;
			}
		}
		
		int idx = 0;
		for (auto &ch : s) {
			if (um.find(ch) != um.end()) {	
				while (idx < 10 && vowelCount[idx] == 0) {
					idx++;
				}
				if (idx >= 10) break;
				ch = revLookup[idx];
				vowelCount[idx]--;
			}
		}
		
		return s;
	}
public:
    string sortVowels(string s) {
        return solveSortVowels(s);
    }
};

// Driver Code for testing.
int main() {
	
}
