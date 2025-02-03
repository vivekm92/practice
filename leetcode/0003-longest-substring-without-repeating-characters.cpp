#include <iostream>
#include <unordered_set>
#include <unordered_map>

using namespace std;

// https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
class Solution {
private:
    // using hash-map
    // T(n) : O(n) ; S(n) : O(n)
    int solvelengthOfLongestSubstring1(string s) {
        int n=s.size();
        int left = 0, right = 0, maxLen = 0;
        unordered_map<char, int> um;
        while (right < n) {
            char r = s[right];
            um[r]++;
            while (um[r] > 1) {
                char l = s[left];
                um[l]--;
                left++;
            }
            maxLen = max(maxLen, right - left + 1);
            right++;
        }
        return maxLen;
    }
    
    // using hash-map
    // T(n) : O(n) ; S(n) : O(n)
    int solvelengthOfLongestSubstring2(string s) {
        int n = s.size(), maxLen = 0;
        unordered_map<char, int> um;
        for (int i=0, j=0; j<n; j++) {
            i = (um[s[j]] > 0) ? max(um[s[j]], i) : i;
            maxLen = max(maxLen, j-i+1);
            um[s[j]] = j+1;
        }
        return maxLen;
        
    }
    
    // using hash-set
    // T(n) : O(n) ; S(n) : O(n)
    int solvelengthOfLongestSubstring(string s) {
        int left=0, right=0, n=s.size(), len=0, maxLen=0;
        unordered_set<char> us;
        while (right < n) {
            while (right < n && us.find(s[right]) == us.end()) {
                us.insert(s[right]);
                right++;
            }
            us.erase(s[left]);
            len = right - left;
            left++;
            maxLen = max(maxLen, len);
        }
        return maxLen;
    }
public:
    int lengthOfLongestSubstring(string s) {
        return solvelengthOfLongestSubstring2(s);
    }
};

// Driver Code for testing
int main() {
    string s = "abacdabc";
    Solution *sol = new Solution();
    cout << sol->lengthOfLongestSubstring(s) << "\n";
    return 0;
}
