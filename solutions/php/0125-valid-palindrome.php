<?php
class Solution
{
    /**
     * @param String $s
     * @return Boolean
     */
    function isPalindrome(string $s): bool
    {
        $s = preg_replace('/[^a-z0-9]/i', '', strtolower($s));

        return $s === strrev($s);
    }
}
