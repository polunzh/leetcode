var process = require('process');
var os = require('os');
// TODO: 算法有待改进
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
var mergeTwoLists = function (l1, l2) {
    if (l1 === null) return l2;
    if (l2 === null) return l1;

    for (; l1 !== null; l1 = l1.next) {
        var t1 = new ListNode(l1.val);
        for (var t2 = l2; t2 !== null; t2 = t2.next) {
            if (t1.val < t2.val) {
                t1.next = t2;
                l2 = t1;
                break;
            } else if (t1.val >= l2.val) {
                if (t2.next !== null) {
                    if (t1.val <= t2.next.val) {
                        var t = t2.next;
                        t1.next = t;
                        t2.next = t1;
                        break;
                    }
                } else {
                    t2.next = t1;
                    break;
                }
            }
        }
    }

    return l2;
};

var print = function (l) {
    while (l !== null) {
        process.stdout.write(l.val + ' ');
        l = l.next;
    }
    process.stdout.write(os.EOL);
}

node1 = new ListNode(11)
node2 = new ListNode(12)
node3 = new ListNode(19)

node1.next = node2;
node2.next = node3;


nodea = new ListNode(3)
nodeb = new ListNode(5)
nodec = new ListNode(7)
noded = new ListNode(10)

nodea.next = nodeb;
nodeb.next = nodec;
nodec.next = noded;

function ListNode(val) {
    this.val = val;
    this.next = null;
}

l = mergeTwoLists(node1, nodea);
print(l);
l2 = mergeTwoLists(new ListNode(1), null);
print(l2);