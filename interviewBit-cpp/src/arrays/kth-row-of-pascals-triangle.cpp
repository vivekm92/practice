#include <iostream>
#include <vector>
#include <cassert>

using namespace std;

// Problem : https://www.interviewbit.com/problems/kth-row-of-pascals-triangle/

// Approach 1
// using recursive approach
// T(n) : O(n2) ; S(n) : O(n)
vector<int> kthRowOfPascalsTriangle1(int A) {

    if (A == 0) return {1};
    else if (A == 1) return {1,1};
    vector<int> row = kthRowOfPascalsTriangle1(A-1);
    for (int i = row.size()-1; i >= 1; i--) {
        row[i] += row[i-1];
    }
    row.push_back(1);
    return row;
}

// Approach 2
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

// Approach 3
// using binomial-coefficient
// T(n) : O(n) : S(n) : O(1)
vector<int> kthRowOfPascalsTriangle3(int A) {

    vector<int> res;
    int t = 1;
    for (int i = 1; i <= A + 1; i++) {
        res.push_back(t);
        t = t * (A + 1 - i) / i;
    }

    return res;
}

void test_correctness(vector<int>& tc, vector<int>& res) {

    cout << "got  : ";
    for (auto &t : tc) cout << t << " ";
    cout << "\n";
    cout << "want : ";
    for (auto &r : res) cout << r << " ";
    cout << "\n";
    assert(tc.size() == res.size());
    for (int i = 0; i < tc.size(); i++) {
        assert(tc[i] == res[i]);
    }
}

// Driver Code for testing
int main() {

    vector<int> arr1 = kthRowOfPascalsTriangle1(4);
    vector<int> brr1 = kthRowOfPascalsTriangle2(4);
    vector<int> crr1 = kthRowOfPascalsTriangle3(4);
    vector<int> want1 = {1,4,6,4,1};
    test_correctness(arr1, want1);
    test_correctness(brr1, want1);
    test_correctness(crr1, want1);

    vector<int> arr2 = kthRowOfPascalsTriangle1(5);
    vector<int> brr2 = kthRowOfPascalsTriangle2(5);
    vector<int> crr2 = kthRowOfPascalsTriangle3(5);
    vector<int> want2 = {1,5,10,10,5,1};
    test_correctness(arr2, want2);
    test_correctness(brr2, want2);
    test_correctness(crr2, want2);

    vector<int> arr3 = kthRowOfPascalsTriangle1(6);
    vector<int> brr3 = kthRowOfPascalsTriangle2(6);
    vector<int> crr3 = kthRowOfPascalsTriangle3(6);
    vector<int> want3 = {1,6,15,20,15,6,1};
    test_correctness(arr3, want3);
    test_correctness(brr3, want3);
    test_correctness(crr3, want3);

    return 0;
}

/*
 * k = 0, o/p = {1}
 * k = 1, o/p = {1,1}
 * k = 2, o/p = {1,2,1}
 * k = 3, o/p = {1,3,3,1}
 * k = 4, o/p = {1,4,6,4,1}
 * k = 5, o/p = {1,5,10,10,5,1}
 * k = 6, o/p = {1,6,15,20,15,6,1}
 * */
