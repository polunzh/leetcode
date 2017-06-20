/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var deleteDuplicates = function (head) {
    let cursor = head;
    while (cursor && cursor.next) {
        if (cursor.next.val === cursor.val) {
            cursor.next = cursor.next.next;
        } else {
            cursor = cursor.next;
        }
    }

    return head;
};

function ListNode(val) {
    this.val = val;
    this.next = null;
}

let n1 = new ListNode(1);
let n2 = new ListNode(2);
let n21 = new ListNode(2);
let n3 = new ListNode(3);
let n4 = new ListNode(3);
n1.next = n2;
n2.next = n21;
n21.next = n3;
n3.next = n4;

console.log(deleteDuplicates(n1));