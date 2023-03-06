<?php

class Solution
{
    /**
     * @param int[] $nums
     * @return int[]
     */
    public function sortedSquares(array $nums): array
    {
        $i = 0;
        $j = count($nums) - 1;
        $index = $j;
        $squares = array_fill(0, count($nums), 0);

        while ($index >= 0) {
            $iSqrt = $nums[$i] ** 2;
            $jSqrt = $nums[$j] ** 2;

            if ($iSqrt > $jSqrt) {
                $squares[$index] = $iSqrt;
                $i++;
            } else {
                $squares[$index] = $jSqrt;
                $j--;
            }
            $index--;
        }

        return $squares;
    }
}
