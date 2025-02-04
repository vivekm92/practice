#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

// https://www.interviewbit.com/problems/distinct-numbers/
// T(n) : O(n) ; S(n) : O()
int solveDistinctNumbers(vector<int>& A, int B) {
    unordered_map<int, int> um;
    for (int &a : A) um[a]++;
    if (um.size() < B) return 0;
    vector<int> arr;
    for (auto &it : um) {
        arr.push_back(it.second);
    }
    sort(arr.begin(), arr.end());
    int count = 0;
    for (int i = 0; i < arr.size() - B; i++) {
        count += arr[i];
    }
    return count;
}


// Driver Code for testing
int main() {
    vector<int> arr {1, 2, 3};
    cout << solveDistinctNumbers(arr, 2) << "\n";
    vector<int> arr1 {1, 2, 2};
    cout << solveDistinctNumbers(arr1, 2) << "\n";
    return 0;
}
