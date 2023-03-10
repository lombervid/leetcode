<?php

class Solution
{
    /**
     * @param string $s
     * @return int
     */
    public function lengthOfLongestSubstring(string $s): int
    {
        $subStr = '';
        $lengths = [];

        foreach (str_split($s) as $char) {
            $pos = strpos($subStr, $char);

            if ($pos !== false) {
                $lengths[] = strlen($subStr);
                $subStr = substr($subStr, $pos + 1);
            }

            $subStr .= $char;
        }

        $lengths[] = strlen($subStr);

        return $lengths ? max($lengths) : 0;
    }
}
