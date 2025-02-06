#include <iostream>
#include <vector>
#include <unordered_map>
#include <unordered_set>

using namespace std;

// https://leetcode.com/problems/tuple-with-same-product/description/
class Solution {
private:
    // T(n) : O(n3) ; S(n) : O(n)
    int solvetupleSameProduct1(vector<int>& nums) {
        int n = nums.size(), tuplesCount = 0;
        sort(nums.begin(), nums.end());
        for (int i=0; i < n-1; i++) {
            for (int j=n-1; j>i; j--) {
                int p = nums[i] * nums[j];
                unordered_set<int> us;
                for (int k=i+1; k<j; k++) {
                    if (p%nums[k] == 0) {
                        int d = p/nums[k];
                        if (us.find(d) != us.end()) {
                            tuplesCount += 8;
                        }
                        us.insert(nums[k]);
                    }
                }
            }
        }
        return tuplesCount;
    }

    // using hash-map
    // T(n) : O(n2) ; S(n) : O(n2)
    int solvetupleSameProduct(vector<int>& nums) {
        int n = nums.size();
        unordered_map<int, int> um;
        for (int i=0; i<n; i++) {
            for (int j=i+1; j<n; j++) {
                um[nums[i]*nums[j]]++;
            }
        }

        int count = 0;
        for (auto &[p, f] : um) {
            int pairsCount = f*(f-1)/2;
            count += 8*pairsCount;
        }

        return count;
    }
public:
    int tupleSameProduct(vector<int>& nums) {
        return solvetupleSameProduct1(nums);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
