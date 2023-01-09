<?php
/**
 * Definition for a singly-linked list.
 * class ListNode {
 *     public $val = 0;
 *     public $next = null;
 *     function __construct($val = 0, $next = null) {
 *         $this->val = $val;
 *         $this->next = $next;
 *     }
 * }
 */
class Solution
{
    /**
     * @param ListNode $l1
     * @param ListNode $l2
     * @return ListNode
     */
    function addTwoNumbers(ListNode $l1, ListNode $l2): ListNode
    {
        $list = new ListNode();
        $last = $list;
        $carry = 0;

        while ($l1 || $l2 || $carry) {
            $sum = ($l1 ? $l1->val : 0) + ($l2 ? $l2->val : 0) + $carry;
            $carry = floor($sum / 10);
            $sum %= 10;

            $last->next = new ListNode($sum);
            $last = $last->next;

            $l1 = $l1?->next;
            $l2 = $l2?->next;
        }

        return $list->next;
    }
}
