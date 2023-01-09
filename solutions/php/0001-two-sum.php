<?php
class Solution
{
    /**
     * @param Integer[] $nums
     * @param Integer $target
     * @return Integer[]
     */
    function twoSum(array $nums, int $target): ?array
    {
        $rest = [];

        foreach ($nums as $i => $num) {
            if (array_key_exists($target - $num, $rest)) {
                return [$i, $rest[$target - $num]];
            }

            $rest[$num] = $i;
        }

        return null;
    }
}
