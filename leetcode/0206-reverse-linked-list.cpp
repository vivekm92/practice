#include <iostream>

using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

// https://leetcode.com/problems/reverse-linked-list/description/
class Solution {
private:
    // Approach 1
    // recursive approach
    // T(n) : O(n) ; S(n) : O(n)
    ListNode* solveReverseList(ListNode *head) {
        if (head == nullptr || head->next == nullptr) {
            return head;
        }

        ListNode *t = solveReverseList(head->next);
        head->next->next = head;
        head->next = nullptr;

        return t;
    }

    // Approach 2
    // iterative approach
    // T(n) : O(n) ; S(n) : O(1)
    ListNode* solveReverseList1(ListNode *head) {
        if (head == nullptr || head->next == nullptr) {
            return head;
        }

        ListNode *curr = head, *prev = nullptr;
        while (curr != nullptr) {
            ListNode *t = curr->next;
            curr->next = prev;
            prev = curr;
            curr = t;
        }

        return prev;
    }
public:
    ListNode* reverseList(ListNode* head) {
        return solveReverseList1(head);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
