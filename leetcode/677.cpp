#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/map-sum-pairs/description/
struct TrieNode {
    vector<TrieNode* > arr;
    bool flag;
    TrieNode() {
        arr.resize(26);
        flag = false;
    }
};

class MapSum {
private:
    unordered_map<string, int> um;
    TrieNode* root;
public:
    MapSum() {
        this->root = new TrieNode();
    }

    void insert(string key, int val) {
        this->um[key] = val;
        TrieNode* node = this->root;
        for (auto &c : key) {
            if (node->arr[c-'a'] == nullptr) {
                node->arr[c-'a'] = new TrieNode();
            }
            node = node->arr[c-'a'];
        }
        node->flag = true;
    }

    bool isPrefix(string word, string prefix) {

        if (word.size() < prefix.size()) return false;
        TrieNode *node = this->root;
        for (int i=0; i<prefix.size(); i++) {
            if (prefix[i] != word[i] || node->arr[prefix[i]-'a'] == nullptr) {
                return false;
            }
            node = node->arr[prefix[i]-'a'];
        }
        return true;
    }

    int sum(string prefix) {
        int total = 0;
        for (auto it = this->um.begin(); it != this->um.end(); it++) {
            if (this->isPrefix(it->first, prefix)) {
                total += it->second;
            }
        }
        return total;
    }
};

/**
 * Your MapSum object will be instantiated and called as such:
 * MapSum* obj = new MapSum();
 * obj->insert(key,val);
 * int param_2 = obj->sum(prefix);
 */

 // Driver code for testing
 int main() {

     return 0;
 }
