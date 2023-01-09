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
     * @param ListNode $head
     * @return ListNode
     */
    function reverseList(?ListNode $head): ?ListNode
    {
        $reversed = null;

        while ($head) {
            [$head->next, $head, $reversed] = [$reversed, $head->next, $head];
        }

        return $reversed;
    }
}
