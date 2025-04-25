#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/number-of-senior-citizens/description/
class Solution {
private:
    // Approach 1
    // using basic simulation
    // T(n) : O(n) ; S(n) : O(1)
    int solveCountSeniors(vector<string> &details) {

        int count = 0;
        for (auto &str : details) {
            string s = str.substr(11, 2);
            int age = stoi(s);
            if (age > 60) count++;
        }

        return count;
    }
public:
    int countSeniors(vector<string>& details) {
        return solveCountSeniors(details);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
