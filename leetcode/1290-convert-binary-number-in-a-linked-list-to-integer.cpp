#include <bits/stdc++.h>

using namespace std;


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
	int solveGetDecimalValue(ListNode *head) {
		
		int value = 0;
		while (head != nullptr) {
			value = value * 2 + head->val;
			head = head->next;
		}
		
		return value;
	}
public:
    int getDecimalValue(ListNode* head) {
        return solveGetDecimalValue(head);
    }
};

// Driver Code for testing
int main() {
	
	return 0;
}
