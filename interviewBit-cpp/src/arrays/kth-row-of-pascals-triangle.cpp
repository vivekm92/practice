#include <iostream>
#include <vector>

using namespace std;

// using recursive approach
// T(n) : O(n2) ; S(n) : O(n)
vector<int> kthRowOfPascalsTriangle(int A) {
    
    if (A == 0) return {1};
    else if (A == 1) return {1,1};
    
    vector<int> prevRow = kthRowOfPascalsTriangle(A-1);
    vector<int> temp;
    for (int i = 0; i < prevRow.size(); i++) {
        if (i == 0) temp.push_back(1);
        else temp.push_back(prevRow[i-1] + prevRow[i]);
    }
    temp.push_back(1);
    return temp;
}

// using binomial-coefficient
// T(n) : O(n) : S(n) : O(1)
vector<int> kthRowOfPascalsTriangle1(int A) {
    
    vector<int> res;
    int t = 1;
    for (int i = 1; i <= A + 1; i++) {
        res.push_back(t);
        t = t * (A + 1 - i) / i;
    }
    
    return res;
}

// using iterative approach
// T(n) : O(n2) ; S(n) : O(n)
vector<int> kthRowOfPascalsTriangle2(int A) {
    
    vector<int> res(A+1, 0);
    res[0] = 1;
    for (int i = 1; i <= A; i++) {
        for (int j = i; j >= 1; j--) {
            res[j] += res[j-1];
        }
    }
    
    return res;
}


// Driver Code for testing
int main() {
    
    vector<int> arr = kthRowOfPascalsTriangle(4);
    for (auto &a : arr) cout << a << " ";
    cout << "\n";
    
    vector<int> arr1 = kthRowOfPascalsTriangle(5);
    for (auto &a : arr1) cout << a << " ";
    cout << "\n";
    
    vector<int> arr2 = kthRowOfPascalsTriangle(6);
    for (auto &a : arr2) cout << a << " ";
    cout << "\n";

    vector<int> arr3 = kthRowOfPascalsTriangle1(4);
    for (auto &a : arr3) cout << a << " ";
    cout << "\n";
    
    vector<int> arr4 = kthRowOfPascalsTriangle1(5);
    for (auto &a : arr4) cout << a << " ";
    cout << "\n";
    
    vector<int> arr5 = kthRowOfPascalsTriangle1(6);
    for (auto &a : arr5) cout << a << " ";
    cout << "\n";
    
    vector<int> arr6 = kthRowOfPascalsTriangle2(4);
    for (auto &a : arr6) cout << a << " ";
    cout << "\n";
    
    vector<int> arr7 = kthRowOfPascalsTriangle2(5);
    for (auto &a : arr7) cout << a << " ";
    cout << "\n";
    
    vector<int> arr8 = kthRowOfPascalsTriangle2(6);
    for (auto &a : arr8) cout << a << " ";
    cout << "\n";
    return 0;
}

/*
 * k = 0, o/p = {1}
 * k = 1, o/p = {1,1}
 * k = 2, o/p = {1,2,1}
 * k = 3, o/p = {1,3,3,1}
 * k = 4, o/p = {1,4,6,4,1}
 * */
