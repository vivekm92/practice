#include <iostream>

using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

// https://leetcode.com/problems/merge-two-sorted-lists/description/
class Solution {
private:

    // Approach 1
    // recursive
    // T(n) : O(n) ; S(n) : O(n)
    ListNode* solveMergeTwoLists(ListNode *list1, ListNode *list2) {
        if (list1 == nullptr) return list2;
        else if (list2 == nullptr) return list1;

        if (list1->val < list2->val) {
            list1->next = solveMergeTwoLists(list1->next, list2);
            return list1;
        }

        list2->next = solveMergeTwoLists(list1, list2->next);
        return list2;
    }

    // Approach 2
    // iterative
    // T(n) : O(n) ; S(n) : O(1)
    ListNode* solveMergeTwoLists1(ListNode *list1, ListNode *list2) {

        ListNode *dummy = new ListNode(-1);
        ListNode *prevHead = dummy;
        while (list1 != nullptr || list2 != nullptr) {

            if (list1 == nullptr) {
                dummy->next = list2;
                list2 = nullptr;
            } else if (list2 == nullptr) {
                dummy->next = list1;
                list1 = nullptr;
            } else if (list1->val <= list2->val) {
                dummy->next = list1;
                list1 = list1->next;
            } else {
                dummy->next = list2;
                list2 = list2->next;
            }
            dummy = dummy->next;
        }

        return prevHead->next;
    }
public:
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        return solveMergeTwoLists1(list1, list2);
    }
};

// Driver Code for testing
int main() {


    return 0;
}
