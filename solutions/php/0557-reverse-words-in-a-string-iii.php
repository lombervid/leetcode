<?php

class Solution
{
    /**
     * @param string $s
     * @return string
     */
    public function reverseWords(string $s): string
    {
        $words = explode(' ', $s);
        $count = count($words);

        for ($i = 0; $i < $count; $i++) {
            $words[$i] = strrev($words[$i]);
        }

        return implode(' ', $words);
    }
}
