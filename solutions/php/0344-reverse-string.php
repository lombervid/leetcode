<?php

class Solution
{
    /**
     * @param string[] $s
     * @return void
     */
    public function reverseString(array &$s): void
    {
        for ($i = 0, $j = count($s) - 1; $i < count($s) / 2; $i++, $j--) {
            [$s[$i], $s[$j]] = [$s[$j], $s[$i]];
        }
    }

    // Another solution using array_reverse() function
    public function reverseString2(array &$s): void
    {
        $s = array_reverse($s);
    }
}
