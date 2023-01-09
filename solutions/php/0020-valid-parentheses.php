<?php
class Solution
{
    /**
     * @param String $s
     * @return Boolean
     */
    function isValid($s)
    {
        $stack = [];
        $brackets = [')' => '(', ']' => '[', '}' => '{'];

        foreach (str_split($s) as $i => $char) {
            if (!array_key_exists($char, $brackets)) {
                $stack[] = $char;
                continue;
            }

            if (array_pop($stack) != $brackets[$char]) {
                return false;
            }
        }

        return count($stack) == 0;
    }
}
