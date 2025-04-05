#include <iostream>

using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

// https://leetcode.com/problems/swap-nodes-in-pairs/description/
class Solution {
private:
    // Approach 1
    // using recursion
    // T(n) : O(n) ; S(n) : O(n)
    ListNode* solveSwapPairs(ListNode *head) {
       if (head == nullptr || head->next == nullptr) {
           return head;
       }

       ListNode* nextNode = head->next->next;
       head->next->next = head;

       ListNode *temp = solveSwapPairs(nextNode);
       ListNode *node = head->next;
       head->next = temp;

       return node;
    }

    // Approach 2
    // iterative approach
    // T(n) : O(n) ; S(n) : O(1)
    ListNode* solveSwapPairs1(ListNode *head) {
        if (head == nullptr || head->next == nullptr) {
            return head;
        }

        ListNode *curr = head, *prev = nullptr;
        head = curr->next;
        while (curr != nullptr && curr->next != nullptr) {
            ListNode *nextNode = curr->next->next;
            curr->next->next = curr;
            if (prev != nullptr) prev->next = curr->next;
            curr->next = nextNode;
            prev = curr;
            curr = nextNode;
        }

        return head;
    }
public:
    ListNode* swapPairs(ListNode* head) {
        return solveSwapPairs1(head);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
