function ListNode(val) {
  this.val = val;
  this.next = null;
}

/**
 * 删除当前节点：使用当前节点的下一个节点覆盖当前节点、删除当前节点的下一个节点
 * @param {ListNode} node
 * @return {void} Do not return anything, modify node in-place instead.
 */
var deleteNode = function(node) {
  node.val = node.next.val;
  node.next = node.next.next;
};
