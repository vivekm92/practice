#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

class Solution {
private:
    // using hash-map and sliding-window
    // T(n) : O(n) ; S(n) : O(1)
    int solveNumberOfSubstrings(string s) {

        unordered_map<char, int> um;
        int n = s.size(), start = 0, validSubstringsCount = 0;
        for (int end = 0; end < n; end++) {
            um[s[end]]++;
            while (um.size() == 3) {
                validSubstringsCount += (n - end);
                um[s[start]]--;
                if (um[s[start]] == 0) um.erase(s[start]);
                start++;
            }
        }

        return validSubstringsCount;
    }

    // using arrays as hashMap and sliding-window
    // T(n) : O(n) ; S(n) : O(1)
    int solveNumberOfSubstrings1(string s) {

        vector<int> arr(3, 0);
        int n = s.size(), start = 0, validSubstringsCount = 0;
        for (int end = 0; end < n; end++) {
            arr[s[end] - 'a']++;
            while (arr[0] > 0 && arr[1] > 0 && arr[2] > 0) {
                validSubstringsCount += (n - end);
                arr[s[start] - 'a']--;
                start++;
            }
        }

        return validSubstringsCount;
    }

    // using last pos tracking
    // T(n) : O(n) ; S(n) : O(1)
    int solveNumberOfSubstrings2(string s) {

        int n = s.size(), validSubstringsCount = 0;
        vector<int> arr(3, -1);
        for (int i = 0; i < n; i++) {
            arr[s[i] - 'a'] = i;
            validSubstringsCount += 1 + min({arr[0], arr[1], arr[2]});
        }

        return validSubstringsCount;
    }
public:
    int numberOfSubstrings(string s) {
        return solveNumberOfSubstrings2(s);
    }
};


// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    cout << sol->numberOfSubstrings("abcabc") << "\n";
    cout << sol->numberOfSubstrings("aaacb") << "\n";
    cout << sol->numberOfSubstrings("abc") << "\n";

    return 0;
}
