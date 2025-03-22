#include <iostream>
#include <vector>
#include <cassert>

using namespace std;


void rotateMatrix(vector<vector<int> >& A) {

    int n = A.size(), m = A[0].size();

}

// Driver Code for testing
int main() {

    vector<vector<int> > arr {{1,2}, {3,4}};
    vector<vector<int> > res {{3,1}, {4,2}};
    rotateMatrix(arr);
    for (int i = 0; i < arr.size(); i++) {
        for (int j = 0; j < arr[0].size(); j++) {
            assert(arr[i][j] == res[i][j]);
        }
    }

    vector<vector<int> > arr1 {{1}};
    vector<vector<int> > res1 {{1}};
    rotateMatrix(arr1);
    for (int i = 0; i < arr1.size(); i++) {
        for (int j = 0; j < arr1[0].size(); j++) {
            assert(arr1[i][j] == res1[i][j]);
        }
    }



    return 0;
}
