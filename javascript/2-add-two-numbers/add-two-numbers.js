/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/** 
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function (l1, l2) {
    let n = 0, list = l1;
    while (l1 !== null && l2 !== null) {
        l1.val = l1.val + l2.val + n;
        if (l1.val > 9) {
            l1.val = l1.val % 10;
            n = 1;
        } else {
            n = 0;
        }

        if (l1.next === null && l2.next === null) {
            if (n === 1) {
                l1.next = new ListNode(n);
            }
            return list;
        }

        l1 = l1.next;
        l2 = l2.next;
    }


    l1 = l1 === null ? l2 : l1;
    console.log(l1)

    while (l1 !== null) {
        l1.val = l1.val + n;

        if (l1.val > 9) {
            l1.val = l1.val % 10;
            n = 1;
        } else {
            n = 0;
        }

        l1 = l1.next;
    }

    if (n === 1) {
        l1 = new ListNode(1);
    }

    return list;
};

function ListNode(val) {
    this.val = val;
    this.next = null;
}

let list1 = new ListNode(2);
let list2 = new ListNode(4);
let list3 = new ListNode(3);
let list4 = new ListNode(5);
let list5 = new ListNode(6);
let list6 = new ListNode(4);

list1.next = list2;
list2.next = list3;
list4.next = list5;
list5.next = list6;

// console.log(addTwoNumbers(list1, list4));

let l1 = new ListNode(0);
let l2 = new ListNode(7);
let l3 = new ListNode(3);
l2.next = l3;

console.log(addTwoNumbers(l1, l2));