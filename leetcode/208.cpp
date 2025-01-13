#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/implement-trie-prefix-tree/description/
struct TrieNode {
	std::vector<TrieNode* > characters;
	bool flag;

    TrieNode() {
        characters.resize(26);
        flag = false;
    }
};

class Trie {
	private:
		TrieNode *root;
	public:
		Trie () {
			this->root = new TrieNode();
		}
		
		void insert(std::string word) {
			
			TrieNode *node = this->root;
			for (auto &c : word) {
				if (node->characters[c-'a'] == nullptr) {
					TrieNode *t = new TrieNode();
					node->characters[c-'a'] = t;
				}
				node = node->characters[c-'a'];
			}
			node->flag = true;
		}
		
		bool search(std::string word, bool checkForPrefix = false) {
			
			TrieNode *node = root;
			for (auto &c : word) {
				if (node->characters[c-'a'] == nullptr) {
					return false;
				}
				node = node->characters[c-'a'];
			}
			
			return checkForPrefix ? true : node->flag;
		}
		
		bool startsWith(std::string prefix) {
			return this->search(prefix, true);
		}
};

// Driver code for testing
int main() {
	
	return 0;
}
