<?php

class Solution
{
    /**
     * @param int[] $nums
     * @param int $k
     * @return NULL
     */
    public function rotate(array &$nums, int $k)
    {
        $len = count($nums);

        $k = $k % $len;

        if ($k === 0) {
            return;
        }

        $nums = array_merge(array_slice($nums, $len - $k, $k), array_slice($nums, 0, $len - $k));
    }
}
