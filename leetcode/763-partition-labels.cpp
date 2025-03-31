#include <iostream>
#include <vector>
#include <unordered_set>

using namespace std;

// https://leetcode.com/problems/partition-labels/description/
class Solution {
private:
    // Approach 1
    // using hash-set
    // T(n) : O(n) ; S(n) : O(n)
    vector<int> solvePartitionLabels(string s) {

        vector<int> lookup(26, 0);
        for (auto &ch : s) lookup[ch - 'a']++;

        int count = 0;
        vector<int> res;
        unordered_set<char> us;
        for (auto &ch : s) {
            lookup[ch - 'a']--;
            count++;
            us.insert(ch);
            if (lookup[ch - 'a'] == 0) {
                bool flag = true;
                for (auto it = us.begin(); it != us.end(); it++) {
                    if (lookup[*it - 'a'] != 0) {
                        flag = false;
                        break;
                    }
                }

                if (flag) {
                    res.push_back(count);
                    count = 0;
                    us.clear();
                }
            }
        }

        return res;
    }

    // Approach 2
    // using two-pointers
    // T(n) : O(n) ; S(n) : O(k) where k = 26
    vector<int> solvePartitionLabels1(string s) {

        int n = s.size();
        vector<int> lastOccurence(26, 0);
        for (int i = 0; i < n; i++) lastOccurence[s[i] - 'a'] = i;

        vector<int> res;
        int partitionStart = 0, partitionEnd = 0;
        for (int i = 0; i < n; i++) {
            partitionEnd = max(partitionEnd, lastOccurence[s[i] - 'a']);
            if (i == partitionEnd) {
                res.push_back(partitionEnd - partitionStart + 1);
                partitionStart = partitionEnd + 1;
            }
        }
        return res;
    }

    // Approach 3
    // using merge operations
    // T(n) : O(n) ; S(n) : O(k) where k = 26
    vector<int> solvePartitionLabels2(string s) {

        int n = s.size();
        vector<int> lastOccurence(26, 0), firstOccurence(26, -1);
        for (int i = 0; i < n; i++) {
            lastOccurence[s[i] - 'a'] = i;
        }

        vector<int> res;
        int partitionStart = 0, partitionEnd = 0;
        for (int i = 0; i < n; i++) {
            if (firstOccurence[s[i] - 'a'] == -1) {
                firstOccurence[s[i] - 'a'] = i;
            }

            if (partitionEnd < firstOccurence[s[i] - 'a']) {
                res.push_back(partitionEnd - partitionStart + 1);
                partitionEnd = i;
                partitionStart = i;
            }

            partitionEnd = max(partitionEnd, lastOccurence[s[i] - 'a']);
        }

        if (partitionEnd - partitionStart + 1 > 0) {
            res.push_back(partitionEnd - partitionStart + 1);
        }

        return res;
    }
public:
    vector<int> partitionLabels(string s) {
        return solvePartitionLabels1(s);
    }
};

// Driver Code for testing
int main() {
    return 0;
}
