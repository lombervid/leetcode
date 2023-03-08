<?php

class Solution
{
    /**
     * @param int[] $numbers
     * @param int $target
     * @return int[]
     */
    public function twoSum(array $numbers, int $target): array
    {
        $pairs = [];

        foreach ($numbers as $key => $value) {
            if (array_key_exists($target - $value, $pairs)) {
                return [$pairs[$target - $value], $key + 1];
            }

            $pairs[$value] = $key + 1;
        }

        return [];
    }
}
