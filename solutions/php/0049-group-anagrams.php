<?php
class Solution
{
    /**
     * @param String[] $strs
     * @return String[][]
     */
    function groupAnagrams(array $strs): array
    {
        $anagrams = [];

        foreach ($strs as $i => $word) {
            $anagrams[$this->sortStr($word)][] = $word;
        }

        return $anagrams;
    }

    private function sortStr(string $str): string
    {
        $arr = str_split($str);
        sort($arr);

        return implode($arr);
    }
}
