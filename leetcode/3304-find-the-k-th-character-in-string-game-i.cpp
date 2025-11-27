#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/find-the-k-th-character-in-string-game-i/description/
class Solution {
private:
	// Approach 1
	// using basic simulation
	// T(n) : O(logn) ; S(n) : O(n)
	char solveKthCharacter(int k) {
		
		string word = "a";
		while (word.size() <= k) {
			string temp = "";
			for (int i = 0; i < word.size(); i++) {
				temp += (word[i] - 'a' + 1) % 26 + 'a';
			}
			word += temp;
		}
		
		return word[k-1];
	}
public:
    char kthCharacter(int k) {
        return solveKthCharacter(k);
    }
};

// Driver Code for testing
int main() {
	
	return 0;
}
