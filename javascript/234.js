function ListNode(val) {
  this.val = val;
  this.next = null;
}

/**
 * @param {ListNode} head
 * @return {boolean}
 */
var isPalindrome = function(head) {
  var arr = [];
  while (head) {
    arr.push(head.val);
    head = head.next;
  }

  for (var i = 0; i < parseInt(arr.length / 2); i++) {
    if (arr[i] !== arr[arr.length - i - 1]) {
      return false;
    }
  }

  return true;
};

/**
 * @param {ListNode} head
 * @return {boolean}
 */
var isPalindrome2 = function(head) {
  var cur = head;
  var prev = null;
  while (cur) {
    cur.prev = prev;
    prev = cur;
    cur = cur.next;
  }

  while (prev && head) {
    if (prev.val !== head.val) return false;
    prev = prev.prev;
    head = head.next;
  }

  return true;
};

const l = new ListNode(10);
l.next = new ListNode(2);
l.next.next = new ListNode(20);
isPalindrome2(l);
