<?php
class Solution
{
    /**
     * @param String $s
     * @return Boolean
     */
    function isPalindrome($s)
    {
        $s = preg_replace('/[^a-z0-9]/i', '', strtolower($s));

        return $s === strrev($s);
    }
}
