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
    public function middleNode(ListNode $head): ListNode
    {
        $counter = 0;
        $middle = $head;

        while ($head = $head->next) {
            if ($counter++ % 2 === 0) {
                $middle = $middle->next;
            }
        }

        return $middle;
    }
}
