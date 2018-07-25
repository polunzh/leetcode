/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

function ListNode(val) {
  this.val = val;
  this.next = null;
}

/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var reverseList = function(head) {
  if(head === null || head.next === null) {
    return head;
  }

  let p = head;
  let q = head.next;
  head.next = null;
  while(q) {
    const r = q.next;
    q.next = p;
    p = q;
    q = r;
  }

  head = p;
  return head;
}