<?php

/* The isBadVersion API is defined in the parent class VersionControl.
      public function isBadVersion($version){} */

class Solution extends VersionControl
{
    /**
     * @param Integer $n
     * @return Integer
     */
    public function firstBadVersion(int $n): int
    {
        $left = 1;
        $right = $n;
        $firstBad = $n;

        while ($left <= $right) {
            $currentVersion = intval(($left + $right) / 2);

            if ($this->isBadVersion($currentVersion)) {
                $firstBad = $currentVersion;
                $right = $currentVersion - 1;
            } else {
                $left = $currentVersion + 1;
            }
        }

        return $firstBad;
    }
}
