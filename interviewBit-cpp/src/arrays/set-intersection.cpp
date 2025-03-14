#include <iostream>
#include <vector>
#include <cassert>

using namespace std;

// https://www.interviewbit.com/problems/set-intersection/

// T(n) : O(nlogn) ; S(n) : O(n)
int setIntersection1(vector<vector<int> > &A) {

    sort(A.begin(), A.end(), [](vector<int>& v1, vector<int> &v2) {
        assert(v1.size() == 2);
        assert(v2.size() == 2);
        if (v1[1] == v2[1]) return v1[0] < v2[0];
        return v1[1] < v2[1];
    });

    int n = A.size();
    vector<int> resultantSet;
    resultantSet.push_back(A[0][1]-1);
    resultantSet.push_back(A[0][1]);
    for (int i = 1; i < n; i++) {
        int start = A[i][0], end = A[i][1];
        int sz = resultantSet.size();
        int last = resultantSet[sz-1], secondLast = resultantSet[sz-2];
        if (start > last) {
            resultantSet.push_back(end-1);
            resultantSet.push_back(end);
        } else if (start == last) {
            resultantSet.push_back(end);
        } else if (start > secondLast) {
            resultantSet.push_back(end);
        }
    }

    return resultantSet.size();
}

// T(n) : O(nlogn) ; S(n) : O(1)
int setIntersection2(vector<vector<int> > &A) {

    sort(A.begin(), A.end(), [](vector<int>& v1, vector<int> &v2) {
        assert(v1.size() == 2);
        assert(v2.size() == 2);
        if (v1[1] == v2[1]) return v1[0] < v2[0];
        return v1[1] < v2[1];
    });

    int n = A.size(), sz = 2;
    int sl = A[0][1] - 1, l = A[0][1];
    for (int i = 1; i < n; i++) {
        int start = A[i][0], end = A[i][1];
        if (start > l) {
            sl = end - 1;
            l = end;
            sz += 2;
        } else if (start == l) {
            sl = l;
            l = end;
            sz += 1;
        } else if (start > sl) {
            sl = l;
            l = end;
            sz += 1;
        }
    }

    return sz;
}

// T(n) : O(nlogn) ; S(n) : O(1)
int setIntersection3(vector<vector<int> > &A) {

    sort(A.begin(), A.end(), [](vector<int>& v1, vector<int> &v2) {
        assert(v1.size() == 2);
        assert(v2.size() == 2);
        if (v1[1] == v2[1]) return v1[0] > v2[0];
        return v1[1] < v2[1];
    });

    int n = A.size(), sz = 2;
    int sl = A[0][1] - 1, l = A[0][1];
    for (int i = 1; i < n; i++) {
        int start = A[i][0], end = A[i][1];
        if (start > sl) {
            if (start > l) {
                l = end;
                sl = end - 1;
                sz += 2;
            } else {
                sl = l;
                l = end;
                sz += 1;
            }
        }
    }

    return sz;
}

// Driver Code for testing
int main() {

    vector<vector<int> > arr = {{1,4}, {1,3}, {3,5}, {2,5}};
    cout << setIntersection1(arr) << "\n";
    cout << setIntersection2(arr) << "\n";
    cout << setIntersection3(arr) << "\n";

    vector<vector<int> > arr1 = {{1,2}, {3,4}, {2,3}, {4,5}};
    cout << setIntersection1(arr1) << "\n";
    cout << setIntersection2(arr1) << "\n";
    cout << setIntersection3(arr1) << "\n";

    vector<vector<int> > arr2 = {{3,9}, {3,5}, {19,20}, {3,6}, {9,16}, {10,14}, {9,17}, {1,3}, {4,16}};
    cout << setIntersection1(arr2) << "\n";
    cout << setIntersection2(arr2) << "\n";
    cout << setIntersection3(arr2) << "\n";

    return 0;
}
