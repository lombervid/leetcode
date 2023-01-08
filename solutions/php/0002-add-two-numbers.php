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
    function addTwoNumbers($l1, $l2)
    {
        $list = new ListNode();
        $last = $list;
        $carry = 0;

        while ($l1 || $l2 || $carry) {
            $num1 = 0;
            $num2 = 0;

            if ($l1) {
                $num1 = $l1->val;
                $l1 = $l1->next;
            }

            if ($l2) {
                $num2 = $l2->val;
                $l2 = $l2->next;
            }

            if (($sum = $num1 + $num2 + $carry) >= 10) {
                $sum -= 10;
                $carry = 1;
            } else {
                $carry = 0;
            }

            $last->next = new ListNode($sum);
            $last = $last->next;
        }

        return $list->next;
    }
}
