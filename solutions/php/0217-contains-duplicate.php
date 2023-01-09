<?php
class Solution
{
    /**
     * @param Integer[] $nums
     * @return Boolean
     */
    function containsDuplicate(array $nums): bool
    {
        $items = [];

        foreach ($nums as $num) {
            if (isset($items[$num])) {
                return true;
            }

            $items[$num] = 1;
        }

        return false;
    }
}
