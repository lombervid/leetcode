<?php

class Solution
{
    /**
     * @param Integer[] $nums
     * @return NULL
     */
    public function moveZeroes(array &$nums): void
    {
        foreach ($nums as $key => $value) {
            if ($value === 0) {
                unset($nums[$key]);
                $nums[] = 0;
            }
        }
    }

    // Another (less efficient) solution with a custom sort function
    public function moveZeroes2(array &$nums): void
    {
        usort($nums, function (int $a, int $b): int {
            if ($a !== 0 && $b !== 0) {
                return 0;
            }

            if ($a === $b) {
                return 0;
            }

            return $a === 0 ? 1 : -1;
        });
    }
}
