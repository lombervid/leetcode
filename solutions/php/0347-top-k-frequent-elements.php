<?php

class Solution {

    /**
     * @param Integer[] $nums
     * @param Integer $k
     * @return Integer[]
     */
    function topKFrequent($nums, $k) {
        $counts = [];

        foreach ($nums as $key => $num) {
            $counts[$num] = ($counts[$num] ?? 0) + 1;
        }

        arsort($counts);

        return array_slice(array_keys($counts), 0, $k);
    }
}
