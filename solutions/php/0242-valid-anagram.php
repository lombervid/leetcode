<?php

class Solution
{
    /**
     * @param String $s
     * @param String $t
     * @return Boolean
     */
    /* function isAnagram(string $s, string $t): bool
    {
        if (strlen($s) != strlen($t)) {
            return false;
        }

        $counts = count_chars($s, 1);

        foreach (str_split($t) as $char) {
            $index = ord($char);

            if (!isset($counts[$index])) {
                return false;
            }

            if (--$counts[$index] < 0) {
                return false;
            }
        }

        return true;
    } */

    function isAnagram(string $s, string $t): bool
    {
        return count_chars($s) === count_chars($t);
    }
}
