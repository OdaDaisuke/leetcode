// https://leetcode.com/problems/add-two-numbers/
/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
 Output: 7 -> 0 -> 8
 Explanation: 342 + 465 = 807.
 */
class Solution {
public:
  ListNode* addTwoNumbers(ListNode *l1, ListNode *l2) {
    int moveup = 0;
    int sum = 0;
    ListNode *head = nullptr;
    ListNode **tail = &head;
    while (l1 != nullptr || l2 != nullptr) {
      sum = moveup + getNextVal(l1) + getNextVal(l2);
      // 前の桁で繰り上がってたら今回の桁で0になる
      ListNode *node = new ListNode(sum % 10);
      *tail = node;
      tail = (&node->next);
      // 下の桁から加算するので、繰り上がり状態をここで保持
      moveup = sum / 10;
    }
    // 1桁入力で繰り上がった時用
    if (moveup > 0) {
      ListNode *node = new ListNode(moveup % 10);
      *tail = node;
    }
    return head;
  }
  int getNextVal(ListNode* &l) {
    if (l == nullptr) {
      return 0;
    }
    int tmpV = l->val;
    l = l->next;
    return tmpV;
  }
};
