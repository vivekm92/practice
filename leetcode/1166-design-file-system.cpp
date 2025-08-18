#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/design-file-system/description/
class FileSystem {
private:
    struct node {
        unordered_map<string, node* > children;
        int val;
    };

    node* root;

    // T(n) : O(n) ; S(n) : O(n)
    vector<string> split(string path) {

        vector<string> res;
        stringstream ss(path);
        string segment;
        while (getline(ss, segment, '/')) {
            if (!segment.empty()) res.push_back(segment);
        }
        return res;
    }
public:
    FileSystem() {
        this->root = new node();
        this->root->val = -1;
    }

    // T(n) : O(n) ; S(n) : O(n)
    bool createPath(string path, int value) {

        vector<string> segments = split(path);
        node *curr = this->root;
        for (int i = 0; i < segments.size() - 1; i++) {
            if (!curr->children.count(segments[i])) return false;
            curr = curr->children[segments[i]];
        }
        string last = segments.back();
        if (curr->children.count(last)) return false;
        curr->children[last] = new node();
        curr->children[last]->val = value;
        return true;
    }

    // T(n) : O(n) ; S(n) : O(1)
    int get(string path) {

        vector<string> segments = split(path);
        node *curr = this->root;
        for (const string &segment : segments) {
            if (!curr->children.count(segment)) return -1;
            curr = curr->children[segment];
        }
        return curr->val;
    }
};

/**
 * Your FileSystem object will be instantiated and called as such:
 * FileSystem* obj = new FileSystem();
 * bool param_1 = obj->createPath(path,value);
 * int param_2 = obj->get(path);
 */

// Driver Code for testing.
int main() {
    FileSystem *f = new FileSystem();
    cout << f->createPath("/leet/code", 1);
    return 0;
}
