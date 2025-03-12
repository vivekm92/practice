#include <iostream>

using namespace std;

// https://www.interviewbit.com/problems/pythagorean-triplets/

// T(n) : O(n2logn) ; S(n) : O(1)
int pythagoreanTriplets(int A) {

    int count = 0;
    for (int i = 1; i <= A; i++) {
        for (int j = i+1; j <= A; j++) {
            int k = i*i + j*j;
            int x = sqrt(k);
            if (x*x == k && x <= A) {
                count++;
            }
        }
    }

    return count;
}

int isPerfectSquare(int A) {

    if (A == 1) return 1;

    long long l = 1, r = (A/2) + 1;
    while (l < r) {
        long long mid = l + (r - l) / 2;
        if (mid * mid == A) {
            return mid;
        } else if (mid * mid < A) {
            l = mid + 1;
        } else {
            r = mid;
        }
    }

    return -1;
}

// T(n) : O(n2logn) ; S(n) : O(1)
int pythagoreanTriplets1(int A) {

    int count = 0;
    for (int i = 1; i <= A; i++) {
        for (int j = i+1; j <= A; j++) {
            int k = i*i + j*j;
            int x = isPerfectSquare(k);
            if (x != -1 && x <= A) {
                count++;
            }
        }
    }

    return count;
}

// Driver Code for testing
int main() {

    cout << pythagoreanTriplets(5) << "\n";
    cout << pythagoreanTriplets(13) << "\n";
    cout << pythagoreanTriplets(900) << "\n";
    cout << pythagoreanTriplets(1000) << "\n";

    cout << isPerfectSquare(1) << "\n";
    cout << isPerfectSquare(2) << "\n";
    cout << isPerfectSquare(4) << "\n";
    cout << isPerfectSquare(9) << "\n";
    cout << isPerfectSquare(15) << "\n";
    cout << isPerfectSquare(1024) << "\n";

    cout << pythagoreanTriplets1(5) << "\n";
    cout << pythagoreanTriplets1(13) << "\n";
    cout << pythagoreanTriplets1(900) << "\n";
    cout << pythagoreanTriplets1(1000) << "\n";

    return 0;
}
