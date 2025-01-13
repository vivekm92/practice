#include <iostream>

struct TrieNodeFromArray {
	std::vector<TrieNodeFromArray* > characters;
	bool flag;
	TrieNodeFromArray() {
        characters.resize(26);
        flag = false;
    }
};

class TrieFromArray {
	private:
		TrieNodeFromArray *root;
	public:
		TrieFromArray () {
			this->root = new TrieNodeFromArray();
		}
		
		void insert(std::string word) {
			
			TrieNodeFromArray *node = this->root;
			for (auto &c : word) {
				if (node->characters[c-'a'] == nullptr) {
					TrieNodeFromArray *t = new TrieNodeFromArray();
					node->characters[c-'a'] = t;
				}
				node = node->characters[c-'a'];
			}
			node->flag = true;
		}
		
		bool search(std::string word, bool checkForPrefix = false) {
			
			TrieNodeFromArray *node = root;
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

struct TrieNodeFromMap {
	std::unordered_map<char, TrieNodeFromMap* > m;
	bool flag;
	
	TrieNodeFromMap() {
		flag = false;
	}
};

class TrieFromMap {
	private:
		TrieNodeFromMap* root;
		
	public:
		TrieFromMap() {
			this->root = new TrieNodeFromMap();
		}
		
		void insert(std::string word) {
			
			TrieNodeFromMap* node = this->root;
			for (auto &c : word) {
				if (node->m.find(c) == node->m.end()) {
					node->m[c] = new TrieNodeFromMap();
				}
				node = node->m[c];
			}
			node->flag = true;
		}
		
		bool search(std::string word, bool isPrefix=false) {
			
			TrieNodeFromMap* node = this->root;
			for (auto &c : word) {
				if (node->m.find(c) == node->m.end()) {
					return false;
				}
				node = node->m[c];
			}
			
			return isPrefix ? true : node->flag;
		}
		
		bool startsWith(std::string prefix) {
			return this->search(prefix, true);
		}
};
