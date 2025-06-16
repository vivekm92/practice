#include <iostream>
#include <functional>
#include <unordered_set>

using namespace std;

// https://leetcode.com/problems/string-matching-in-an-array/description/
class Solution {
private:
    // Approach 1
    // using brute-force
    // T(n) : O(n2*m2) ; S(n) : O(n)
    vector<string> solveStringMatching(vector<string> &words) {
        
       int n = words.size();
       unordered_set<string> us;
       for (int i = 0; i < n; i++) {
           for (int j = i+1; j < n; j++) {
               string s1 = words[i];
               string s2 = words[j];
               if (s1.size() == s2.size() && s1 == s2) {
                   us.insert(s1);
               } else if (s1.size() > s2.size() && s1.find(s2) != string::npos) {
                    us.insert(s2);
               } else if (s2.size() > s1.size() && s2.find(s1) != string::npos) {
                   us.insert(s1);
               }
           } 
       }
       
       vector<string> res;
       for (auto &s : us) res.push_back(s);

       return res;
    }
    
    // Approach 2
    // using optimised brute-force
    // T(n) : O(m2*n2) ; S(n) : O(1)
    vector<string> solveStringMatching1(vector<string> &words) {
        
        int n = words.size();
        vector<string> res;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (j == i) continue;
                if (words[j].find(words[i]) != string::npos) {
                    res.push_back(words[i]);
                    break;
                }
            }
        }
        
        return res;
    }
    
    // Approach 3
    // using optimised brute-force
    // T(n) : O(m2*n2) ; S(n) : O(1)
    vector<string> solveStringMatching2(vector<string> words) {
        
        function<bool(string&, string&)> isSubstring = [&](string &s1, string &s2) {
            int k = s2.size() - s1.size() + 1;
            for (int i = 0; i < k; i++) {
                if (s2.substr(i, s1.size()) == s1) return true;
            }
            return false;  
        };
        
        int n = words.size();
        vector<string> res;
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                if (j == i) continue;
                if (isSubstring(words[i], words[j])) {
                    res.push_back(words[i]);
                    break;
                }
            }
        }
        
        return res;
    }
    
    // Approach 4
    // using KMP Algorithm
    // T(n) : O(n) ; S(n) : O(1)
    vector<string> solveStringMatching3(vector<string> &words) {
        
    }
public:
    vector<string> stringMatching(vector<string>& words) {
        return solveStringMatching(words);
    }
};

// Driver Code for testing
int main() {
    
    return 0;
}
