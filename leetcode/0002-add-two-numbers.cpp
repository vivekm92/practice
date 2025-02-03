#include <iostream>

using namespace std;

// https://leetcode.com/problems/add-two-numbers/description/

// Definition for singly-linked list.
struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};
 
class Solution {
private:
    // T(n) : O(n) ; S(n) : O(n)
    ListNode* solveaddTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode *dummy = new ListNode(-1);
        ListNode *curr = dummy;
        int carry = 0;
        while (l1 != nullptr || l2 != nullptr) {
            int sum = carry;
            if (l1 != nullptr) {
                sum += l1->val;
                l1 = l1->next;
            }
            if (l2 != nullptr) { 
                sum += l2->val;
                l2 = l2->next;
            }
            curr->next = new ListNode(sum % 10);
            curr = curr->next;
            carry = sum / 10;
        }
        if (carry != 0) curr->next = new ListNode(carry);
        // ListNode *res = dummy->next;
        // delete dummy;
        return res;
    }
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        return solveaddTwoNumbers(l1, l2);
    }
};

// Driver Code for testing
int main() {
    
}
