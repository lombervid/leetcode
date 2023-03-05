<?php

class Solution
{
    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer
     */
    public function search(array $nums, int $target): int
    {
        $left = 0;
        $right = count($nums) - 1;

        while ($left <= $right) {
            $mid = intval(($left + $right) / 2);

            if ($target === $nums[$mid]) {
                return $mid;
            }

            if ($target > $nums[$mid]) {
                $left = $mid + 1;
            } else {
                $right = $mid - 1;
            }
        }

        return -1;
    }
}
