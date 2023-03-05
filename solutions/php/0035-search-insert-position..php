<?php

class Solution
{
    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer
     */
    public function searchInsert(array $nums, int $target): int
    {
        $index = count($nums);
        $left = 0;
        $right = $index - 1;

        while ($left <= $right) {
            $mid = intval(($left + $right) / 2);

            if ($target === $nums[$mid]) {
                return $mid;
            }

            if ($target < $nums[$mid]) {
                $index = $mid;
                $right = $mid - 1;
            } else {
                $left = $mid + 1;
            }
        }

        return $index;
    }
}
