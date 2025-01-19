#include <iostream>

using namespace std;

// https://leetcode.com/problems/trapping-rain-water/description/
class Solution {
private:
    // using brute-force approach
    // T(n) : O(n2) ; S(n) : O(1)
    int solvetrap(vector<int>& height) {

        int n = height.size(), trappedWater = 0;
        for (int i = 0; i < n; i++) {
            if (i == 0 || i == n-1) continue;

            int ml = *max_element(height.begin(), height.begin()+i);
            int mr = *max_element(height.begin()+i+1, height.end());

            int minH = min(ml, mr);
            trappedWater += max(0, minH - height[i]);
        }

        return trappedWater;
    }

    // using dynamic-programming
    // T(n) : O(n) ; S(n) : O(n)
    int solvetrap1(vector<int>& height) {

        int n = height.size(), trappedWater = 0;
        vector<int> arr(n, 0);
        arr[n-1] = height[n-1];
        for (int i = n-2; i>=0; i--) {
            arr[i] = max(height[i], arr[i+1]);
        }

        int lm = INT_MIN;
        for (int i=0; i<n; i++) {
            lm = max(lm, height[i]);
            int minH = min(lm, arr[i]);
            trappedWater += max(0, minH - height[i]);
        }

        return trappedWater;
    }

    // using stacks
    // T(n) : O(n) ; S(n) : O(n)
    int solvetrap2(vector<int>& height) {

        int idx = 0, n = height.size(), trappedWater = 0;
        stack<int> st;
        while (idx < n) {
            while (!st.empty() && height[idx] > height[st.top()]) {
                int top = st.top();
                st.pop();
                if (st.empty()) break;
                int distance = idx - st.top() -1;
                int bounded_height = min(height[idx], height[st.top()] - height[top]);
                trappedWater += distance * bounded_height;
            }
            st.push(idx);
            idx++;
        }

        return trappedWater;
    }

    // using two-pointers
    // T(n) : O(n) ; S(n) : O(1)
    int solvetrap3(vector<int>& height) {

        int n = height.size(), trappedWater = 0;
        int left = 0, right = n - 1, lm = 0, rm = 0;
        while (left < right) {
            if (height[left] < height[right]) {
                if (height[left] >= lm) {
                    lm = height[left];
                } else {
                    trappedWater += lm - height[left];
                }
                left++;
            } else {
                if (height[right] >= rm) {
                    rm = height[right];
                } else {
                    trappedWater += rm - height[right];
                }
                right--;
            }
        }

        return trappedWater;
    }
public:
    int trap(vector<int>& height) {
        return solvetrap3(height);
    }
};

int main() {
    vector<int> height {0,1,0,2,1,0,1,3,2,1,2,1};
    Solution *sol = new Solution();
    cout << sol->trap(height) << "\n";
    return 0;
}
