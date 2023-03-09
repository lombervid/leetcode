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
     * @param int $n
     * @return ListNode
     */
    public function removeNthFromEnd(ListNode $head, int $n): ?ListNode
    {
        $newHead = new ListNode(next: $head);
        $nth = $newHead;
        $len = 0;

        while ($head = $head->next) {
            if (++$len < $n) {
                continue;
            }
            $nth = $nth->next;
        }
        $nth->next = $nth->next->next;

        return $newHead->next;
    }
}
