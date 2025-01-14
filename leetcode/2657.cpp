#include <iostream>
#include <vector>
#include <unordered_set>

using namespace std;

// https://leetcode.com/problems/find-the-prefix-common-array-of-two-arrays/description/
class Solution {
    private:
        // using three for loops : brute-force
        // T(n) : O(n3) ; S(n) : O(1)
        vector<int> solvefindThePrefixCommonArray1(vector<int>& A, vector<int>& B) {

            int n = A.size();
            vector<int> arr(n, 0);
            for (int i = 0; i < n; i++) {
                int count = 0;
                for (int j = 0; j <= i; j++) {
                    for (int k = 0; k <= i; k++) {
                        if (A[j] == B[k]) {
                            count++;
                            break;
                        }
                    }
                }
                arr[i] = count;
            }
            return arr;
        }

        // using two for loops and arr map
        // T(n) : O(n2) ; S(n) : O(n)
        vector<int> solvefindThePrefixCommonArray(vector<int>& A, vector<int>& B) {
            int n = A.size();
            vector<int> arr(n+1, 0), brr(n+1, 0), crr(n, 0);
            for (int i = 0; i < n; i++) {
                arr[A[i]] += 1;
                brr[B[i]] += 1;
                for (int j = 0; j < n+1; j++) {
                    if (arr[j] > 0 && brr[j] > 0 && arr[j] == brr[j]) {
                        crr[i] += 1;
                    }
                }
            }
            return crr;
        }

        // using two hash-set
        // T(n) : O(n2) ; S(n) : O(n)
        vector<int> solvefindThePrefixCommonArray2(vector<int>& A, vector<int>& B) {
            int n = A.size();
            unordered_set<int> us1, us2;
            vector<int> res(n, 0);
            for (int i = 0; i < n; i++) {
                us1.insert(A[i]);
                us2.insert(B[i]);
                int count = 0;
                for (auto it = us1.begin(); it != us1.end(); it++) {
                    if (us2.find(*it) != us2.end()) {
                        count++;
                    }
                }
                res[i] = count;
            }

            return res;
        }

        // using freq arr : single paas appproach
        // T(n) : O(n) ; S(n) : O(n)
        vector<int> solvefindThePrefixCommonArray3(vector<int>& A, vector<int>& B) {
            int n = A.size(), count = 0;
            vector<int> freq(n + 1, 0), res(n, 0);
            for (int i=0; i<n; i++) {
                freq[A[i]]++;
                if (freq[A[i]] == 2) count++;
                freq[B[i]]++;
                if (freq[B[i]] == 2) count++;
                res[i] = count;
            }
            return res;
        }
    public:
        vector<int> findThePrefixCommonArray(vector<int>& A, vector<int>& B) {
            return solvefindThePrefixCommonArray3(A, B);
        }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    vector<int> A = {1, 3 ,2, 4};
    vector<int> B = {3, 1, 2, 4};
    vector<int> res = sol->findThePrefixCommonArray(A, B);
    for (auto &a : res) {
        cout << a << " ";
    }
    cout << "\n";

    return 0;
}
