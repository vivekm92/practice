#include <iostream>
#include <unordered_set>
#includ <set>

using namespace std;

// https://leetcode.com/problems/design-a-number-container-system/description/
class NumberContainers {
private:
    unordered_map<int, set<int> > numberToIndices;
    unordered_map<int, int> indicesToNumber;
public:
    NumberContainers() {
        
    }
    
    void change(int index, int number) {
        if (indicesToNumber.find(index) != indicesToNumber.end()) {
            int prevNum = indicesToNumber[index];
            numberToIndices[prevNum].erase(index);
            if (numberToIndices[prevNum].empty()) numberToIndices.erase(prevNum);
        }
        numberToIndices[number].insert(index);
        indicesToNumber[index] = number;
    }
    
    int find(int number) {
        if (numberToIndices.find(number) != numberToIndices.end()) {
            return *numberToIndices[number].begin();
        }
        return -1;
    }
};


class NumberContainers1 {
private:
    
public:
    NumberContainers1() {
        
    }
    
    void change(int index, int number) {
        
    }
    
    int find(int number) {
        
    }
};

/**
 * Your NumberContainers object will be instantiated and called as such:
 * NumberContainers* obj = new NumberContainers();
 * obj->change(index,number);
 * int param_2 = obj->find(number);
 */

// Driver Code for testing
int main() {
    
    return 0;
}
