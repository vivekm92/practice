#include <iostream>
#include <vector>
#include <stack>

using namespace std;

// https://www.interviewbit.com/problems/evaluate-expression/

// using stacks
// T(n) : O(n) ; S(n) : O(n)
int evaluateExpression(vector<string> &A) {

    size_t n = A.size();
    stack<long long> st;
    for (int i=0; i<n; i++) {
        if (A[i] != "*" && A[i] != "/" && A[i] != "+" && A[i] != "-") {
            long long v1 = stoll(A[i]);
            st.push(v1);
            
            continue;
        }

        long long v1 = 0, v2 = 0;
        if (!st.empty()) {
            v1 = st.top();
            st.pop();
        } else {
            return -1;
        }

        if (!st.empty()) {
            v2 = st.top();
            st.pop();
        } else {
            return -1;
        }

        long long res = 0;
        if (A[i] == "*") {
            res = v2 * v1;
        } else if (A[i] == "/") {
            res = v2 / v1;
        } else if (A[i] == "+") {
            res = v2 + v1;
        } else if (A[i] == "-") {
            res = v2 - v1;
        }

        st.push(res);
    }

    return st.top();
}

// Driver Code for testing
int main () {

    vector<string> A = {"2", "1", "+", "3", "*"};
    cout << evaluateExpression(A) << "\n";

    return 0;
}
