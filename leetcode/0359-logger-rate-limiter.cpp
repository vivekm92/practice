#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/logger-rate-limiter/description/

// Approach 1
// using hash-table
// T(n) : O(1) ; S(n) ; O(n)
class Logger {
private:
    unordered_map<string, int> um;
public:
    Logger() {
        
    }
    
    bool shouldPrintMessage(int timestamp, string message) {
        
        if (this->um.find(message) != this->um.end() &&
                this->um[message] > timestamp) {
            return false;
        }
        
        this->um[message] = timestamp + 10;
        return true;
    }
};

/**
 * Your Logger object will be instantiated and called as such:
 * Logger* obj = new Logger();
 * bool param_1 = obj->shouldPrintMessage(timestamp,message);
 */

// Driver Code for testing
int main() {
 
    return 0;
}
