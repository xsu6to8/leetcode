/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2) {
        ListNode* t1 = l1; 
        ListNode* t2 = l2;
        
        bool isCarry = false; 

        ListNode* dummy = new ListNode(0);
        ListNode* curr = dummy;
        while(t1 != nullptr || t2 != nullptr || isCarry){
            int v1 = (t1 != nullptr) ? t1->val : 0;
            int v2 = (t2 != nullptr) ? t2->val : 0;

            int sum = v1 + v2;
            if (isCarry) {
                sum += 1;
                isCarry = false;
            }
            
            if (sum >= 10) {
                isCarry = true;
                sum %= 10;
            }
            
            curr->next = new ListNode(sum);
            curr = curr->next;

            if (t1 != nullptr) t1 = t1->next;
            if (t2 != nullptr) t2 = t2->next;
        }
        return dummy->next;
    }
};