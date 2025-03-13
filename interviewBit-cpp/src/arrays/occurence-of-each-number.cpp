#include <iostream>
#include <vector>
#include <map>

using namespace std;

// https://www.interviewbit.com/problems/occurence-of-each-number/

// Approach 1
// using ordered-map
// T(n) : O(nlogn) ; S(n) : O(n)
vector<int> occurenceOfEachNumber(vector<int> &A) {

    map<int, int> m;
    for (auto &a : A) m[a]++;

    vector<int> freqArr;
    for (auto it = m.begin(); it != m.end(); it++) {
        freqArr.push_back(it->second);
    }

    return freqArr;
}

// Driver Code for testing
int main() {

    vector<int> arr = {5,5,1,1,1,2,2,3,3,3};
    vector<int> res = occurenceOfEachNumber(arr);
    for (auto &r : res) cout << r << " ";
    cout << "\n";

    return 0;
}
